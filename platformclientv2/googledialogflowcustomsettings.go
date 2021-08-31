package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Googledialogflowcustomsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Googledialogflowcustomsettings
	
	return json.Marshal(&struct { 
		Environment *string `json:"environment,omitempty"`
		
		EventName *string `json:"eventName,omitempty"`
		
		WebhookQueryParameters *map[string]string `json:"webhookQueryParameters,omitempty"`
		
		EventInputParameters *map[string]string `json:"eventInputParameters,omitempty"`
		*Alias
	}{ 
		Environment: o.Environment,
		
		EventName: o.EventName,
		
		WebhookQueryParameters: o.WebhookQueryParameters,
		
		EventInputParameters: o.EventInputParameters,
		Alias:    (*Alias)(o),
	})
}

func (o *Googledialogflowcustomsettings) UnmarshalJSON(b []byte) error {
	var GoogledialogflowcustomsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &GoogledialogflowcustomsettingsMap)
	if err != nil {
		return err
	}
	
	if Environment, ok := GoogledialogflowcustomsettingsMap["environment"].(string); ok {
		o.Environment = &Environment
	}
	
	if EventName, ok := GoogledialogflowcustomsettingsMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
	
	if WebhookQueryParameters, ok := GoogledialogflowcustomsettingsMap["webhookQueryParameters"].(map[string]interface{}); ok {
		WebhookQueryParametersString, _ := json.Marshal(WebhookQueryParameters)
		json.Unmarshal(WebhookQueryParametersString, &o.WebhookQueryParameters)
	}
	
	if EventInputParameters, ok := GoogledialogflowcustomsettingsMap["eventInputParameters"].(map[string]interface{}); ok {
		EventInputParametersString, _ := json.Marshal(EventInputParameters)
		json.Unmarshal(EventInputParametersString, &o.EventInputParameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Googledialogflowcustomsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
