package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Callmediapolicyconditions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callmediapolicyconditions

	

	return json.Marshal(&struct { 
		ForUsers *[]User `json:"forUsers,omitempty"`
		
		DateRanges *[]string `json:"dateRanges,omitempty"`
		
		ForQueues *[]Queue `json:"forQueues,omitempty"`
		
		WrapupCodes *[]Wrapupcode `json:"wrapupCodes,omitempty"`
		
		Languages *[]Language `json:"languages,omitempty"`
		
		TimeAllowed *Timeallowed `json:"timeAllowed,omitempty"`
		
		Directions *[]string `json:"directions,omitempty"`
		
		Duration *Durationcondition `json:"duration,omitempty"`
		*Alias
	}{ 
		ForUsers: u.ForUsers,
		
		DateRanges: u.DateRanges,
		
		ForQueues: u.ForQueues,
		
		WrapupCodes: u.WrapupCodes,
		
		Languages: u.Languages,
		
		TimeAllowed: u.TimeAllowed,
		
		Directions: u.Directions,
		
		Duration: u.Duration,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Callmediapolicyconditions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
