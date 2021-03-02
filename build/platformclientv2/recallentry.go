package platformclientv2
import (
	"encoding/json"
)

// Recallentry
type Recallentry struct { 
	// NbrAttempts
	NbrAttempts *int `json:"nbrAttempts,omitempty"`


	// MinutesBetweenAttempts
	MinutesBetweenAttempts *int `json:"minutesBetweenAttempts,omitempty"`

}

// String returns a JSON representation of the model
func (o *Recallentry) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
