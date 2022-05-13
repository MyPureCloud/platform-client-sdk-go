package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialogflowintent
type Dialogflowintent struct { 
	// Name - The intent name
	Name *string `json:"name,omitempty"`


	// Parameters - An object mapping parameter names to Parameter objects
	Parameters *map[string]Dialogflowparameter `json:"parameters,omitempty"`

}

func (o *Dialogflowintent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialogflowintent
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Parameters *map[string]Dialogflowparameter `json:"parameters,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Parameters: o.Parameters,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialogflowintent) UnmarshalJSON(b []byte) error {
	var DialogflowintentMap map[string]interface{}
	err := json.Unmarshal(b, &DialogflowintentMap)
	if err != nil {
		return err
	}
	
	if Name, ok := DialogflowintentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Parameters, ok := DialogflowintentMap["parameters"].(map[string]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialogflowintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
