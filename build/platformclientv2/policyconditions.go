package platformclientv2
import (
	"encoding/json"
)

// Policyconditions
type Policyconditions struct { 
	// ForUsers
	ForUsers *[]User `json:"forUsers,omitempty"`


	// Directions
	Directions *[]string `json:"directions,omitempty"`


	// DateRanges
	DateRanges *[]string `json:"dateRanges,omitempty"`


	// MediaTypes
	MediaTypes *[]string `json:"mediaTypes,omitempty"`


	// ForQueues
	ForQueues *[]Queue `json:"forQueues,omitempty"`


	// Duration
	Duration *Durationcondition `json:"duration,omitempty"`


	// WrapupCodes
	WrapupCodes *[]Wrapupcode `json:"wrapupCodes,omitempty"`


	// TimeAllowed
	TimeAllowed *Timeallowed `json:"timeAllowed,omitempty"`

}

// String returns a JSON representation of the model
func (o *Policyconditions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
