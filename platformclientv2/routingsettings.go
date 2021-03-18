package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Routingsettings
type Routingsettings struct { 
	// ResetAgentScoreOnPresenceChange - Reset agent score when agent presence changes from off-queue to on-queue
	ResetAgentScoreOnPresenceChange *bool `json:"resetAgentScoreOnPresenceChange,omitempty"`

}

// String returns a JSON representation of the model
func (o *Routingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
