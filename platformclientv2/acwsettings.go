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

func (o *Acwsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Acwsettings
	
	return json.Marshal(&struct { 
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		TimeoutMs *int `json:"timeoutMs,omitempty"`
		*Alias
	}{ 
		WrapupPrompt: o.WrapupPrompt,
		
		TimeoutMs: o.TimeoutMs,
		Alias:    (*Alias)(o),
	})
}

func (o *Acwsettings) UnmarshalJSON(b []byte) error {
	var AcwsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &AcwsettingsMap)
	if err != nil {
		return err
	}
	
	if WrapupPrompt, ok := AcwsettingsMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
    
	if TimeoutMs, ok := AcwsettingsMap["timeoutMs"].(float64); ok {
		TimeoutMsInt := int(TimeoutMs)
		o.TimeoutMs = &TimeoutMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Acwsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
