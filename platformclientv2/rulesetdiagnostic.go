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

func (u *Rulesetdiagnostic) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Rulesetdiagnostic

	

	return json.Marshal(&struct { 
		RuleSet *Domainentityref `json:"ruleSet,omitempty"`
		
		Warnings *[]string `json:"warnings,omitempty"`
		*Alias
	}{ 
		RuleSet: u.RuleSet,
		
		Warnings: u.Warnings,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Rulesetdiagnostic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
