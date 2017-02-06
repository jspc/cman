package compiler

import (
	"testing"
)

func TestCompile(t *testing.T) {
	compiled, err := Compile("./fixtures/web-server.yml")
	if err != nil {
		t.Errorf("Compile() error: received %q", err)
	}

	script, err := compiled.Generate()
	if err != nil {
		t.Errorf("Compiled.Generate() error: received %q", err)
	}
}
