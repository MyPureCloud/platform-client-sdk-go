package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Settings
type Settings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CommunicationBasedACW - Communication Based ACW
	CommunicationBasedACW *bool `json:"communicationBasedACW,omitempty"`

	// IncludeNonAgentConversationSummary - Display communication summary
	IncludeNonAgentConversationSummary *bool `json:"includeNonAgentConversationSummary,omitempty"`

	// AllowCallbackQueueSelection - Allow Callback Queue Selection
	AllowCallbackQueueSelection *bool `json:"allowCallbackQueueSelection,omitempty"`

	// CompleteAcwWhenAgentTransitionsOffline - Complete ACW When Agent Transitions Offline
	CompleteAcwWhenAgentTransitionsOffline *bool `json:"completeAcwWhenAgentTransitionsOffline,omitempty"`

	// TotalActiveCallback - Exclude the 'interacting' duration from the handle calculations of callbacks
	TotalActiveCallback *bool `json:"totalActiveCallback,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Settings) SetField(field string, fieldValue interface{}) {
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

func (o Settings) MarshalJSON() ([]byte, error) {
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
	type Alias Settings
	
	return json.Marshal(&struct { 
		CommunicationBasedACW *bool `json:"communicationBasedACW,omitempty"`
		
		IncludeNonAgentConversationSummary *bool `json:"includeNonAgentConversationSummary,omitempty"`
		
		AllowCallbackQueueSelection *bool `json:"allowCallbackQueueSelection,omitempty"`
		
		CompleteAcwWhenAgentTransitionsOffline *bool `json:"completeAcwWhenAgentTransitionsOffline,omitempty"`
		
		TotalActiveCallback *bool `json:"totalActiveCallback,omitempty"`
		Alias
	}{ 
		CommunicationBasedACW: o.CommunicationBasedACW,
		
		IncludeNonAgentConversationSummary: o.IncludeNonAgentConversationSummary,
		
		AllowCallbackQueueSelection: o.AllowCallbackQueueSelection,
		
		CompleteAcwWhenAgentTransitionsOffline: o.CompleteAcwWhenAgentTransitionsOffline,
		
		TotalActiveCallback: o.TotalActiveCallback,
		Alias:    (Alias)(o),
	})
}

func (o *Settings) UnmarshalJSON(b []byte) error {
	var SettingsMap map[string]interface{}
	err := json.Unmarshal(b, &SettingsMap)
	if err != nil {
		return err
	}
	
	if CommunicationBasedACW, ok := SettingsMap["communicationBasedACW"].(bool); ok {
		o.CommunicationBasedACW = &CommunicationBasedACW
	}
    
	if IncludeNonAgentConversationSummary, ok := SettingsMap["includeNonAgentConversationSummary"].(bool); ok {
		o.IncludeNonAgentConversationSummary = &IncludeNonAgentConversationSummary
	}
    
	if AllowCallbackQueueSelection, ok := SettingsMap["allowCallbackQueueSelection"].(bool); ok {
		o.AllowCallbackQueueSelection = &AllowCallbackQueueSelection
	}
    
	if CompleteAcwWhenAgentTransitionsOffline, ok := SettingsMap["completeAcwWhenAgentTransitionsOffline"].(bool); ok {
		o.CompleteAcwWhenAgentTransitionsOffline = &CompleteAcwWhenAgentTransitionsOffline
	}
    
	if TotalActiveCallback, ok := SettingsMap["totalActiveCallback"].(bool); ok {
		o.TotalActiveCallback = &TotalActiveCallback
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Settings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
