package platformclientv2
import (
	"encoding/json"
)

// Posttextresponse
type Posttextresponse struct { 
	// BotState - The state of the bot after completion of the request
	BotState *string `json:"botState,omitempty"`


	// ReplyMessages - The list of messages to respond with, if any
	ReplyMessages *[]Posttextmessage `json:"replyMessages,omitempty"`


	// IntentName - The name of the intent the bot is either processing or has processed, this will be blank if no intent could be detected.
	IntentName *string `json:"intentName,omitempty"`


	// Slots - Data parameters detected and filled by the bot.
	Slots *map[string]string `json:"slots,omitempty"`


	// BotCorrelationId - The optional ID specified in the request
	BotCorrelationId *string `json:"botCorrelationId,omitempty"`


	// AmazonLex - Raw data response from AWS (if called)
	AmazonLex *map[string]interface{} `json:"amazonLex,omitempty"`


	// GoogleDialogFlow - Raw data response from Google Dialogflow (if called)
	GoogleDialogFlow *map[string]interface{} `json:"googleDialogFlow,omitempty"`


	// GenesysDialogEngine - Raw data response from Genesys' Dialogengine (if called)
	GenesysDialogEngine *map[string]interface{} `json:"genesysDialogEngine,omitempty"`

}

// String returns a JSON representation of the model
func (o *Posttextresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
