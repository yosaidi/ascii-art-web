package ascii

import "errors"

func ErrorHandler(text, banner string) error {
	if text == "" || banner == "" {
		return errors.New("text and banner fields cannot be empty")
	}
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		return errors.New("invalid banner type")
	}
	helpertext := []rune(text)
	if !AreStringValid(helpertext) {
		return errors.New("text contains invalid characters")
	}
	if len(helpertext) > 600 {
		return errors.New("text exceeds 600 characters limit")
	}
	return nil // No error
}
