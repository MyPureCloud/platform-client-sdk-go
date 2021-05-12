package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Domainresourceconditionnode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
