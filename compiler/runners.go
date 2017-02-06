package compiler

type Runners struct {
	runners map[string]RunnerFuncs
}

type RunnerFuncs struct {
	Validate func(interface{}) error
	Preamble func() string
	Command  func(interface{}) string
}

func runners() (r Runners) {
	r.runners = make(map[string]RunnerFuncs)
	r.runners["pkg"] = RunnerFuncs{
		Validate: PKGValidate,
		Preamble: PKGPreamble,
		Command:  PKGCommand,
	}

	r.runners["service"] = RunnerFuncs{
		Validate: ServiceValidate,
		Preamble: ServicePreamble,
		Command:  ServiceCommand,
	}

	return
}
