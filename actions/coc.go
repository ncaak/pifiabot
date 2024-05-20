package actions

type CoCAction struct {
	command string
}

func (a CoCAction) Resolve() (string, error) {
	return a.command, nil
}
