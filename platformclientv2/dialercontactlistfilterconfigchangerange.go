package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangerange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
