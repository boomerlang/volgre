//
// This is the interface with grule-rule-engine.
//
// Represents the core of every rule engine derived from VolanteRuleEngine
//
// Author Bogdan Peta
//
package volantere

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"

	"strconv"
)

type RuleEngine struct {
	fileRes pkg.Resource
	kb_lib *ast.KnowledgeLibrary
	dataContext ast.IDataContext
	rule_builder *builder.RuleBuilder

	rule_path string
	crt_version string
	new_nersion string

	// rule engine name
	name string
}

func (re *RuleEngine) init(rule_path string) {
	re.rule_path = rule_path
	re.crt_version = "v0"
	re.fileRes = pkg.NewFileResource(rule_path)
	re.kb_lib = ast.NewKnowledgeLibrary()
	re.rule_builder = builder.NewRuleBuilder(re.kb_lib)

	err := re.rule_builder.BuildRuleFromResource(re.name, re.crt_version, re.fileRes)

	if err != nil {
		panic(err)
	}

	re.dataContext = ast.NewDataContext()
}

func (re *RuleEngine) refresh() {
	
	i, _ := strconv.Atoi(re.crt_version[1:])

	re.new_nersion = re.crt_version[:1] + strconv.Itoa(i+1)
	
	re.fileRes = pkg.NewFileResource(re.rule_path)
	err := re.rule_builder.BuildRuleFromResource(re.name, re.new_nersion, re.fileRes)

	if err != nil {
		panic(err)
	}

	re.crt_version = re.new_nersion
	fmt.Println(fmt.Sprintf("Rules refreshed for engine %s with version %s", re.name, re.crt_version))
	re.dataContext = ast.NewDataContext()
}

func (re *RuleEngine) add_datacontext(name string, data interface{}) {
	err := re.dataContext.Add(name, data)
	if err != nil {
		panic(err)
	}
}

func (re *RuleEngine) run() {
	kb_inst := re.kb_lib.NewKnowledgeBaseInstance(re.name, re.crt_version)
	reng := &engine.GruleEngine{MaxCycle: 3001}
	err := reng.Execute(re.dataContext, kb_inst)
	
	if err != nil {
		panic(err)
	}
}
