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
	Wrapup *Extendedwrapup `json:"wrapup,omitempty"`


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

func (o *Mediaparticipantrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediaparticipantrequest
	
	return json.Marshal(&struct { 
		Wrapup *Extendedwrapup `json:"wrapup,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Recording *bool `json:"recording,omitempty"`
		
		Muted *bool `json:"muted,omitempty"`
		
		Confined *bool `json:"confined,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		*Alias
	}{ 
		Wrapup: o.Wrapup,
		
		State: o.State,
		
		Recording: o.Recording,
		
		Muted: o.Muted,
		
		Confined: o.Confined,
		
		Held: o.Held,
		
		WrapupSkipped: o.WrapupSkipped,
		Alias:    (*Alias)(o),
	})
}

func (o *Mediaparticipantrequest) UnmarshalJSON(b []byte) error {
	var MediaparticipantrequestMap map[string]interface{}
	err := json.Unmarshal(b, &MediaparticipantrequestMap)
	if err != nil {
		return err
	}
	
	if Wrapup, ok := MediaparticipantrequestMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if State, ok := MediaparticipantrequestMap["state"].(string); ok {
		o.State = &State
	}
    
	if Recording, ok := MediaparticipantrequestMap["recording"].(bool); ok {
		o.Recording = &Recording
	}
    
	if Muted, ok := MediaparticipantrequestMap["muted"].(bool); ok {
		o.Muted = &Muted
	}
    
	if Confined, ok := MediaparticipantrequestMap["confined"].(bool); ok {
		o.Confined = &Confined
	}
    
	if Held, ok := MediaparticipantrequestMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if WrapupSkipped, ok := MediaparticipantrequestMap["wrapupSkipped"].(bool); ok {
		o.WrapupSkipped = &WrapupSkipped
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Mediaparticipantrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
