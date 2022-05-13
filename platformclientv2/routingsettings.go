package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingsettings
type Routingsettings struct { 
	// ResetAgentScoreOnPresenceChange - Reset agent score when agent presence changes from off-queue to on-queue
	ResetAgentScoreOnPresenceChange *bool `json:"resetAgentScoreOnPresenceChange,omitempty"`

}

func (o *Routingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingsettings
	
	return json.Marshal(&struct { 
		ResetAgentScoreOnPresenceChange *bool `json:"resetAgentScoreOnPresenceChange,omitempty"`
		*Alias
	}{ 
		ResetAgentScoreOnPresenceChange: o.ResetAgentScoreOnPresenceChange,
		Alias:    (*Alias)(o),
	})
}

func (o *Routingsettings) UnmarshalJSON(b []byte) error {
	var RoutingsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &RoutingsettingsMap)
	if err != nil {
		return err
	}
	
	if ResetAgentScoreOnPresenceChange, ok := RoutingsettingsMap["resetAgentScoreOnPresenceChange"].(bool); ok {
		o.ResetAgentScoreOnPresenceChange = &ResetAgentScoreOnPresenceChange
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Routingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
