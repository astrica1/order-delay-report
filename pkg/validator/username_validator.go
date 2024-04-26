package validator

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	USERNAME_MIN_LENGTH = 5
	USERNAME_MAX_LENGTH = 50
)

func ValidateUsername(username string) error {
	if len(username) < USERNAME_MIN_LENGTH {
		return fmt.Errorf("username must be at least %d characters", USERNAME_MIN_LENGTH)
	}

	if len(username) > USERNAME_MAX_LENGTH && USERNAME_MAX_LENGTH > USERNAME_MIN_LENGTH {
		return fmt.Errorf("username must be at most %d characters", USERNAME_MAX_LENGTH)
	}

	username = strings.ToLower(username)

	pattern := "^[a-z0-9_]+$"
	regex := regexp.MustCompile(pattern)
	isMatched := regex.MatchString(username)
	if !isMatched {
		return fmt.Errorf("username '%s' is Invalid", username)
	}

	return nil
}
