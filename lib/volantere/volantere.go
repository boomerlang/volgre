//
// This is the rule engine root package that will hold multiple rule engines
// defined by various business rules
//
// Implements the Factory design pattern
//
// Conventions: 
//
// 1. Every rule engine will live in a file with name of the rule engine
// in this package and in this directory.
//
// 2. Every rule engine will implement the interface VolanteRuleEngine.
//
// 3. Every rule engine will load its rules from the directory:
// $PATH_TO/volgre/rules/<rule_engine_name>/<version>.grl
//
// Author Bogdan Peta
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
	Refresh()
	Version() string
}

const (
	CreditCardRuleEngineT = 1
	BankNotExistentRuleEngine = 2
)

// registered rules engines
// TODO: load them from a config file
var RegisteredRuleEngines map[string]int = map[string]int{
	"credit_card": CreditCardRuleEngineT,
}

// Here will be pre-loadede the registered rules engines
// Singleton design pattern
var LoadedRuleEngines map[string]VolanteRuleEngine = map[string]VolanteRuleEngine{}

func GetRuleEngine(m int) (VolanteRuleEngine, error) {
	switch m {
	case CreditCardRuleEngineT:
		return new(CreditCardRuleEngine), nil
	default:
		return nil, errors.New(fmt.Sprintf("RuleEngine %d not available\n", m))
	}
}
