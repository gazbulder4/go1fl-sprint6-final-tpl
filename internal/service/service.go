package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(s string) bool {
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' && r != '\n' && r != '\r' {
			return false
		}
	}
	return true
}

func Convert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("строка пуста")
	}

	if isMorse(trimmed) {
		return morse.ToText(trimmed), nil
	}

	return morse.ToMorse(trimmed), nil
}
