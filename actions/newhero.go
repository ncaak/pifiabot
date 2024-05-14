package actions

type NewHeroAction struct {
	command string
}

func (a NewHeroAction) Resolve() (string, error) {
	// TODO: Currently only DnD chars are returned
	return RepeatAction{command: "/repite 6 4d6-L"}.Resolve()
}
