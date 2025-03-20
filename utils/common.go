package utils

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	defaultPassLen int = 10
	minPassLen     int = 10
	maxPassLen     int = 255
)

func ToString(s string) *string {
	s = strings.TrimSpace(s)

	if len(s) < 1 {
		return nil
	}

	return &s
}

func MiniumPasswordLenthg() int {
	pass, err := strconv.Atoi(os.Getenv("MIN_PASSWORD_LENGTH"))
	if err != nil {
		pass = defaultPassLen
	}

	if pass < minPassLen {
		pass = minPassLen
	}

	if pass > maxPassLen {
		pass = maxPassLen
	}

	return pass
}

func AddError(m fiber.Map, k string, v string) fiber.Map {
	slog.Warn(fmt.Sprintf("valor de m: %v, k:%v, v:%v", m, k, v))

	if _, ok := m[k]; !ok {
		m[k] = []string{v}
		slog.Warn(fmt.Sprintf("valor de m[k]: %v", m[k]))

	} else {
		m[k] = append(m[k].([]string), v)
		slog.Warn(fmt.Sprintf("valor de append m[k]: %v", m[k]))
	}

	slog.Warn(fmt.Sprintf("valor de return m: %v", m))

	return m
}
