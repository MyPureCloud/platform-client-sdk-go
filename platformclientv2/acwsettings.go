package platformclientv2
import (
	"encoding/json"
)

// Acwsettings
type Acwsettings struct { 
	// WrapupPrompt - This field controls how the UI prompts the agent for a wrapup.
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`


	// TimeoutMs - The amount of time the agent can stay in ACW (Min: 1 sec, Max: 1 day).  Can only be used when ACW is MANDATORY_TIMEOUT or MANDATORY_FORCED_TIMEOUT.
	TimeoutMs *int32 `json:"timeoutMs,omitempty"`

}

// String returns a JSON representation of the model
func (o *Acwsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
