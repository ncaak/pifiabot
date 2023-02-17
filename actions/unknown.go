package actions

import (
	"fmt"
)

type UnknownAction struct{}

func (a UnknownAction) Resolve() (string, error) {
	return "", fmt.Errorf(MSG_UNKNOWN_ACTION)
}
