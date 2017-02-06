package compiler

import (
	"gopkg.in/yaml.v2"
)

type Compiled struct {
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name   string      `yaml:"name"`
	Runner string      `yaml:"runner"`
	Args   interface{} `yaml:"args"`
}

func Compile(p string) (c Compiled, err error) {
	var data string
	if data, err = ParseFile(p); err != nil {
		return
	}

	err = yaml.Unmarshal([]byte(data), &c)
	return
}

func (c Compiled) Generate() (script string, err error) {
	script += scriptPreamble()

	rData := runners().runners
	for _, r := range c.distinctRunners() {
		script += rData[r].Preamble()
	}

	for _, s := range c.Steps {
		if err = rData[s.Runner].Validate(s.Args); err != nil {
			return
		}

		script += rData[s.Runner].Command(s.Args)
	}

	return
}

func (c Compiled) distinctRunners() (r []string) {
	for _, s := range c.Steps {
		if !contains(r, s.Runner) {
			r = append(r, s.Runner)
		}
	}
	return
}

func scriptPreamble() string {
	return `#!/usr/bin/env bash
# cman runner

set -axe

export E_BAD_DISTRO=1
export E_BAD_CMD=2

export PACKAGE_MANAGER_UPDATE=0   # Cause a $package_manager update on first pkg install

err() {
  local msg="${1}"
  local code="${2}"
  echo msg >&2
  exit "${code}"
}

`
}

func contains(s []string, k string) bool {
	for _, i := range s {
		if i == k {
			return true
		}
	}
	return false
}
