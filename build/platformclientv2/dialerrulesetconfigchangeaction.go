package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangeaction
type Dialerrulesetconfigchangeaction struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// ActionTypeName
	ActionTypeName *string `json:"actionTypeName,omitempty"`


	// UpdateOption
	UpdateOption *string `json:"updateOption,omitempty"`


	// Properties
	Properties *map[string]string `json:"properties,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangeaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
