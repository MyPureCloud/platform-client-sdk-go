package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationappsettings - Conversation settings that handles chats within the messenger
type Conversationappsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - The toggle to enable or disable conversations
	Enabled *bool `json:"enabled,omitempty"`

	// ShowAgentTypingIndicator - The toggle to enable or disable typing indicator for messenger
	ShowAgentTypingIndicator *bool `json:"showAgentTypingIndicator,omitempty"`

	// ShowUserTypingIndicator - The toggle to enable or disable typing indicator for messenger
	ShowUserTypingIndicator *bool `json:"showUserTypingIndicator,omitempty"`

	// AutoStart - The auto start for the messenger conversation
	AutoStart *Autostart `json:"autoStart,omitempty"`

	// Markdown - The markdown for the messenger app
	Markdown *Markdown `json:"markdown,omitempty"`

	// ConversationDisconnect - The conversation disconnect settings for the messenger app
	ConversationDisconnect *Conversationdisconnectsettings `json:"conversationDisconnect,omitempty"`

	// ConversationClear - The conversation clear settings for the messenger app
	ConversationClear *Conversationclearsettings `json:"conversationClear,omitempty"`

	// Humanize - The humanize conversations settings for the messenger app
	Humanize *Humanize `json:"humanize,omitempty"`

	// Notifications - The notification settings for messenger apps
	Notifications *Notificationssettings `json:"notifications,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationappsettings) SetField(field string, fieldValue interface{}) {
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

func (o Conversationappsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationappsettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		ShowAgentTypingIndicator *bool `json:"showAgentTypingIndicator,omitempty"`
		
		ShowUserTypingIndicator *bool `json:"showUserTypingIndicator,omitempty"`
		
		AutoStart *Autostart `json:"autoStart,omitempty"`
		
		Markdown *Markdown `json:"markdown,omitempty"`
		
		ConversationDisconnect *Conversationdisconnectsettings `json:"conversationDisconnect,omitempty"`
		
		ConversationClear *Conversationclearsettings `json:"conversationClear,omitempty"`
		
		Humanize *Humanize `json:"humanize,omitempty"`
		
		Notifications *Notificationssettings `json:"notifications,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		ShowAgentTypingIndicator: o.ShowAgentTypingIndicator,
		
		ShowUserTypingIndicator: o.ShowUserTypingIndicator,
		
		AutoStart: o.AutoStart,
		
		Markdown: o.Markdown,
		
		ConversationDisconnect: o.ConversationDisconnect,
		
		ConversationClear: o.ConversationClear,
		
		Humanize: o.Humanize,
		
		Notifications: o.Notifications,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationappsettings) UnmarshalJSON(b []byte) error {
	var ConversationappsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationappsettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := ConversationappsettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if ShowAgentTypingIndicator, ok := ConversationappsettingsMap["showAgentTypingIndicator"].(bool); ok {
		o.ShowAgentTypingIndicator = &ShowAgentTypingIndicator
	}
    
	if ShowUserTypingIndicator, ok := ConversationappsettingsMap["showUserTypingIndicator"].(bool); ok {
		o.ShowUserTypingIndicator = &ShowUserTypingIndicator
	}
    
	if AutoStart, ok := ConversationappsettingsMap["autoStart"].(map[string]interface{}); ok {
		AutoStartString, _ := json.Marshal(AutoStart)
		json.Unmarshal(AutoStartString, &o.AutoStart)
	}
	
	if Markdown, ok := ConversationappsettingsMap["markdown"].(map[string]interface{}); ok {
		MarkdownString, _ := json.Marshal(Markdown)
		json.Unmarshal(MarkdownString, &o.Markdown)
	}
	
	if ConversationDisconnect, ok := ConversationappsettingsMap["conversationDisconnect"].(map[string]interface{}); ok {
		ConversationDisconnectString, _ := json.Marshal(ConversationDisconnect)
		json.Unmarshal(ConversationDisconnectString, &o.ConversationDisconnect)
	}
	
	if ConversationClear, ok := ConversationappsettingsMap["conversationClear"].(map[string]interface{}); ok {
		ConversationClearString, _ := json.Marshal(ConversationClear)
		json.Unmarshal(ConversationClearString, &o.ConversationClear)
	}
	
	if Humanize, ok := ConversationappsettingsMap["humanize"].(map[string]interface{}); ok {
		HumanizeString, _ := json.Marshal(Humanize)
		json.Unmarshal(HumanizeString, &o.Humanize)
	}
	
	if Notifications, ok := ConversationappsettingsMap["notifications"].(map[string]interface{}); ok {
		NotificationsString, _ := json.Marshal(Notifications)
		json.Unmarshal(NotificationsString, &o.Notifications)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationappsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
