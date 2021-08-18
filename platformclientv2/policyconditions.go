package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Policyconditions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Policyconditions

	

	return json.Marshal(&struct { 
		ForUsers *[]User `json:"forUsers,omitempty"`
		
		Directions *[]string `json:"directions,omitempty"`
		
		DateRanges *[]string `json:"dateRanges,omitempty"`
		
		MediaTypes *[]string `json:"mediaTypes,omitempty"`
		
		ForQueues *[]Queue `json:"forQueues,omitempty"`
		
		Duration *Durationcondition `json:"duration,omitempty"`
		
		WrapupCodes *[]Wrapupcode `json:"wrapupCodes,omitempty"`
		
		TimeAllowed *Timeallowed `json:"timeAllowed,omitempty"`
		*Alias
	}{ 
		ForUsers: u.ForUsers,
		
		Directions: u.Directions,
		
		DateRanges: u.DateRanges,
		
		MediaTypes: u.MediaTypes,
		
		ForQueues: u.ForQueues,
		
		Duration: u.Duration,
		
		WrapupCodes: u.WrapupCodes,
		
		TimeAllowed: u.TimeAllowed,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Policyconditions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
