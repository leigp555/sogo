package test

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	str := "adas[error] failed to initialize database, got error %v\n"
	t.Log(strings.HasPrefix(str, "[error]"))
}
