package actions

import "fmt"

type ErrorAction struct {
	messageId string
}

func NewErrorAction(id string) ActionInterface {
	var action = ErrorAction{messageId: ERR_UNKNOWN}

	if id != "" {
		action.messageId = id
	}

	return action
}

func (a ErrorAction) Resolve() (string, error) {
	return "", fmt.Errorf(a.messageId)
}
