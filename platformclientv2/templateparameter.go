package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Templateparameter
type Templateparameter struct { 
	// Id - Response substitution identifier
	Id *string `json:"id,omitempty"`


	// Value - Response substitution value
	Value *string `json:"value,omitempty"`

}

func (o *Templateparameter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Templateparameter
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Templateparameter) UnmarshalJSON(b []byte) error {
	var TemplateparameterMap map[string]interface{}
	err := json.Unmarshal(b, &TemplateparameterMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TemplateparameterMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Value, ok := TemplateparameterMap["value"].(string); ok {
		o.Value = &Value
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Templateparameter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
