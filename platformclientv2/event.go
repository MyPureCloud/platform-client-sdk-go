package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Event
type Event struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - System-generated UUID for the event.
	Id *string `json:"id,omitempty"`

	// CorrelationId - UUID corresponding to triggering action that caused this event (e.g. HTTP POST, SIP invite, another event).
	CorrelationId *string `json:"correlationId,omitempty"`

	// CustomerId - Primary identifier of the customer in the source of the events.
	CustomerId *string `json:"customerId,omitempty"`

	// CustomerIdType - Type of primary identifier (e.g. cookie, email, phone).
	CustomerIdType *string `json:"customerIdType,omitempty"`

	// Session - The session that the event belongs to.
	Session *Eventsession `json:"session,omitempty"`

	// EventType - The name representing the type of event.
	EventType *string `json:"eventType,omitempty"`

	// OutcomeAchievedEvent - Event where a customer has achieved a specific outcome or goal.
	OutcomeAchievedEvent *Outcomeachievedevent `json:"outcomeAchievedEvent,omitempty"`

	// SegmentAssignmentEvent - Event that represents a segment being assigned.
	SegmentAssignmentEvent *Segmentassignmentevent `json:"segmentAssignmentEvent,omitempty"`

	// WebActionEvent - Event triggered by web actions.
	WebActionEvent *Webactionevent `json:"webActionEvent,omitempty"`

	// WebEvent - Event that tracks user interactions with content in a browser such as pageviews, downloads, mobile ad clicks, etc.
	WebEvent *Webevent `json:"webEvent,omitempty"`

	// AppEvent - Event that tracks user interactions with content in an application such as screen views, searches, etc.
	AppEvent *Appevent `json:"appEvent,omitempty"`

	// CreatedDate - Timestamp indicating when the event actually took place. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Event) SetField(field string, fieldValue interface{}) {
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

func (o Event) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate", }
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
	type Alias Event
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		CustomerId *string `json:"customerId,omitempty"`
		
		CustomerIdType *string `json:"customerIdType,omitempty"`
		
		Session *Eventsession `json:"session,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		
		OutcomeAchievedEvent *Outcomeachievedevent `json:"outcomeAchievedEvent,omitempty"`
		
		SegmentAssignmentEvent *Segmentassignmentevent `json:"segmentAssignmentEvent,omitempty"`
		
		WebActionEvent *Webactionevent `json:"webActionEvent,omitempty"`
		
		WebEvent *Webevent `json:"webEvent,omitempty"`
		
		AppEvent *Appevent `json:"appEvent,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		CorrelationId: o.CorrelationId,
		
		CustomerId: o.CustomerId,
		
		CustomerIdType: o.CustomerIdType,
		
		Session: o.Session,
		
		EventType: o.EventType,
		
		OutcomeAchievedEvent: o.OutcomeAchievedEvent,
		
		SegmentAssignmentEvent: o.SegmentAssignmentEvent,
		
		WebActionEvent: o.WebActionEvent,
		
		WebEvent: o.WebEvent,
		
		AppEvent: o.AppEvent,
		
		CreatedDate: CreatedDate,
		Alias:    (Alias)(o),
	})
}

func (o *Event) UnmarshalJSON(b []byte) error {
	var EventMap map[string]interface{}
	err := json.Unmarshal(b, &EventMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EventMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if CorrelationId, ok := EventMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if CustomerId, ok := EventMap["customerId"].(string); ok {
		o.CustomerId = &CustomerId
	}
    
	if CustomerIdType, ok := EventMap["customerIdType"].(string); ok {
		o.CustomerIdType = &CustomerIdType
	}
    
	if Session, ok := EventMap["session"].(map[string]interface{}); ok {
		SessionString, _ := json.Marshal(Session)
		json.Unmarshal(SessionString, &o.Session)
	}
	
	if EventType, ok := EventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if OutcomeAchievedEvent, ok := EventMap["outcomeAchievedEvent"].(map[string]interface{}); ok {
		OutcomeAchievedEventString, _ := json.Marshal(OutcomeAchievedEvent)
		json.Unmarshal(OutcomeAchievedEventString, &o.OutcomeAchievedEvent)
	}
	
	if SegmentAssignmentEvent, ok := EventMap["segmentAssignmentEvent"].(map[string]interface{}); ok {
		SegmentAssignmentEventString, _ := json.Marshal(SegmentAssignmentEvent)
		json.Unmarshal(SegmentAssignmentEventString, &o.SegmentAssignmentEvent)
	}
	
	if WebActionEvent, ok := EventMap["webActionEvent"].(map[string]interface{}); ok {
		WebActionEventString, _ := json.Marshal(WebActionEvent)
		json.Unmarshal(WebActionEventString, &o.WebActionEvent)
	}
	
	if WebEvent, ok := EventMap["webEvent"].(map[string]interface{}); ok {
		WebEventString, _ := json.Marshal(WebEvent)
		json.Unmarshal(WebEventString, &o.WebEvent)
	}
	
	if AppEvent, ok := EventMap["appEvent"].(map[string]interface{}); ok {
		AppEventString, _ := json.Marshal(AppEvent)
		json.Unmarshal(AppEventString, &o.AppEvent)
	}
	
	if createdDateString, ok := EventMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Event) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
