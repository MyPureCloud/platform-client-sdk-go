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

func (o *Domainresourceconditionnode) MarshalJSON() ([]byte, error) {
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
		VariableName: o.VariableName,
		
		Operator: o.Operator,
		
		Operands: o.Operands,
		
		Conjunction: o.Conjunction,
		
		Terms: o.Terms,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainresourceconditionnode) UnmarshalJSON(b []byte) error {
	var DomainresourceconditionnodeMap map[string]interface{}
	err := json.Unmarshal(b, &DomainresourceconditionnodeMap)
	if err != nil {
		return err
	}
	
	if VariableName, ok := DomainresourceconditionnodeMap["variableName"].(string); ok {
		o.VariableName = &VariableName
	}
    
	if Operator, ok := DomainresourceconditionnodeMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Operands, ok := DomainresourceconditionnodeMap["operands"].([]interface{}); ok {
		OperandsString, _ := json.Marshal(Operands)
		json.Unmarshal(OperandsString, &o.Operands)
	}
	
	if Conjunction, ok := DomainresourceconditionnodeMap["conjunction"].(string); ok {
		o.Conjunction = &Conjunction
	}
    
	if Terms, ok := DomainresourceconditionnodeMap["terms"].([]interface{}); ok {
		TermsString, _ := json.Marshal(Terms)
		json.Unmarshal(TermsString, &o.Terms)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainresourceconditionnode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
