package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowlaunchrequest - Settings for launching an instance of a bot flow.
type Textbotflowlaunchrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Flow - Specifies which Bot Flow to launch.
	Flow *Textbotflow `json:"flow,omitempty"`

	// ExternalSessionId - The ID of the external session that is associated with the bot flow.
	ExternalSessionId *string `json:"externalSessionId,omitempty"`

	// ConversationId - A conversation ID to associate with the bot flow, if available.
	ConversationId *string `json:"conversationId,omitempty"`

	// InputData - Input values to the flow. Valid values are defined by the flow's input JSON schema.
	InputData *Textbotinputoutputdata `json:"inputData,omitempty"`

	// Channel - Channel information relevant to the bot flow.
	Channel *Textbotchannel `json:"channel,omitempty"`

	// Language - The language that the bot will use in the session. Validated against list of supported languages and if the value is omitted or is invalid, the default language will be used.
	Language *string `json:"language,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Textbotflowlaunchrequest) SetField(field string, fieldValue interface{}) {
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

func (o Textbotflowlaunchrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Textbotflowlaunchrequest
	
	return json.Marshal(&struct { 
		Flow *Textbotflow `json:"flow,omitempty"`
		
		ExternalSessionId *string `json:"externalSessionId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		InputData *Textbotinputoutputdata `json:"inputData,omitempty"`
		
		Channel *Textbotchannel `json:"channel,omitempty"`
		
		Language *string `json:"language,omitempty"`
		Alias
	}{ 
		Flow: o.Flow,
		
		ExternalSessionId: o.ExternalSessionId,
		
		ConversationId: o.ConversationId,
		
		InputData: o.InputData,
		
		Channel: o.Channel,
		
		Language: o.Language,
		Alias:    (Alias)(o),
	})
}

func (o *Textbotflowlaunchrequest) UnmarshalJSON(b []byte) error {
	var TextbotflowlaunchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotflowlaunchrequestMap)
	if err != nil {
		return err
	}
	
	if Flow, ok := TextbotflowlaunchrequestMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if ExternalSessionId, ok := TextbotflowlaunchrequestMap["externalSessionId"].(string); ok {
		o.ExternalSessionId = &ExternalSessionId
	}
    
	if ConversationId, ok := TextbotflowlaunchrequestMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if InputData, ok := TextbotflowlaunchrequestMap["inputData"].(map[string]interface{}); ok {
		InputDataString, _ := json.Marshal(InputData)
		json.Unmarshal(InputDataString, &o.InputData)
	}
	
	if Channel, ok := TextbotflowlaunchrequestMap["channel"].(map[string]interface{}); ok {
		ChannelString, _ := json.Marshal(Channel)
		json.Unmarshal(ChannelString, &o.Channel)
	}
	
	if Language, ok := TextbotflowlaunchrequestMap["language"].(string); ok {
		o.Language = &Language
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotflowlaunchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
