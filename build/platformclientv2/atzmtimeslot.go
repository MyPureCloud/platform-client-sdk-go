package platformclientv2
import (
	"encoding/json"
)

// Atzmtimeslot
type Atzmtimeslot struct { 
	// EarliestCallableTime - The earliest time to dial a contact. Valid format is HH:mm
	EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`


	// LatestCallableTime - The latest time to dial a contact. Valid format is HH:mm
	LatestCallableTime *string `json:"latestCallableTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Atzmtimeslot) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
