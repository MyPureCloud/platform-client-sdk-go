package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Incomingmessagerequest - Incoming Message request
type Incomingmessagerequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// BotId - The unique id of the bot.
	BotId *string `json:"botId,omitempty"`

	// BotName - Name of the bot
	BotName *string `json:"botName,omitempty"`

	// BotVersion - The version of the bot.
	BotVersion *string `json:"botVersion,omitempty"`

	// IntegrationId - The Integration Id for the bot's configuration
	IntegrationId *string `json:"integrationId,omitempty"`

	// BotSessionId - The id of the session. This id will be used for an entire conversation with the bot (a series of back and forth between the bot until the bot has fulfilled its intent).
	BotSessionId *string `json:"botSessionId,omitempty"`

	// AutomateFlowExecId - Execution ID of the Automate Flow.
	AutomateFlowExecId *string `json:"automateFlowExecId,omitempty"`

	// ConversationId - Genesys conversation ID.
	ConversationId *string `json:"conversationId,omitempty"`

	// LanguageCode - Language code for the conversation (e.g., 'en-US').
	LanguageCode *string `json:"languageCode,omitempty"`

	// InputMessage - Message received from the user to send to the bot
	InputMessage *Inputmessage `json:"inputMessage,omitempty"`

	// MessagingPlatformType - Type of messaging platform being used
	MessagingPlatformType *string `json:"messagingPlatformType,omitempty"`

	// Channels - The channels this bot is utilizing.
	Channels *[]string `json:"channels,omitempty"`

	// BotSessionTimeout - Timeout duration for bot session in minutes.
	BotSessionTimeout *int `json:"botSessionTimeout,omitempty"`

	// Parameters - This is a map of string-string key, value pairs containing optional fields that can be passed to the bot for custom behavior, tracking, etc.
	Parameters *map[string]string `json:"parameters,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Incomingmessagerequest) SetField(field string, fieldValue interface{}) {
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

func (o Incomingmessagerequest) MarshalJSON() ([]byte, error) {
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
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
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
	type Alias Incomingmessagerequest
	
	return json.Marshal(&struct { 
		BotId *string `json:"botId,omitempty"`
		
		BotName *string `json:"botName,omitempty"`
		
		BotVersion *string `json:"botVersion,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		BotSessionId *string `json:"botSessionId,omitempty"`
		
		AutomateFlowExecId *string `json:"automateFlowExecId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		LanguageCode *string `json:"languageCode,omitempty"`
		
		InputMessage *Inputmessage `json:"inputMessage,omitempty"`
		
		MessagingPlatformType *string `json:"messagingPlatformType,omitempty"`
		
		Channels *[]string `json:"channels,omitempty"`
		
		BotSessionTimeout *int `json:"botSessionTimeout,omitempty"`
		
		Parameters *map[string]string `json:"parameters,omitempty"`
		Alias
	}{ 
		BotId: o.BotId,
		
		BotName: o.BotName,
		
		BotVersion: o.BotVersion,
		
		IntegrationId: o.IntegrationId,
		
		BotSessionId: o.BotSessionId,
		
		AutomateFlowExecId: o.AutomateFlowExecId,
		
		ConversationId: o.ConversationId,
		
		LanguageCode: o.LanguageCode,
		
		InputMessage: o.InputMessage,
		
		MessagingPlatformType: o.MessagingPlatformType,
		
		Channels: o.Channels,
		
		BotSessionTimeout: o.BotSessionTimeout,
		
		Parameters: o.Parameters,
		Alias:    (Alias)(o),
	})
}

func (o *Incomingmessagerequest) UnmarshalJSON(b []byte) error {
	var IncomingmessagerequestMap map[string]interface{}
	err := json.Unmarshal(b, &IncomingmessagerequestMap)
	if err != nil {
		return err
	}
	
	if BotId, ok := IncomingmessagerequestMap["botId"].(string); ok {
		o.BotId = &BotId
	}
    
	if BotName, ok := IncomingmessagerequestMap["botName"].(string); ok {
		o.BotName = &BotName
	}
    
	if BotVersion, ok := IncomingmessagerequestMap["botVersion"].(string); ok {
		o.BotVersion = &BotVersion
	}
    
	if IntegrationId, ok := IncomingmessagerequestMap["integrationId"].(string); ok {
		o.IntegrationId = &IntegrationId
	}
    
	if BotSessionId, ok := IncomingmessagerequestMap["botSessionId"].(string); ok {
		o.BotSessionId = &BotSessionId
	}
    
	if AutomateFlowExecId, ok := IncomingmessagerequestMap["automateFlowExecId"].(string); ok {
		o.AutomateFlowExecId = &AutomateFlowExecId
	}
    
	if ConversationId, ok := IncomingmessagerequestMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if LanguageCode, ok := IncomingmessagerequestMap["languageCode"].(string); ok {
		o.LanguageCode = &LanguageCode
	}
    
	if InputMessage, ok := IncomingmessagerequestMap["inputMessage"].(map[string]interface{}); ok {
		InputMessageString, _ := json.Marshal(InputMessage)
		json.Unmarshal(InputMessageString, &o.InputMessage)
	}
	
	if MessagingPlatformType, ok := IncomingmessagerequestMap["messagingPlatformType"].(string); ok {
		o.MessagingPlatformType = &MessagingPlatformType
	}
    
	if Channels, ok := IncomingmessagerequestMap["channels"].([]interface{}); ok {
		ChannelsString, _ := json.Marshal(Channels)
		json.Unmarshal(ChannelsString, &o.Channels)
	}
	
	if BotSessionTimeout, ok := IncomingmessagerequestMap["botSessionTimeout"].(float64); ok {
		BotSessionTimeoutInt := int(BotSessionTimeout)
		o.BotSessionTimeout = &BotSessionTimeoutInt
	}
	
	if Parameters, ok := IncomingmessagerequestMap["parameters"].(map[string]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Incomingmessagerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
