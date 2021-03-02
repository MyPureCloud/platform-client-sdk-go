package platformclientv2
import (
	"encoding/json"
)

// Contactlistfilterrange
type Contactlistfilterrange struct { 
	// Min - The minimum value of the range. Required for the operator BETWEEN.
	Min *string `json:"min,omitempty"`


	// Max - The maximum value of the range. Required for the operator BETWEEN.
	Max *string `json:"max,omitempty"`


	// MinInclusive - Whether or not to include the minimum in the range.
	MinInclusive *bool `json:"minInclusive,omitempty"`


	// MaxInclusive - Whether or not to include the maximum in the range.
	MaxInclusive *bool `json:"maxInclusive,omitempty"`


	// InSet - A set of values that the contact data should be in. Required for the IN operator.
	InSet *[]string `json:"inSet,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contactlistfilterrange) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
