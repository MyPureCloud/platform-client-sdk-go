package platformclientv2
import (
	"encoding/json"
)

// Googledialogflowcustomsettings
type Googledialogflowcustomsettings struct { 
	// Environment - If set this environment will be used to initiate the dialogflow bot, otherwise the default configuration will be used.  See https://cloud.google.com/dialogflow/docs/agents-versions
	Environment *string `json:"environment,omitempty"`


	// EventName - If set this eventName will be used to initiate the dialogflow bot rather than language processing on the input text.  See https://cloud.google.com/dialogflow/es/docs/events-overview
	EventName *string `json:"eventName,omitempty"`


	// WebhookQueryParameters - Parameters passed to the fulfillment webhook of the bot (if any).
	WebhookQueryParameters *map[string]string `json:"webhookQueryParameters,omitempty"`


	// EventInputParameters - Parameters passed to the event input of the bot.
	EventInputParameters *map[string]string `json:"eventInputParameters,omitempty"`

}

// String returns a JSON representation of the model
func (o *Googledialogflowcustomsettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
