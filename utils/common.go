package utils

import (
	"strings"
)

type Argon2Config struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func ToString(s string) *string {
	s = strings.TrimSpace(s)

	if len(s) < 1 {
		return nil
	}

	return &s
}
