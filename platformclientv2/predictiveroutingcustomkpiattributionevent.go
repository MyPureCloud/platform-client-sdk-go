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
	// EventId - A unique UUID v4 used to identify each attribution event uniquely
	EventId *string `json:"eventId,omitempty"`

	// EventDateTime - The date and time at which the attributed business result was achieved, in UTC, formatted as an ISO-8601 string. For example, 2023-01-07T08:30:00.123Z. Use the time the business result actually occurred or was confirmed, not a later ingestion, synchronisation, reporting, or batch-processing time. When conversationId is supplied, this timestamp must fall within the conversation attribution window, or the event is rejected. This means eventDateTime must be no earlier than the conversation start time, and no later than the conversation end time plus the grace period for the conversation's media type (1 hour for voice, message and email).
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`

	// ExternalContactId - The id of the External Contact of the customer who was involved in the KPI attribution. If this field is populated, we will retrieve the relevant conversation internally - thus this field is not compatible with conversationId. You must chose to use either externalContactId or conversationId but not both on the same event.
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// ConversationId - The id of the conversation. This field must be omitted if externalContactId is already populated on the event.
	ConversationId *string `json:"conversationId,omitempty"`

	// AgentId - The id of the agent who achieved the KPI attribution that is getting the attribution.
	AgentId *string `json:"agentId,omitempty"`

	// KpiId - The id of the relevant revenue KPI within the Genesys Platform - this id can be retrieved from the Revenue KPI Admin UI.
	KpiId *string `json:"kpiId,omitempty"`

	// AssociatedValue - For Sales Conversion, Retention and Churn KPI type, the numeric value expected is \"1\". For SalesValue KPI type, a numeric value representing the aggregated value of all KPI attributions achieved by an agent during a conversation, we do not allow multiple event uploads for the same conversationId and agentId pair.
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
