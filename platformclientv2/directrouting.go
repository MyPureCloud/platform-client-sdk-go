package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Directrouting
type Directrouting struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CallMediaSettings - Direct Routing Settings specific to Call media.
	CallMediaSettings *Directroutingmediasettings `json:"callMediaSettings,omitempty"`

	// EmailMediaSettings - Direct Routing Settings specific to Email media.
	EmailMediaSettings *Directroutingmediasettings `json:"emailMediaSettings,omitempty"`

	// MessageMediaSettings - Direct Routing Settings specific to Message media.
	MessageMediaSettings *Directroutingmediasettings `json:"messageMediaSettings,omitempty"`

	// BackupQueueId - ID of another queue to be used as the default backup if an agent does not have their Backup Settings configured. If not set, the current queue will be used as backup, but with Direct Routing criteria removed from the conversation.
	BackupQueueId *string `json:"backupQueueId,omitempty"`

	// WaitForAgent - Flag indicating if Direct Routing interactions should wait for Direct Routing agent or go immediately to selected backup.
	WaitForAgent *bool `json:"waitForAgent,omitempty"`

	// AgentWaitSeconds - Time (in seconds) that a Direct Routing interaction will wait for Direct Routing agent before going to selected backup. Valid range [60, 864000].
	AgentWaitSeconds *int `json:"agentWaitSeconds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Directrouting) SetField(field string, fieldValue interface{}) {
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

func (o Directrouting) MarshalJSON() ([]byte, error) {
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
	type Alias Directrouting
	
	return json.Marshal(&struct { 
		CallMediaSettings *Directroutingmediasettings `json:"callMediaSettings,omitempty"`
		
		EmailMediaSettings *Directroutingmediasettings `json:"emailMediaSettings,omitempty"`
		
		MessageMediaSettings *Directroutingmediasettings `json:"messageMediaSettings,omitempty"`
		
		BackupQueueId *string `json:"backupQueueId,omitempty"`
		
		WaitForAgent *bool `json:"waitForAgent,omitempty"`
		
		AgentWaitSeconds *int `json:"agentWaitSeconds,omitempty"`
		Alias
	}{ 
		CallMediaSettings: o.CallMediaSettings,
		
		EmailMediaSettings: o.EmailMediaSettings,
		
		MessageMediaSettings: o.MessageMediaSettings,
		
		BackupQueueId: o.BackupQueueId,
		
		WaitForAgent: o.WaitForAgent,
		
		AgentWaitSeconds: o.AgentWaitSeconds,
		Alias:    (Alias)(o),
	})
}

func (o *Directrouting) UnmarshalJSON(b []byte) error {
	var DirectroutingMap map[string]interface{}
	err := json.Unmarshal(b, &DirectroutingMap)
	if err != nil {
		return err
	}
	
	if CallMediaSettings, ok := DirectroutingMap["callMediaSettings"].(map[string]interface{}); ok {
		CallMediaSettingsString, _ := json.Marshal(CallMediaSettings)
		json.Unmarshal(CallMediaSettingsString, &o.CallMediaSettings)
	}
	
	if EmailMediaSettings, ok := DirectroutingMap["emailMediaSettings"].(map[string]interface{}); ok {
		EmailMediaSettingsString, _ := json.Marshal(EmailMediaSettings)
		json.Unmarshal(EmailMediaSettingsString, &o.EmailMediaSettings)
	}
	
	if MessageMediaSettings, ok := DirectroutingMap["messageMediaSettings"].(map[string]interface{}); ok {
		MessageMediaSettingsString, _ := json.Marshal(MessageMediaSettings)
		json.Unmarshal(MessageMediaSettingsString, &o.MessageMediaSettings)
	}
	
	if BackupQueueId, ok := DirectroutingMap["backupQueueId"].(string); ok {
		o.BackupQueueId = &BackupQueueId
	}
    
	if WaitForAgent, ok := DirectroutingMap["waitForAgent"].(bool); ok {
		o.WaitForAgent = &WaitForAgent
	}
    
	if AgentWaitSeconds, ok := DirectroutingMap["agentWaitSeconds"].(float64); ok {
		AgentWaitSecondsInt := int(AgentWaitSeconds)
		o.AgentWaitSeconds = &AgentWaitSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Directrouting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
