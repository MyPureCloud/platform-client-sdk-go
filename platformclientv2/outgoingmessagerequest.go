package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outgoingmessagerequest - Outgoing Message request
type Outgoingmessagerequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// BotId - The unique id of the bot.
	BotId *string `json:"botId,omitempty"`

	// BotVersion - The version of the bot.
	BotVersion *string `json:"botVersion,omitempty"`

	// BotSessionId - The id of the session. This id will be used for an entire conversation with the bot (a series of back and forth between the bot until the bot has fulfilled its intent).
	BotSessionId *string `json:"botSessionId,omitempty"`

	// BotState - The state of the bot reported
	BotState *string `json:"botState,omitempty"`

	// LanguageCode - The language used for this message. EG 'en-us' or 'es', etc; These language codes are W3C language identification tags (ISO 639-1 for the language name and ISO 3166 for the country code).
	LanguageCode *string `json:"languageCode,omitempty"`

	// ReplyMessages - This is a list of messages to send back to the user, this field can be null or an empty list.
	ReplyMessages *[]Replymessage `json:"replyMessages,omitempty"`

	// Intent - The name of the intent the bot is either processing or has processed, this will be blank if no intent could be detected.
	Intent *string `json:"intent,omitempty"`

	// Confidence - A value between 0 and 1.0 denoting the confidence of the discovered intent (if found) this is optional and if left null genesys assumes a confidence of 1.0 on success and 0 on fail.
	Confidence *float64 `json:"confidence,omitempty"`

	// ErrorInfo - If the botState is Failed the bot can add this error object with more details about the error.
	ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`

	// Parameters - This is a map of string-string key, value pairs containing optional fields that can be passed from the bot for custom behavior, tracking, etc, which can be used by the flow.
	Parameters *map[string]string `json:"parameters,omitempty"`

	// Entities - A set of entity values that go along with the intent.
	Entities *[]Botentityvalue `json:"entities,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outgoingmessagerequest) SetField(field string, fieldValue interface{}) {
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

func (o Outgoingmessagerequest) MarshalJSON() ([]byte, error) {
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
	type Alias Outgoingmessagerequest
	
	return json.Marshal(&struct { 
		BotId *string `json:"botId,omitempty"`
		
		BotVersion *string `json:"botVersion,omitempty"`
		
		BotSessionId *string `json:"botSessionId,omitempty"`
		
		BotState *string `json:"botState,omitempty"`
		
		LanguageCode *string `json:"languageCode,omitempty"`
		
		ReplyMessages *[]Replymessage `json:"replyMessages,omitempty"`
		
		Intent *string `json:"intent,omitempty"`
		
		Confidence *float64 `json:"confidence,omitempty"`
		
		ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`
		
		Parameters *map[string]string `json:"parameters,omitempty"`
		
		Entities *[]Botentityvalue `json:"entities,omitempty"`
		Alias
	}{ 
		BotId: o.BotId,
		
		BotVersion: o.BotVersion,
		
		BotSessionId: o.BotSessionId,
		
		BotState: o.BotState,
		
		LanguageCode: o.LanguageCode,
		
		ReplyMessages: o.ReplyMessages,
		
		Intent: o.Intent,
		
		Confidence: o.Confidence,
		
		ErrorInfo: o.ErrorInfo,
		
		Parameters: o.Parameters,
		
		Entities: o.Entities,
		Alias:    (Alias)(o),
	})
}

func (o *Outgoingmessagerequest) UnmarshalJSON(b []byte) error {
	var OutgoingmessagerequestMap map[string]interface{}
	err := json.Unmarshal(b, &OutgoingmessagerequestMap)
	if err != nil {
		return err
	}
	
	if BotId, ok := OutgoingmessagerequestMap["botId"].(string); ok {
		o.BotId = &BotId
	}
    
	if BotVersion, ok := OutgoingmessagerequestMap["botVersion"].(string); ok {
		o.BotVersion = &BotVersion
	}
    
	if BotSessionId, ok := OutgoingmessagerequestMap["botSessionId"].(string); ok {
		o.BotSessionId = &BotSessionId
	}
    
	if BotState, ok := OutgoingmessagerequestMap["botState"].(string); ok {
		o.BotState = &BotState
	}
    
	if LanguageCode, ok := OutgoingmessagerequestMap["languageCode"].(string); ok {
		o.LanguageCode = &LanguageCode
	}
    
	if ReplyMessages, ok := OutgoingmessagerequestMap["replyMessages"].([]interface{}); ok {
		ReplyMessagesString, _ := json.Marshal(ReplyMessages)
		json.Unmarshal(ReplyMessagesString, &o.ReplyMessages)
	}
	
	if Intent, ok := OutgoingmessagerequestMap["intent"].(string); ok {
		o.Intent = &Intent
	}
    
	if Confidence, ok := OutgoingmessagerequestMap["confidence"].(float64); ok {
		o.Confidence = &Confidence
	}
    
	if ErrorInfo, ok := OutgoingmessagerequestMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Parameters, ok := OutgoingmessagerequestMap["parameters"].(map[string]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	
	if Entities, ok := OutgoingmessagerequestMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outgoingmessagerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
