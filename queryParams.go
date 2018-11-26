package bingSpellCheck

import "net/url"

const (
	// ActionTypeParam is string that's used by logging to determine whether
	// the request is coming from an interactive session or a page load
	//
	//  Notes
	//    See EditActionType and LoadActionType
	ActionTypeParam = "actionType"

	// AppNameParam is the unique name of your app
	//
	//  Notes
	//    AppName must be known by Bing. Do not include this parameter unless you
	//    have previously contacted Bing to get a unique app name.
	//
	AppNameParam = "appName"

	// ClientMachineNameParam is a unique name of the device that the request is
	// being made from
	//
	//  Notes
	//    The specific value of ClientMachineNameParam is not important, but it
	//    should be unique for each device
	//
	ClientMachineNameParam = "clientMachineName"

	// DocumentIDParam is a unique ID that identifies the document that the
	// text belongs to
	//
	//  Notes
	//    The specific value of DocumentIDParam is not important, but it should
	//    be unique for each document
	//
	DocumentIDParam = "docId"

	// ModeParam is the type of spelling and grammar checks to perform
	//
	//  Notes
	//    See ProofMode and SpellMode
	ModeParam = "mode"

	// TextParam is the string to check for spelling and/or grammar errors
	TextParam = "text"

	// PreContextTextParam is a string that gives context to the text that precedes
	// the TextParam string
	PreContextTextParam = "preContextText"

	// PostContextTextParam is a string that gives context to the text that follows
	// the TextParam string
	PostContextTextParam = "postContextText"

	// SessionIDParam is a unique ID that identifies the user session
	//
	//  Notes
	//    The specific value of SessionIDParam is not important, but it should
	//    be unique for each user session
	//
	SessionIDParam = "sessionId"

	// UserIDParam is a unique ID that identifies the user
	//
	//  Notes
	//    The specific value of UserIDParam is not important, but it should
	//    be unique for each user
	//
	UserIDParam = "userId"

	// MarketParam is the market where the results come from
	//
	//  Notes
	//    The value for MarketParam must be in the form:
	//       <language code>-<country code>.
	//
	//    Example: en-US.
	//
	//    The string is case insensitive
	MarketParam = "mkt"

	// CountryCodeParam is the 2-character country code of the country where the
	// results come from
	CountryCodeParam = "cc"

	// LanguageParam is the language to use for user interface strings
	//
	//  Notes
	//    For the value of Language use the ISO 639-1 2-letter language code.
	//    Although optional, you should always specify Language. Typically, you
	//    set Language to the same language specified by Market unless the user
	//    wants the user interface strings displayed in a different language.
	//    The Language parameter and the Accept-Language header are mutually
	//    exclusive; do not specify both
	//
	LanguageParam = "setLang"

	// EditActionType indicates the request is from an interactive session
	EditActionType = "edit"

	// LoadActionType indicates the request is from a page load
	LoadActionType = "load"

	// ProofMode is grammar and spell checking
	//
	//  Notes
	//    ProofMode is the default value of ModeParam
	//
	ProofMode = "proof"

	// SpellMode provides spell checking
	SpellMode = "spell"
)

// SpellCheckParams encapsulates the possible query parameters for the
// SpellCheckAPI
//
//  Notes
//    For complete documentation see:
//      https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-spell-check-api-v7-reference#query-parameters
//
type SpellCheckParams struct {
	Values url.Values
}

// NewSpellCheckParams creates an instance of SpellCheckParams
func NewSpellCheckParams() *SpellCheckParams {
	return &SpellCheckParams{Values: url.Values{}}
}

// SetParam sets the value of a parameter, or if value is empty, removes it
func (scp *SpellCheckParams) SetParam(param, value string) *SpellCheckParams {
	if len(value) > 0 {
		scp.Values.Set(param, value)
	} else {
		scp.Values.Del(param)
	}

	return scp
}

// WithEditAction sets the ActionType parameter to 'edit'
func (scp *SpellCheckParams) WithEditAction() *SpellCheckParams {
	return scp.SetParam(ActionTypeParam, EditActionType)
}

// WithLoadAction sets the ActionType parameter to 'load'
func (scp *SpellCheckParams) WithLoadAction() *SpellCheckParams {
	return scp.SetParam(ActionTypeParam, LoadActionType)
}

// WithAppName sets the AppName parameter
func (scp *SpellCheckParams) WithAppName(appName string) *SpellCheckParams {
	return scp.SetParam(AppNameParam, appName)
}

// WithClientMachineName sets the ClientMachineName parameter
func (scp *SpellCheckParams) WithClientMachineName(clientMachineName string) *SpellCheckParams {
	return scp.SetParam(ClientMachineNameParam, clientMachineName)
}

// WithDocumentID sets the DocumentID parameter
func (scp *SpellCheckParams) WithDocumentID(documentID string) *SpellCheckParams {
	return scp.SetParam(DocumentIDParam, documentID)
}

// WithProofMode sets the Mode parameter to 'proof'
func (scp *SpellCheckParams) WithProofMode() *SpellCheckParams {
	return scp.SetParam(ModeParam, ProofMode)
}

// WithSpellMode sets the Mode parameter to 'spell'
func (scp *SpellCheckParams) WithSpellMode() *SpellCheckParams {
	return scp.SetParam(ModeParam, SpellMode)
}

// WithPreContextText sets the PreContextText parameter
func (scp *SpellCheckParams) WithPreContextText(text string) *SpellCheckParams {
	return scp.SetParam(PreContextTextParam, text)
}

// WithPostContextText sets the PostContextText parameter
func (scp *SpellCheckParams) WithPostContextText(text string) *SpellCheckParams {
	return scp.SetParam(PostContextTextParam, text)
}

// WithText sets the Text parameter
func (scp *SpellCheckParams) WithText(text string) *SpellCheckParams {
	return scp.SetParam(TextParam, text)
}

// WithTextAndContext sets the Text, Pre, and Post Context parameters
func (scp *SpellCheckParams) WithTextAndContext(text, preContext, postContext string) *SpellCheckParams {
	return scp.WithText(text).WithPreContextText(preContext).WithPostContextText(postContext)
}

// WithSessionID sets the SessionID parameter
func (scp *SpellCheckParams) WithSessionID(sessionID string) *SpellCheckParams {
	return scp.SetParam(SessionIDParam, sessionID)
}

// WithUserID sets the UserID parameter
func (scp *SpellCheckParams) WithUserID(userID string) *SpellCheckParams {
	return scp.SetParam(UserIDParam, userID)
}

// WithMarket sets the Market parameter
func (scp *SpellCheckParams) WithMarket(market MarketCode) *SpellCheckParams {
	return scp.SetParam(MarketParam, string(market))
}

// WithCountryCode sets the CountryCode parameter
func (scp *SpellCheckParams) WithCountryCode(cc CountryCode) *SpellCheckParams {
	return scp.SetParam(CountryCodeParam, string(cc))
}

// WithLangauge sets the Langague parameter
func (scp *SpellCheckParams) WithLangauge(lang string) *SpellCheckParams {
	return scp.SetParam(LanguageParam, lang)
}

// TotalTextLength returns the sum of the length of all text fields
func (scp *SpellCheckParams) TotalTextLength() int {
	return len(scp.Values[TextParam]) + len(scp.Values[PreContextTextParam]) + len(scp.Values[PostContextTextParam])
}
