package actions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ncaak/pifiabot/dice"
)

type CoCAction struct {
	command string
}

func (a CoCAction) Resolve() (string, error) {
	var params []string = strings.SplitN(a.command, " ", 3)
	if len(params) < 2 {
		return "", fmt.Errorf(ERR_COC_WRONG_ARGUMENT)
	}

	skillValue, err := strconv.Atoi(params[1])
	if err != nil {
		return "", fmt.Errorf(ERR_COC_WRONG_ARGUMENT)
	}

	if skillValue < 1 || skillValue > 100 {
		return "", fmt.Errorf(ERR_COC_SKILL_LIMITS)
	}

	return a.result(skillValue, dice.RollD100()), nil
}

func (a CoCAction) result(skillValue, roll int) string {

	switch {
	case roll > skillValue:
		return fmt.Sprintf("Fallo [%d]", roll)
	case roll <= skillValue/5:
		return fmt.Sprintf("Éxito (crítico) [%d]", roll)
	case roll <= skillValue/2:
		return fmt.Sprintf("Éxito (difícil) [%d]", roll)
	case roll <= skillValue:
		return fmt.Sprintf("Éxito (normal) [%d]", roll)
	}

	return ""
}
