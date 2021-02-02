package platformclientv2
import (
	"encoding/json"
)

// Dialerattemptlimitsconfigchangerecallentry
type Dialerattemptlimitsconfigchangerecallentry struct { 
	// NbrAttempts
	NbrAttempts *int `json:"nbrAttempts,omitempty"`


	// MinutesBetweenAttempts
	MinutesBetweenAttempts *int `json:"minutesBetweenAttempts,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerattemptlimitsconfigchangerecallentry) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
