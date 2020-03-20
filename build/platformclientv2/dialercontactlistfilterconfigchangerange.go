package platformclientv2
import (
	"encoding/json"
)

// Dialercontactlistfilterconfigchangerange
type Dialercontactlistfilterconfigchangerange struct { 
	// Min
	Min *string `json:"min,omitempty"`


	// Max
	Max *string `json:"max,omitempty"`


	// MinInclusive
	MinInclusive *bool `json:"minInclusive,omitempty"`


	// MaxInclusive
	MaxInclusive *bool `json:"maxInclusive,omitempty"`


	// InSet
	InSet *[]string `json:"inSet,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangerange) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
