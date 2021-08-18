package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// GenesysBotConnector - Raw data response from Genesys' BotConnector (if called)
	GenesysBotConnector *map[string]interface{} `json:"genesysBotConnector,omitempty"`


	// NuanceMixDlg - Raw data response from Nuance Mix Dlg (if called)
	NuanceMixDlg *map[string]interface{} `json:"nuanceMixDlg,omitempty"`

}

func (u *Posttextresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Posttextresponse

	

	return json.Marshal(&struct { 
		BotState *string `json:"botState,omitempty"`
		
		ReplyMessages *[]Posttextmessage `json:"replyMessages,omitempty"`
		
		IntentName *string `json:"intentName,omitempty"`
		
		Slots *map[string]string `json:"slots,omitempty"`
		
		BotCorrelationId *string `json:"botCorrelationId,omitempty"`
		
		AmazonLex *map[string]interface{} `json:"amazonLex,omitempty"`
		
		GoogleDialogFlow *map[string]interface{} `json:"googleDialogFlow,omitempty"`
		
		GenesysDialogEngine *map[string]interface{} `json:"genesysDialogEngine,omitempty"`
		
		GenesysBotConnector *map[string]interface{} `json:"genesysBotConnector,omitempty"`
		
		NuanceMixDlg *map[string]interface{} `json:"nuanceMixDlg,omitempty"`
		*Alias
	}{ 
		BotState: u.BotState,
		
		ReplyMessages: u.ReplyMessages,
		
		IntentName: u.IntentName,
		
		Slots: u.Slots,
		
		BotCorrelationId: u.BotCorrelationId,
		
		AmazonLex: u.AmazonLex,
		
		GoogleDialogFlow: u.GoogleDialogFlow,
		
		GenesysDialogEngine: u.GenesysDialogEngine,
		
		GenesysBotConnector: u.GenesysBotConnector,
		
		NuanceMixDlg: u.NuanceMixDlg,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Posttextresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
