//
// Controllers
//
// Author Bogdan Peta
//
package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/boomerlang/volgre/lib/volantere"
)

//
// Preloads all registered rule engines at application start up
//
func PreloadRuleEngines() {
	for engine_name, engine_id := range volantere.RegisteredRuleEngines {
		rule_engine, err := volantere.GetRuleEngine(engine_id)
		
		if err != nil {
			panic(err)
		}

		volantere.LoadedRuleEngines[engine_name] = rule_engine
		rule_path := "rules/" + engine_name + "_rules.grl"
		rule_engine.Init(string(rule_path))
	}
}

// RunRuleEngineHandler godoc
// @Summary Runs a rule engine
// @Description Runs a rule engine
// @Tags rule engine
// @Accept  json
// @Produce  json
// @Success 200
// @Router /run/engine/{rule_engine} [post]
func RunRuleEngineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	engine_name := vars["rule_engine"]

	w.Header().Set("Content-Type", "application/json")

	if rule_engine, ok := volantere.LoadedRuleEngines[engine_name]; ok {
	
		fn := func(data interface{}){json.NewDecoder(r.Body).Decode(data)}
		
		rule_engine.Load(fn)
		
		rule_engine.Run()

		fn1 := func(data interface{}) ([]byte, error) { return json.Marshal(data) }
		rez := rule_engine.Dump(fn1)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(rez)+"\n")
	} else {
		w.WriteHeader(404)	
		fmt.Fprintf(w, `{"error":"Rule Engine not found!"}`)
	}
}


// RefreshRuleEngineHandler godoc
// @Summary Refreshes a rule engine
// @Description Refreshes a rule engine with a new set of rules from disk
// @Tags rule engine
// @Accept  json
// @Produce  json
// @Success 200
// @Router /refresh/engine/{rule_engine} [post]
func RefreshRuleEngineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	engine_name := vars["rule_engine"]

	w.Header().Set("Content-Type", "application/json")

	if rule_engine, ok := volantere.LoadedRuleEngines[engine_name]; ok {
				
		rule_engine.Refresh()

		rez := `{"ok":"success"}`
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(rez)+"\n")
	} else {
		w.WriteHeader(404)
	
		fmt.Fprintf(w, `{"error":"Rule Engine not found!"}`)
	}
}

// VersionRuleEngineHandler godoc
// @Summary Version of a rule engine rule set
// @Description Refreshes a rule engine with a new set of rules from disk
// @Tags rule engine
// @Accept  json
// @Produce  json
// @Success 200
// @Router /version/engine/{rule_engine} [post]
func VersionRuleEngineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	engine_name := vars["rule_engine"]

	w.Header().Set("Content-Type", "application/json")

	if rule_engine, ok := volantere.LoadedRuleEngines[engine_name]; ok {
				
		ver := rule_engine.Version()

		rez := fmt.Sprintf(`{"version":"%s"}`, ver)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(rez)+"\n")
	} else {
		w.WriteHeader(404)
	
		fmt.Fprintf(w, `{"error":"Rule Engine not found!"}`)
	}
}