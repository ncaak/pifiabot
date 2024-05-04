package actions

type NewHeroAction struct {
	command string
}

func (a NewHeroAction) Resolve() (string, error) {
	return "test", nil
}
