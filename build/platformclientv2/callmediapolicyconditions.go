package platformclientv2
import (
	"encoding/json"
)

// Callmediapolicyconditions
type Callmediapolicyconditions struct { 
	// ForUsers
	ForUsers *[]User `json:"forUsers,omitempty"`


	// DateRanges
	DateRanges *[]string `json:"dateRanges,omitempty"`


	// ForQueues
	ForQueues *[]Queue `json:"forQueues,omitempty"`


	// WrapupCodes
	WrapupCodes *[]Wrapupcode `json:"wrapupCodes,omitempty"`


	// Languages
	Languages *[]Language `json:"languages,omitempty"`


	// TimeAllowed
	TimeAllowed *Timeallowed `json:"timeAllowed,omitempty"`


	// Directions
	Directions *[]string `json:"directions,omitempty"`


	// Duration
	Duration *Durationcondition `json:"duration,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callmediapolicyconditions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
