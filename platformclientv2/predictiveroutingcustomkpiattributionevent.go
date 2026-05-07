package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictiveroutingcustomkpiattributionevent
type Predictiveroutingcustomkpiattributionevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventId - A unique (UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`

	// EventDateTime - A timestamp as epoch representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`

	// ExternalContactId - The UUID of the external contact associated with this event
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// ConversationId - The UUID of the conversation associated with this event
	ConversationId *string `json:"conversationId,omitempty"`

	// AgentId - The UUID of the agent associated with this event
	AgentId *string `json:"agentId,omitempty"`

	// KpiId - The UUID of the KPI associated with this event
	KpiId *string `json:"kpiId,omitempty"`

	// AssociatedValue - The value associated with this outcome attribution
	AssociatedValue *float64 `json:"associatedValue,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Predictiveroutingcustomkpiattributionevent) SetField(field string, fieldValue interface{}) {
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

func (o Predictiveroutingcustomkpiattributionevent) MarshalJSON() ([]byte, error) {
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
	type Alias Predictiveroutingcustomkpiattributionevent
	
	EventDateTime := new(string)
	if o.EventDateTime != nil {
		
		*EventDateTime = timeutil.Strftime(o.EventDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventDateTime = nil
	}
	
	return json.Marshal(&struct { 
		EventId *string `json:"eventId,omitempty"`
		
		EventDateTime *string `json:"eventDateTime,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		AgentId *string `json:"agentId,omitempty"`
		
		KpiId *string `json:"kpiId,omitempty"`
		
		AssociatedValue *float64 `json:"associatedValue,omitempty"`
		Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ExternalContactId: o.ExternalContactId,
		
		ConversationId: o.ConversationId,
		
		AgentId: o.AgentId,
		
		KpiId: o.KpiId,
		
		AssociatedValue: o.AssociatedValue,
		Alias:    (Alias)(o),
	})
}

func (o *Predictiveroutingcustomkpiattributionevent) UnmarshalJSON(b []byte) error {
	var PredictiveroutingcustomkpiattributioneventMap map[string]interface{}
	err := json.Unmarshal(b, &PredictiveroutingcustomkpiattributioneventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := PredictiveroutingcustomkpiattributioneventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := PredictiveroutingcustomkpiattributioneventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ExternalContactId, ok := PredictiveroutingcustomkpiattributioneventMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if ConversationId, ok := PredictiveroutingcustomkpiattributioneventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if AgentId, ok := PredictiveroutingcustomkpiattributioneventMap["agentId"].(string); ok {
		o.AgentId = &AgentId
	}
    
	if KpiId, ok := PredictiveroutingcustomkpiattributioneventMap["kpiId"].(string); ok {
		o.KpiId = &KpiId
	}
    
	if AssociatedValue, ok := PredictiveroutingcustomkpiattributioneventMap["associatedValue"].(float64); ok {
		o.AssociatedValue = &AssociatedValue
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Predictiveroutingcustomkpiattributionevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
