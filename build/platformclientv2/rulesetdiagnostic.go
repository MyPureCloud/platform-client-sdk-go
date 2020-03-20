package platformclientv2
import (
	"encoding/json"
)

// Rulesetdiagnostic
type Rulesetdiagnostic struct { 
	// RuleSet - A campaign rule set
	RuleSet *Domainentityref `json:"ruleSet,omitempty"`


	// Warnings - Diagnostic warnings for the rule set
	Warnings *[]string `json:"warnings,omitempty"`

}

// String returns a JSON representation of the model
func (o *Rulesetdiagnostic) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
