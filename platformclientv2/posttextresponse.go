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

func (o *Posttextresponse) MarshalJSON() ([]byte, error) {
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
		BotState: o.BotState,
		
		ReplyMessages: o.ReplyMessages,
		
		IntentName: o.IntentName,
		
		Slots: o.Slots,
		
		BotCorrelationId: o.BotCorrelationId,
		
		AmazonLex: o.AmazonLex,
		
		GoogleDialogFlow: o.GoogleDialogFlow,
		
		GenesysDialogEngine: o.GenesysDialogEngine,
		
		GenesysBotConnector: o.GenesysBotConnector,
		
		NuanceMixDlg: o.NuanceMixDlg,
		Alias:    (*Alias)(o),
	})
}

func (o *Posttextresponse) UnmarshalJSON(b []byte) error {
	var PosttextresponseMap map[string]interface{}
	err := json.Unmarshal(b, &PosttextresponseMap)
	if err != nil {
		return err
	}
	
	if BotState, ok := PosttextresponseMap["botState"].(string); ok {
		o.BotState = &BotState
	}
    
	if ReplyMessages, ok := PosttextresponseMap["replyMessages"].([]interface{}); ok {
		ReplyMessagesString, _ := json.Marshal(ReplyMessages)
		json.Unmarshal(ReplyMessagesString, &o.ReplyMessages)
	}
	
	if IntentName, ok := PosttextresponseMap["intentName"].(string); ok {
		o.IntentName = &IntentName
	}
    
	if Slots, ok := PosttextresponseMap["slots"].(map[string]interface{}); ok {
		SlotsString, _ := json.Marshal(Slots)
		json.Unmarshal(SlotsString, &o.Slots)
	}
	
	if BotCorrelationId, ok := PosttextresponseMap["botCorrelationId"].(string); ok {
		o.BotCorrelationId = &BotCorrelationId
	}
    
	if AmazonLex, ok := PosttextresponseMap["amazonLex"].(map[string]interface{}); ok {
		AmazonLexString, _ := json.Marshal(AmazonLex)
		json.Unmarshal(AmazonLexString, &o.AmazonLex)
	}
	
	if GoogleDialogFlow, ok := PosttextresponseMap["googleDialogFlow"].(map[string]interface{}); ok {
		GoogleDialogFlowString, _ := json.Marshal(GoogleDialogFlow)
		json.Unmarshal(GoogleDialogFlowString, &o.GoogleDialogFlow)
	}
	
	if GenesysDialogEngine, ok := PosttextresponseMap["genesysDialogEngine"].(map[string]interface{}); ok {
		GenesysDialogEngineString, _ := json.Marshal(GenesysDialogEngine)
		json.Unmarshal(GenesysDialogEngineString, &o.GenesysDialogEngine)
	}
	
	if GenesysBotConnector, ok := PosttextresponseMap["genesysBotConnector"].(map[string]interface{}); ok {
		GenesysBotConnectorString, _ := json.Marshal(GenesysBotConnector)
		json.Unmarshal(GenesysBotConnectorString, &o.GenesysBotConnector)
	}
	
	if NuanceMixDlg, ok := PosttextresponseMap["nuanceMixDlg"].(map[string]interface{}); ok {
		NuanceMixDlgString, _ := json.Marshal(NuanceMixDlg)
		json.Unmarshal(NuanceMixDlgString, &o.NuanceMixDlg)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Posttextresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
