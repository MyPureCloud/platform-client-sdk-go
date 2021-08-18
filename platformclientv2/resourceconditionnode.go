package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Resourceconditionnode
type Resourceconditionnode struct { 
	// VariableName
	VariableName *string `json:"variableName,omitempty"`


	// Conjunction
	Conjunction *string `json:"conjunction,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Operands
	Operands *[]Resourceconditionvalue `json:"operands,omitempty"`


	// Terms
	Terms *[]Resourceconditionnode `json:"terms,omitempty"`

}

func (u *Resourceconditionnode) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Resourceconditionnode

	

	return json.Marshal(&struct { 
		VariableName *string `json:"variableName,omitempty"`
		
		Conjunction *string `json:"conjunction,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Operands *[]Resourceconditionvalue `json:"operands,omitempty"`
		
		Terms *[]Resourceconditionnode `json:"terms,omitempty"`
		*Alias
	}{ 
		VariableName: u.VariableName,
		
		Conjunction: u.Conjunction,
		
		Operator: u.Operator,
		
		Operands: u.Operands,
		
		Terms: u.Terms,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Resourceconditionnode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
