package actions

type UnknownAction struct{}

func (a UnknownAction) Resolve() (string, error) {
	return "El comando no se ha reconocido. Consulta la ayuda.", nil
}
