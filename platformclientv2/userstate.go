package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userstate
type Userstate struct { 
	// State - User's current state.
	State *string `json:"state,omitempty"`


	// Version - Version of this user.
	Version *int `json:"version,omitempty"`


	// StateChangeReason - Reason for a change in the user's state.
	StateChangeReason *string `json:"stateChangeReason,omitempty"`


	// StateChangeDate - Date that the state was last changed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StateChangeDate *time.Time `json:"stateChangeDate,omitempty"`

}

func (o *Userstate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userstate
	
	StateChangeDate := new(string)
	if o.StateChangeDate != nil {
		
		*StateChangeDate = timeutil.Strftime(o.StateChangeDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StateChangeDate = nil
	}
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		StateChangeReason *string `json:"stateChangeReason,omitempty"`
		
		StateChangeDate *string `json:"stateChangeDate,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		Version: o.Version,
		
		StateChangeReason: o.StateChangeReason,
		
		StateChangeDate: StateChangeDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Userstate) UnmarshalJSON(b []byte) error {
	var UserstateMap map[string]interface{}
	err := json.Unmarshal(b, &UserstateMap)
	if err != nil {
		return err
	}
	
	if State, ok := UserstateMap["state"].(string); ok {
		o.State = &State
	}
	
	if Version, ok := UserstateMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if StateChangeReason, ok := UserstateMap["stateChangeReason"].(string); ok {
		o.StateChangeReason = &StateChangeReason
	}
	
	if stateChangeDateString, ok := UserstateMap["stateChangeDate"].(string); ok {
		StateChangeDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", stateChangeDateString)
		o.StateChangeDate = &StateChangeDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userstate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
