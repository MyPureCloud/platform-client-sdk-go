package platformclientv2
import (
	"github.com/leekchan/timeutil"
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
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialerrulesetconfigchangeaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrulesetconfigchangeaction

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		ActionTypeName *string `json:"actionTypeName,omitempty"`
		
		UpdateOption *string `json:"updateOption,omitempty"`
		
		Properties *map[string]string `json:"properties,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		ActionTypeName: u.ActionTypeName,
		
		UpdateOption: u.UpdateOption,
		
		Properties: u.Properties,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangeaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
