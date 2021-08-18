package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediaparticipantrequest
type Mediaparticipantrequest struct { 
	// Wrapup - Wrap-up to assign to this participant.
	Wrapup *Wrapup `json:"wrapup,omitempty"`


	// State - The state to update to set for this participant's communications.  Possible values are: 'connected' and 'disconnected'.
	State *string `json:"state,omitempty"`


	// Recording - True to enable recording of this participant, otherwise false to disable recording.
	Recording *bool `json:"recording,omitempty"`


	// Muted - True to mute this conversation participant.
	Muted *bool `json:"muted,omitempty"`


	// Confined - True to confine this conversation participant.  Should only be used for ad-hoc conferences
	Confined *bool `json:"confined,omitempty"`


	// Held - True to hold this conversation participant.
	Held *bool `json:"held,omitempty"`


	// WrapupSkipped - True to skip wrap-up for this participant.
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`

}

func (u *Mediaparticipantrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediaparticipantrequest

	

	return json.Marshal(&struct { 
		Wrapup *Wrapup `json:"wrapup,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Recording *bool `json:"recording,omitempty"`
		
		Muted *bool `json:"muted,omitempty"`
		
		Confined *bool `json:"confined,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		*Alias
	}{ 
		Wrapup: u.Wrapup,
		
		State: u.State,
		
		Recording: u.Recording,
		
		Muted: u.Muted,
		
		Confined: u.Confined,
		
		Held: u.Held,
		
		WrapupSkipped: u.WrapupSkipped,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Mediaparticipantrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
