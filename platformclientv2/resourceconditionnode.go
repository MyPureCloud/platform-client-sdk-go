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

func (o *Resourceconditionnode) MarshalJSON() ([]byte, error) {
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
		VariableName: o.VariableName,
		
		Conjunction: o.Conjunction,
		
		Operator: o.Operator,
		
		Operands: o.Operands,
		
		Terms: o.Terms,
		Alias:    (*Alias)(o),
	})
}

func (o *Resourceconditionnode) UnmarshalJSON(b []byte) error {
	var ResourceconditionnodeMap map[string]interface{}
	err := json.Unmarshal(b, &ResourceconditionnodeMap)
	if err != nil {
		return err
	}
	
	if VariableName, ok := ResourceconditionnodeMap["variableName"].(string); ok {
		o.VariableName = &VariableName
	}
    
	if Conjunction, ok := ResourceconditionnodeMap["conjunction"].(string); ok {
		o.Conjunction = &Conjunction
	}
    
	if Operator, ok := ResourceconditionnodeMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Operands, ok := ResourceconditionnodeMap["operands"].([]interface{}); ok {
		OperandsString, _ := json.Marshal(Operands)
		json.Unmarshal(OperandsString, &o.Operands)
	}
	
	if Terms, ok := ResourceconditionnodeMap["terms"].([]interface{}); ok {
		TermsString, _ := json.Marshal(Terms)
		json.Unmarshal(TermsString, &o.Terms)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Resourceconditionnode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
