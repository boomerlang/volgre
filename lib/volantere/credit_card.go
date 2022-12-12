//
// Rule engine for credit card
//
// Author Bogdan Peta
//
package volantere

import (
	"fmt"

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

type CreditCardRuleEngine struct {
	data *CreditCard

	vr *ValidationResult

	re *RuleEngine
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
	ccre.re = new(RuleEngine)
	ccre.re.name = "CreditCard"
	ccre.re.init(rule_path)
}

func (ccre *CreditCardRuleEngine) Version() string {
	return ccre.re.crt_version
}

func (ccre *CreditCardRuleEngine) Refresh() {
	ccre.re.refresh()
}

func (ccre *CreditCardRuleEngine) Run() {
	start := time.Now()

	ccre.vr = &ValidationResult{}

	ccre.re.add_datacontext("CreditCard", ccre.data)
	
	ccre.re.add_datacontext("ValidationResults", ccre.vr)
	
	ccre.re.run()

	ccre.data.Result = *ccre.vr
	elapsed := time.Since(start)
	ccre.data.Rule_engine_execution_time = fmt.Sprintf("%s", elapsed)
}


// Rules data structure
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

