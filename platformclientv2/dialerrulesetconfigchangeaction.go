package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangeaction
type Dialerrulesetconfigchangeaction struct { 
	// VarType - Type of the action
	VarType *string `json:"type,omitempty"`


	// ActionTypeName - Identifier of the action
	ActionTypeName *string `json:"actionTypeName,omitempty"`


	// UpdateOption - Indicator of the type of update action (applicable only to certain types of actions)
	UpdateOption *string `json:"updateOption,omitempty"`


	// Properties - Map of key-value pairs pertinent to the action (different actions require different properties)
	Properties *map[string]string `json:"properties,omitempty"`

}

func (o *Dialerrulesetconfigchangeaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrulesetconfigchangeaction
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		ActionTypeName *string `json:"actionTypeName,omitempty"`
		
		UpdateOption *string `json:"updateOption,omitempty"`
		
		Properties *map[string]string `json:"properties,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		ActionTypeName: o.ActionTypeName,
		
		UpdateOption: o.UpdateOption,
		
		Properties: o.Properties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerrulesetconfigchangeaction) UnmarshalJSON(b []byte) error {
	var DialerrulesetconfigchangeactionMap map[string]interface{}
	err := json.Unmarshal(b, &DialerrulesetconfigchangeactionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DialerrulesetconfigchangeactionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if ActionTypeName, ok := DialerrulesetconfigchangeactionMap["actionTypeName"].(string); ok {
		o.ActionTypeName = &ActionTypeName
	}
	
	if UpdateOption, ok := DialerrulesetconfigchangeactionMap["updateOption"].(string); ok {
		o.UpdateOption = &UpdateOption
	}
	
	if Properties, ok := DialerrulesetconfigchangeactionMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangeaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
