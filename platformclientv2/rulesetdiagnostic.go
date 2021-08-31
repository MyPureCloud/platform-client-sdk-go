package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Rulesetdiagnostic
type Rulesetdiagnostic struct { 
	// RuleSet - A campaign rule set
	RuleSet *Domainentityref `json:"ruleSet,omitempty"`


	// Warnings - Diagnostic warnings for the rule set
	Warnings *[]string `json:"warnings,omitempty"`

}

func (o *Rulesetdiagnostic) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Rulesetdiagnostic
	
	return json.Marshal(&struct { 
		RuleSet *Domainentityref `json:"ruleSet,omitempty"`
		
		Warnings *[]string `json:"warnings,omitempty"`
		*Alias
	}{ 
		RuleSet: o.RuleSet,
		
		Warnings: o.Warnings,
		Alias:    (*Alias)(o),
	})
}

func (o *Rulesetdiagnostic) UnmarshalJSON(b []byte) error {
	var RulesetdiagnosticMap map[string]interface{}
	err := json.Unmarshal(b, &RulesetdiagnosticMap)
	if err != nil {
		return err
	}
	
	if RuleSet, ok := RulesetdiagnosticMap["ruleSet"].(map[string]interface{}); ok {
		RuleSetString, _ := json.Marshal(RuleSet)
		json.Unmarshal(RuleSetString, &o.RuleSet)
	}
	
	if Warnings, ok := RulesetdiagnosticMap["warnings"].([]interface{}); ok {
		WarningsString, _ := json.Marshal(Warnings)
		json.Unmarshal(WarningsString, &o.Warnings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Rulesetdiagnostic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
