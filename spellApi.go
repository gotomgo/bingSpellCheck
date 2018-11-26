package bingSpellCheck

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// BingHost is the host/domain for Bing Spell Check API
const BingHost = "https://api.cognitive.microsoft.com"

// BingSpellCheckPath is the url path of the Bing spell check, version 7, API
const BingSpellCheckPath = "/bing/v7.0/spellcheck"

// Client is a simple Bing Spell Check API client that is most useful when
// the values of Params and Headers remain constant.
//
//  Notes
//    For server implementations, use a unique instance of Client per user
//    as it is not thread-safe. Alternatively, build params and headers
//    on the fly and call func SpellCheck.
//
type Client struct {
	Params  *SpellCheckParams
	Headers *SpellCheckHeaders

	spellCheckURL string
	httpClient    *http.Client
}

// GetSpellCheckURL returns the URL for the Bing Spell Check version 7 API
func GetSpellCheckURL() string {
	u, _ := url.ParseRequestURI(BingHost)
	u.Path = BingSpellCheckPath
	return u.String()
}

// NewClient creates an instance of the Bing Spell Check API client
func NewClient(subscriptionKey string) *Client {
	return &Client{
		Params:        NewSpellCheckParams(),
		Headers:       NewSpellCheckHeaders(subscriptionKey),
		spellCheckURL: GetSpellCheckURL(),
		httpClient:    &http.Client{Timeout: time.Second * 30},
	}
}

// SpellCheck is the core function for accessing the Bing Spell Check API
func SpellCheck(
	httpClient *http.Client,
	targetURL string,
	params *SpellCheckParams,
	headers *SpellCheckHeaders) (*SpellCheckResponse, error) {

	var err error

	// if the length of the text is excessively long we need to POST, not GET
	postRequired := params.TotalTextLength() > 1500

	q := params.Values

	var r *http.Request

	if postRequired {
		r, err = http.NewRequest(http.MethodPost, targetURL, strings.NewReader(q.Encode()))
		if err != nil {
			return nil, err
		}
		// form encoded
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		r, err = http.NewRequest(http.MethodGet, targetURL, nil)
		if err != nil {
			return nil, err
		}
		r.URL.RawQuery = q.Encode()
	}

	r.Header = headers.Headers

	resp, err := httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var spellCheck SpellCheckResponse
	err = json.Unmarshal(bodyBytes, &spellCheck)
	if err != nil {
		return nil, err
	}

	return &spellCheck, nil
}

// SpellCheck performs a spelling and/or grammar check on text
func (client *Client) SpellCheck(text string) (*SpellCheckResponse, error) {
	client.Params.WithTextAndContext(text, "", "")
	return SpellCheck(client.httpClient, client.spellCheckURL, client.Params, client.Headers)
}

// SpellCheckWithContext performs a spelling and/or grammar check on text with optional
// pre/post context
func (client *Client) SpellCheckWithContext(text, preContext, postContext string) (*SpellCheckResponse, error) {
	client.Params.WithTextAndContext(text, preContext, postContext)
	return SpellCheck(client.httpClient, client.spellCheckURL, client.Params, client.Headers)
}

// AutoCorrect performs a spell check and corrects the text based on corrections
// from the response
func (client *Client) AutoCorrect(text string) (string, error) {
	scr, err := client.SpellCheck(text)
	if err != nil {
		return "", err
	}

	return BuildAutoCorrectedText(text, scr)
}
