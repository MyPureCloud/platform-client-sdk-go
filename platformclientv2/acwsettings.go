package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Acwsettings
type Acwsettings struct { 
	// WrapupPrompt - This field controls how the UI prompts the agent for a wrapup.
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`


	// TimeoutMs - The amount of time the agent can stay in ACW (Min: 1 sec, Max: 1 day).  Can only be used when ACW is MANDATORY_TIMEOUT or MANDATORY_FORCED_TIMEOUT.
	TimeoutMs *int `json:"timeoutMs,omitempty"`

}

func (u *Acwsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Acwsettings

	

	return json.Marshal(&struct { 
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		TimeoutMs *int `json:"timeoutMs,omitempty"`
		*Alias
	}{ 
		WrapupPrompt: u.WrapupPrompt,
		
		TimeoutMs: u.TimeoutMs,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Acwsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
