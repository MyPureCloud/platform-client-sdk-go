package platformclientv2
import (
	"encoding/json"
)

// Googledialogflowcustomsettings
type Googledialogflowcustomsettings struct { 
	// Environment - If set this environment will be used to initiate the dialogflow bot, otherwise the default configuration will be used.  See https://cloud.google.com/dialogflow/docs/agents-versions
	Environment *string `json:"environment,omitempty"`

}

// String returns a JSON representation of the model
func (o *Googledialogflowcustomsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
