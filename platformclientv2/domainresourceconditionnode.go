package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainresourceconditionnode
type Domainresourceconditionnode struct { 
	// VariableName
	VariableName *string `json:"variableName,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Operands
	Operands *[]Domainresourceconditionvalue `json:"operands,omitempty"`


	// Conjunction
	Conjunction *string `json:"conjunction,omitempty"`


	// Terms
	Terms *[]Domainresourceconditionnode `json:"terms,omitempty"`

}

func (u *Domainresourceconditionnode) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainresourceconditionnode

	

	return json.Marshal(&struct { 
		VariableName *string `json:"variableName,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Operands *[]Domainresourceconditionvalue `json:"operands,omitempty"`
		
		Conjunction *string `json:"conjunction,omitempty"`
		
		Terms *[]Domainresourceconditionnode `json:"terms,omitempty"`
		*Alias
	}{ 
		VariableName: u.VariableName,
		
		Operator: u.Operator,
		
		Operands: u.Operands,
		
		Conjunction: u.Conjunction,
		
		Terms: u.Terms,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Domainresourceconditionnode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
