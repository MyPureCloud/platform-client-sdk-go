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

func (u *Routingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingsettings

	

	return json.Marshal(&struct { 
		ResetAgentScoreOnPresenceChange *bool `json:"resetAgentScoreOnPresenceChange,omitempty"`
		*Alias
	}{ 
		ResetAgentScoreOnPresenceChange: u.ResetAgentScoreOnPresenceChange,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Routingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
