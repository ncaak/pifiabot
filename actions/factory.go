package actions

type Action interface {
	Resolve() (string, error)
}

func Factory(command string) Action {
	switch {
	// TODO: Define commands as they are coded
	default:
		return UnknownAction{}
	}
}
