package bingSpellCheck

import "fmt"

const (
	// ErrorResponseType is used as a value for SpellCheckResponse.Type and
	// indicates a request error occured
	ErrorResponseType = "ErrorResponse"
	// SpellCheckResponseType is used as a value for SpellCheckResponse.Type and
	// indicates a successful request
	SpellCheckResponseType = "SpellCheck"

	// RepeatedTokenType indicates a repeated word in a spell/grammar check
	RepeatedTokenType = "RepeatedToken"
	// UnknownTokenType indicates a spelling / grammatical error in the text
	UnknownTokenType = "UnknownToken"
)

// SpellCheckResponse is the main data return from spell check API calls
//
//  Fields
//    Type          - Type hint ("SpellCheck" or "ErrorResponse")
//    FlaggedTokens - A list of words in text that were flagged as not being
//      spelled correctly or are grammatically incorrect.
//    Errors        - A list of errors that describe the reasons why the
//      request failed
//
//  Notes
//    If no spelling or grammar errors were found, or the specified market is
//    not supported, the FlaggedTokens array is empty.
//
//    When Errors is non-empty, Type will be "ErrorResponse", otherwise it
//    is "SpellCheck"
//
type SpellCheckResponse struct {
	Type          string         `json:"_type"`
	FlaggedTokens []FlaggedToken `json:"flaggedTokens"`
	Errors        []Error        `json:"errors"`
}

// IsErrorResponse determines if the SpellCheckResponse indicates an error
func (scr *SpellCheckResponse) IsErrorResponse() bool {
	return scr.Type == ErrorResponseType
}

// IsSpellCheckResponse determines if the SpellCheckResponse was successful
func (scr *SpellCheckResponse) IsSpellCheckResponse() bool {
	return scr.Type == SpellCheckResponseType
}

// HasSuggestions determines if the SpellCheckResponse returned suggestions
func (scr *SpellCheckResponse) HasSuggestions() bool {
	return len(scr.FlaggedTokens) > 0
}

// TokenSuggestion is a suggested replacement entry for a token
//
//  Fields
//    Score      - A value that indicates the level of confidence that the
//      suggested correction is correct
//    Suggestion - The suggested word to replace the flagged word.
//
//  Notes
//    If the mode query parameter is set to 'spell', this field is set to 1.0
//
//    If the flagged word is a repeated word (see FlaggedToken.Type),
//    this string is empty.
//
type TokenSuggestion struct {
	Score      float64 `json:"score"`
	Suggestion string  `json:"suggestion"`
}

// FlaggedToken represents information about the word that is not spelled
// correctly or is grammatically incorrect
//
//  Fields
//    Offset      - The zero-based offset from the beginning of the text query
//      string to the word that was flagged
//    Suggestions - A list of words that correct the spelling or grammar error.
//      The list is in decreasing order of preference
//    Token       - The word in the text query string that is not spelled
//      correctly or is grammatically incorrect.
//    Type        - The type of error that caused the word to be flagged
//
//  Notes
//    The possible values for Type are:
//      'RepeatedToken' — The word was repeated in succession
//        (for example, the warm warm weather)
//      'UnknownToken'  — All other spelling or grammar errors
//
type FlaggedToken struct {
	Offset      int               `json:"offset"`
	Suggestions []TokenSuggestion `json:"suggestions"`
	Token       string            `json:"token"`
	Type        string            `json:"type"`
}

// IsRepeatedToken determines if the flagged token represents a repeated word
func (token FlaggedToken) IsRepeatedToken() bool {
	return token.Type == RepeatedTokenType
}

// IsUnknownToken determines if the flagged token represents a spelling or
// grammatical error
func (token FlaggedToken) IsUnknownToken() bool {
	return token.Type == UnknownTokenType
}

// Error defines an error that occurred
//
//  Fields
//    Code        - the error code that identifies the category of error
//    Message     - A description of the error
//    MoreDetails - A description that provides additional information about the error
//    Parameter   - The query parameter in the request that caused the error
//    SubCode     - The error code that identifies the error
//    Value       - The query parameter's value that was not valid
//
//  Notes
//    Code is an HTTP status code
//
type Error struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	MoreDetails string `json:"moreDetails"`
	Parameter   string `json:"parameter"`
	SubCode     string `json:"subCode"`
	Value       string `json:"value"`
}

func (err Error) Error() string {
	return fmt.Sprintf("%s: %s. Parameter=%s", err.Code, err.Message, err.Parameter)
}
