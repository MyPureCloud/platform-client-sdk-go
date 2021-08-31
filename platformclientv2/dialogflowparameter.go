package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialogflowparameter
type Dialogflowparameter struct { 
	// Name - The parameter name
	Name *string `json:"name,omitempty"`


	// VarType - The parameter type
	VarType *string `json:"type,omitempty"`

}

func (o *Dialogflowparameter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialogflowparameter
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialogflowparameter) UnmarshalJSON(b []byte) error {
	var DialogflowparameterMap map[string]interface{}
	err := json.Unmarshal(b, &DialogflowparameterMap)
	if err != nil {
		return err
	}
	
	if Name, ok := DialogflowparameterMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if VarType, ok := DialogflowparameterMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialogflowparameter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
