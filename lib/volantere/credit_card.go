//
// Rules engine for credit card
//
package volantere

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"

	"strconv"
	"time"

	"github.com/boomerlang/volgre/lib"
)

type ValidationResult struct {
	CardNumber bool `json:"cardNumber"`

	OwnerName bool `json:"ownerName"`

	ExpireMonth bool `json:"expireMonth"`

	ExpireYear bool `json:"expireYear"`

	SecurityCode bool `json:"securityCode"`

	Currency bool `json:"currency"`

	Amount bool `json:"amount"`
}

type CreditCard struct {
	Tenant interface{}	`json:"tenant"`
	MessageContent interface{}	`json:"messageContent"`
	MessageType  string `json:"messageType"`

	Result ValidationResult `json:"volgre:validationResults"`

	Rule_engine_execution_time string `json:"volgre:rule_engine_execution_time"`
}

func (cc *CreditCard) ValidateCardNumber() bool {
	ccn := cc.MessageContent.(map[string]interface{})["cardNumber"].(string)

	return validators.CheckCreditCard(ccn)
}

func (cc *CreditCard) ValidateExpMonth() bool {
	mc := cc.MessageContent.(map[string]interface{})

	now_dt := time.Now()

	expireMonth, err := strconv.Atoi(mc["expireMonth"].(string))
	if err != nil {
		return false
	}
	if expireMonth < 0 && expireMonth > 12 {
		return false
	}
	expireYear, err := strconv.Atoi(mc["expireYear"].(string))
	if err != nil {
		return false
	}

	now_mo, _ := strconv.Atoi(fmt.Sprintf("%2d", now_dt.Month()))
	now_yr, _ := strconv.Atoi(fmt.Sprintf("%d", now_dt.Year()))

	if expireYear < now_yr {
		return false
	} else if expireYear == now_yr {
		if expireMonth < now_mo {
			return false
		} 
		return true
	}
	return true
}

func (cc *CreditCard) ValidateExpYear() bool {
	mc := cc.MessageContent.(map[string]interface{})

	now_dt := time.Now()

	expireYear, err := strconv.Atoi(mc["expireYear"].(string))
	if err != nil {
		return false
	}

	now_yr, _ := strconv.Atoi(fmt.Sprintf("%d", now_dt.Year()))

	if expireYear < now_yr {
		return false
	}

	return true
}

func (cc *CreditCard) ValidateSecurityCode() bool {
	mc := cc.MessageContent.(map[string]interface{})

	sec_code := mc["securityCode"].(string)
	
	//if is integer 
	_, err := strconv.Atoi(sec_code)
	if err != nil {
		return false
	}

	return len(sec_code) == 3
}

func (cc *CreditCard) ValidateHolderName() bool {
	mc := cc.MessageContent.(map[string]interface{})

	holder := mc["ownerName"].(string)
	
	return validators.CheckHolderName(holder)
}

type CreditCardRuleEngine struct {
	fileRes pkg.Resource
	kb_lib *ast.KnowledgeLibrary
	dataContext ast.IDataContext
	rule_builder *builder.RuleBuilder

	data *CreditCard

	vr *ValidationResult
}

func (ccre *CreditCardRuleEngine) Load(fn func(interface{})) {
	ccre.data = new(CreditCard)
	fn(&ccre.data)
}

func (ccre *CreditCardRuleEngine) Dump(fn func(interface{}) ([]byte, error)) (rez []byte) {
	rez, err := fn(ccre.data)

	if err != nil {
		panic(err)
	}

	return rez
}

func (ccre *CreditCardRuleEngine) Init(rule_path string) {
	
	ccre.fileRes = pkg.NewFileResource(rule_path)
	ccre.kb_lib = ast.NewKnowledgeLibrary()
	ccre.rule_builder = builder.NewRuleBuilder(ccre.kb_lib)

	err := ccre.rule_builder.BuildRuleFromResource("Test", "0.1.1", ccre.fileRes)

	if err != nil {
		panic(err)
	}
}

func (ccre *CreditCardRuleEngine) Run() {
	start := time.Now()

	ccre.vr = &ValidationResult{}
	ccre.dataContext = ast.NewDataContext()
	err := ccre.dataContext.Add("CreditCard", ccre.data)
	if err != nil {
		panic(err)
	}

	err = ccre.dataContext.Add("ValidationResults", ccre.vr)
	if err != nil {
		panic(err)
	}

	kb_inst := ccre.kb_lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	reng := &engine.GruleEngine{MaxCycle: 100}
	err = reng.Execute(ccre.dataContext, kb_inst)
	
	if err != nil {
		panic(err)
	}

	ccre.data.Result = *ccre.vr
	elapsed := time.Since(start)
	ccre.data.Rule_engine_execution_time = fmt.Sprintf("%s", elapsed)
}
