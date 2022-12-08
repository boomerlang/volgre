//
// This is a rule engine root package that will hold multiple rule engines
// defined by various business rules
//
package volantere

import (
	"errors"
	"fmt"
)

type VolanteRuleEngine interface {
	Init(rule_path string)
	Load(f func(interface{}))
	Dump(f func(interface{}) ([]byte, error)) []byte
	Run()
}

const (
	CreditCardRuleEngineT = 1
	BankNotExistentRuleEngine = 2
)

var RegisteredRuleEngines map[string]int = map[string]int{
	"credit_card": CreditCardRuleEngineT,
}

var LoadedRuleEngines map[string]VolanteRuleEngine = map[string]VolanteRuleEngine{}


func GetRuleEngine(m int) (VolanteRuleEngine, error) {
	switch m {
	case CreditCardRuleEngineT:
		return new(CreditCardRuleEngine), nil
	default:
		return nil, errors.New(fmt.Sprintf("RuleEngine %d not available\n", m))
	}
}

