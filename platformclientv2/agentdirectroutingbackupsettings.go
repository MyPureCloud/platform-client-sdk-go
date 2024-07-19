package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentdirectroutingbackupsettings
type Agentdirectroutingbackupsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QueueId - ID of queue to be used as backup. If queueId and userId are both specified, queue behaves as secondary backup.
	QueueId *string `json:"queueId,omitempty"`

	// UserId - ID of user to be used as backup. If queueId and userId are both specified, user behaves as primary backup.
	UserId *string `json:"userId,omitempty"`

	// WaitForAgent - Flag indicating if Direct Routing interactions should wait for Direct Routing agent or go immediately to selected backup.
	WaitForAgent *bool `json:"waitForAgent,omitempty"`

	// AgentWaitSeconds - Time (in seconds) that a Direct Routing interaction will wait for Direct Routing agent before going to selected backup. Valid range [60, 864000].
	AgentWaitSeconds *int `json:"agentWaitSeconds,omitempty"`

	// BackedUpUsers - Set of users that this user is a backup for.
	BackedUpUsers *[]string `json:"backedUpUsers,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentdirectroutingbackupsettings) SetField(field string, fieldValue interface{}) {
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

func (o Agentdirectroutingbackupsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Agentdirectroutingbackupsettings
	
	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		WaitForAgent *bool `json:"waitForAgent,omitempty"`
		
		AgentWaitSeconds *int `json:"agentWaitSeconds,omitempty"`
		
		BackedUpUsers *[]string `json:"backedUpUsers,omitempty"`
		Alias
	}{ 
		QueueId: o.QueueId,
		
		UserId: o.UserId,
		
		WaitForAgent: o.WaitForAgent,
		
		AgentWaitSeconds: o.AgentWaitSeconds,
		
		BackedUpUsers: o.BackedUpUsers,
		Alias:    (Alias)(o),
	})
}

func (o *Agentdirectroutingbackupsettings) UnmarshalJSON(b []byte) error {
	var AgentdirectroutingbackupsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &AgentdirectroutingbackupsettingsMap)
	if err != nil {
		return err
	}
	
	if QueueId, ok := AgentdirectroutingbackupsettingsMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if UserId, ok := AgentdirectroutingbackupsettingsMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if WaitForAgent, ok := AgentdirectroutingbackupsettingsMap["waitForAgent"].(bool); ok {
		o.WaitForAgent = &WaitForAgent
	}
    
	if AgentWaitSeconds, ok := AgentdirectroutingbackupsettingsMap["agentWaitSeconds"].(float64); ok {
		AgentWaitSecondsInt := int(AgentWaitSeconds)
		o.AgentWaitSeconds = &AgentWaitSecondsInt
	}
	
	if BackedUpUsers, ok := AgentdirectroutingbackupsettingsMap["backedUpUsers"].([]interface{}); ok {
		BackedUpUsersString, _ := json.Marshal(BackedUpUsers)
		json.Unmarshal(BackedUpUsersString, &o.BackedUpUsers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentdirectroutingbackupsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
