package bingSpellCheck

import "bytes"

// BuildAutoCorrectedText updates text to reflect the corrections in response
func BuildAutoCorrectedText(text string, response *SpellCheckResponse) (result string, err error) {
	// guard against a panic
	defer func() {
		if r := recover(); r != nil {
			result = ""
			err = r.(Error)
		}
	}()

	// the API response was an error
	if response.IsErrorResponse() {
		return "", response.Errors[0]
	}

	// Nothing to correct
	if !response.HasSuggestions() {
		return text, nil
	}

	var buf bytes.Buffer

	srcIndex := 0

	// range thru each suggestion and replace with 1st available suggestion
	// Note: Assumes flagged tokens are ordered by offset
	for _, token := range response.FlaggedTokens {
		// write anything up to this suggestion
		buf.WriteString(text[srcIndex:token.Offset])

		// move srcIndex past the token being corrected
		srcIndex = token.Offset + len(token.Token)

		if token.IsRepeatedToken() {
			// assume there is a ' ' character that needs to be skipped as well
			srcIndex++
		} else if token.IsUnknownToken() {
			buf.WriteString(token.Suggestions[0].Suggestion)
		} else {
			// unsupported/unknown type, skip it
		}
	}

	// trailing part of original text (if any)
	if srcIndex < len(text) {
		// flush anything remaining
		buf.WriteString(text[srcIndex:])
	}

	return buf.String(), nil
}
