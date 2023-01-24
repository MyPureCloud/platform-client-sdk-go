package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userpresenceevent
type Userpresenceevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`

	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`

	// UserId - The User ID of the user associated with this UserPresence
	UserId *string `json:"userId,omitempty"`

	// SourceId - The id (V4 UUID) of the presence source being updated
	SourceId *string `json:"sourceId,omitempty"`

	// PresenceDefinitionId - The id (UUID) of the presence definition that the user presence is associated with
	PresenceDefinitionId *string `json:"presenceDefinitionId,omitempty"`

	// Message - The message associated with the presence
	Message *string `json:"message,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userpresenceevent) SetField(field string, fieldValue interface{}) {
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

func (o Userpresenceevent) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EventDateTime", }
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
	type Alias Userpresenceevent
	
	EventDateTime := new(string)
	if o.EventDateTime != nil {
		
		*EventDateTime = timeutil.Strftime(o.EventDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventDateTime = nil
	}
	
	return json.Marshal(&struct { 
		EventId *string `json:"eventId,omitempty"`
		
		EventDateTime *string `json:"eventDateTime,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		SourceId *string `json:"sourceId,omitempty"`
		
		PresenceDefinitionId *string `json:"presenceDefinitionId,omitempty"`
		
		Message *string `json:"message,omitempty"`
		Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		UserId: o.UserId,
		
		SourceId: o.SourceId,
		
		PresenceDefinitionId: o.PresenceDefinitionId,
		
		Message: o.Message,
		Alias:    (Alias)(o),
	})
}

func (o *Userpresenceevent) UnmarshalJSON(b []byte) error {
	var UserpresenceeventMap map[string]interface{}
	err := json.Unmarshal(b, &UserpresenceeventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := UserpresenceeventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := UserpresenceeventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if UserId, ok := UserpresenceeventMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if SourceId, ok := UserpresenceeventMap["sourceId"].(string); ok {
		o.SourceId = &SourceId
	}
    
	if PresenceDefinitionId, ok := UserpresenceeventMap["presenceDefinitionId"].(string); ok {
		o.PresenceDefinitionId = &PresenceDefinitionId
	}
    
	if Message, ok := UserpresenceeventMap["message"].(string); ok {
		o.Message = &Message
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userpresenceevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
