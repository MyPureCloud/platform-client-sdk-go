package platformclientv2
import (
	"encoding/json"
)

// Emailmediapolicyconditions
type Emailmediapolicyconditions struct { 
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

}

// String returns a JSON representation of the model
func (o *Emailmediapolicyconditions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
