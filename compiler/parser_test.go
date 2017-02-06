package compiler

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	_, err := ParseFile("./fixtures/web-server.yml")
	if err != nil {
		t.Errorf("%q", err)
	}
}
