package main

import (
	"context"
	"errors"
	"strings"
)

// option interface about string
type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

// option struct
type stringService struct{}

// func realize
func (str stringService) Uppercase(ctx context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (str stringService) Count(ctx context.Context, s string) int {
	return len(s)
}

var ErrEmpty = errors.New("Empty string")
