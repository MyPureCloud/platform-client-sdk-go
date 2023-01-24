package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Posttextresponse
type Posttextresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Posttextresponse) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Posttextresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		Alias
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
		Alias:    (Alias)(o),
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
