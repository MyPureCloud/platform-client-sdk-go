package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Parameter
type Parameter struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// ParameterType
	ParameterType *string `json:"parameterType,omitempty"`


	// Domain
	Domain *string `json:"domain,omitempty"`


	// Required
	Required *bool `json:"required,omitempty"`

}

func (o *Parameter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Parameter
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ParameterType *string `json:"parameterType,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		Required *bool `json:"required,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		ParameterType: o.ParameterType,
		
		Domain: o.Domain,
		
		Required: o.Required,
		Alias:    (*Alias)(o),
	})
}

func (o *Parameter) UnmarshalJSON(b []byte) error {
	var ParameterMap map[string]interface{}
	err := json.Unmarshal(b, &ParameterMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ParameterMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ParameterType, ok := ParameterMap["parameterType"].(string); ok {
		o.ParameterType = &ParameterType
	}
    
	if Domain, ok := ParameterMap["domain"].(string); ok {
		o.Domain = &Domain
	}
    
	if Required, ok := ParameterMap["required"].(bool); ok {
		o.Required = &Required
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Parameter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
