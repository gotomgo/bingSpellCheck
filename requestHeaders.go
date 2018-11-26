package bingSpellCheck

import "net/http"

const (
	// SubscriptionKeyHeader is the header for the required API Key
	SubscriptionKeyHeader = "Ocp-Apim-Subscription-Key"

	// AcceptHeader is the standard HTTP Accept header, and defaults to 'application/json'
	AcceptHeader = "Accept"

	// AcceptLanguageHeader is the standard HTTP Accept-Language header
	//
	//  Notes
	//    Comma seperated lists of languges in decreasing order of preference.
	//    See RFC2616 for more info on format
	//
	AcceptLanguageHeader = "Accept-Language"

	// AcceptApplicationJSON is the standard Accept header value for JSON
	AcceptApplicationJSON = "application/json"

	// UserAgentHeader is the standard HTTP User-Agent header
	UserAgentHeader = "User-Agent"

	// PragmaHeader is the standard HTTP Pragma header
	PragmaHeader = "Pragma"

	// ClientIDHeader is used to provide users with consistent
	// behavior across Bing API calls
	ClientIDHeader = "X-MSEdge-ClientID"

	// ClientIPHeader The IPv4 or IPv6 address of the client device
	//
	//  Notes
	//    Bing uses the ClientIPto discover the user's location. Bing uses the
	//    location information to determine safe search behavior
	//
	ClientIPHeader = "X-MSEdge-ClientIP"

	// SearchLocationHeader is semicolon-delimited list of key/value pairs that
	// describe the client's geographical location
	SearchLocationHeader = "X-Search-Location"

	// AcceptApplicationLinkedJSON is the Accept header value for Linked JSON
	AcceptApplicationLinkedJSON = "application/ld+json"

	// PragmaNoCache is the standard Pragma header value of no-cache
	PragmaNoCache = "no-cache"
)

// SpellCheckHeaders are HTTP header values that may be sent with a request
// to the Bing Spell Check API
//
//  Notes
//    For more information see:
//      https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-spell-check-api-v7-reference#request-headers
//
type SpellCheckHeaders struct {
	Headers http.Header
}

// NewSpellCheckHeaders returns a *SpellCheckHeaders using a specified
// Bing subscription key
func NewSpellCheckHeaders(subscriptionKey string) *SpellCheckHeaders {
	return (&SpellCheckHeaders{Headers: http.Header{}}).WithSubscriptionKey(subscriptionKey).WithJSON()
}

// SetHeader sets the value of a header, or removes it if value is empty
func (sch *SpellCheckHeaders) SetHeader(header, value string) *SpellCheckHeaders {
	if len(value) > 0 {
		sch.Headers.Set(header, value)
	} else {
		sch.Headers.Del(header)
	}

	return sch
}

// WithJSON sets the accept header to 'application/json'
func (sch *SpellCheckHeaders) WithJSON() *SpellCheckHeaders {
	return sch.SetHeader(AcceptHeader, AcceptApplicationJSON)
}

// WithLinkedJSON sets the accept header to 'application/ld+json'
func (sch *SpellCheckHeaders) WithLinkedJSON() *SpellCheckHeaders {
	return sch.SetHeader(AcceptHeader, AcceptApplicationLinkedJSON)
}

// WithSubscriptionKey sets the SubscriptionKey header
//
//  Notes
//    This header values is *required*
//
func (sch *SpellCheckHeaders) WithSubscriptionKey(subscriptionKey string) *SpellCheckHeaders {
	return sch.SetHeader(SubscriptionKeyHeader, subscriptionKey)
}

// WithUserAgent sets the UserAgent header (see RFC 2616 for format)
//
//  Notes
//    Passing in an empty string effectively removes the Header
//
func (sch *SpellCheckHeaders) WithUserAgent(userAgent string) *SpellCheckHeaders {
	return sch.SetHeader(UserAgentHeader, userAgent)
}

// WithNoCachePragma sets the Pragma header to no-cache
func (sch *SpellCheckHeaders) WithNoCachePragma() *SpellCheckHeaders {
	return sch.SetHeader(PragmaHeader, PragmaNoCache)
}

// WithAllowCachePragma removes the Pragma header
func (sch *SpellCheckHeaders) WithAllowCachePragma() *SpellCheckHeaders {
	sch.Headers.Del(PragmaHeader)
	return sch
}

// WithClientID sets the ClientID header
//
//  Notes
//    Passing in an empty string effectively removes the Header
//
func (sch *SpellCheckHeaders) WithClientID(clientID string) *SpellCheckHeaders {
	return sch.SetHeader(ClientIDHeader, clientID)
}

// WithClientIP sets the ClientIP header
//
//  Notes
//    Passing in an empty string effectively removes the Header
//
func (sch *SpellCheckHeaders) WithClientIP(clientIP string) *SpellCheckHeaders {
	return sch.SetHeader(ClientIPHeader, clientIP)
}

// WithAcceptLanguages sets the AcceptLanaguage header (see RFC2616 for format)
//
//  Notes
//    Passing in an empty string effectively removes the Header
//
func (sch *SpellCheckHeaders) WithAcceptLanguages(languages string) *SpellCheckHeaders {
	return sch.SetHeader(AcceptLanguageHeader, languages)
}

// WithSearchLocation sets the SearchLocation header (see Bing Spell Check docs
// for values and expected format)
//
//  Notes
//    Passing in an empty string effectively removes the Header
//
func (sch *SpellCheckHeaders) WithSearchLocation(location string) *SpellCheckHeaders {
	return sch.SetHeader(SearchLocationHeader, location)
}
