package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationwebeventsnotification
type Journeywebeventsnotificationwebeventsnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`

	// ExternalContact
	ExternalContact *Journeywebeventsnotificationexternalcontact `json:"externalContact,omitempty"`

	// CreatedDate
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// CustomerId
	CustomerId *string `json:"customerId,omitempty"`

	// CustomerIdType
	CustomerIdType *string `json:"customerIdType,omitempty"`

	// Session
	Session *Journeywebeventsnotificationsession `json:"session,omitempty"`

	// EventType
	EventType *string `json:"eventType,omitempty"`

	// WebEvent
	WebEvent *Journeywebeventsnotificationwebmessage `json:"webEvent,omitempty"`

	// WebActionEvent
	WebActionEvent *Journeywebeventsnotificationwebactionmessage `json:"webActionEvent,omitempty"`

	// OutcomeAchievedEvent
	OutcomeAchievedEvent *Journeywebeventsnotificationoutcomeachievedmessage `json:"outcomeAchievedEvent,omitempty"`

	// SegmentAssignmentEvent
	SegmentAssignmentEvent *Journeywebeventsnotificationsegmentassignmentmessage `json:"segmentAssignmentEvent,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeywebeventsnotificationwebeventsnotification) SetField(field string, fieldValue interface{}) {
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

func (o Journeywebeventsnotificationwebeventsnotification) MarshalJSON() ([]byte, error) {
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
	type Alias Journeywebeventsnotificationwebeventsnotification
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		ExternalContact *Journeywebeventsnotificationexternalcontact `json:"externalContact,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		CustomerId *string `json:"customerId,omitempty"`
		
		CustomerIdType *string `json:"customerIdType,omitempty"`
		
		Session *Journeywebeventsnotificationsession `json:"session,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		
		WebEvent *Journeywebeventsnotificationwebmessage `json:"webEvent,omitempty"`
		
		WebActionEvent *Journeywebeventsnotificationwebactionmessage `json:"webActionEvent,omitempty"`
		
		OutcomeAchievedEvent *Journeywebeventsnotificationoutcomeachievedmessage `json:"outcomeAchievedEvent,omitempty"`
		
		SegmentAssignmentEvent *Journeywebeventsnotificationsegmentassignmentmessage `json:"segmentAssignmentEvent,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		CorrelationId: o.CorrelationId,
		
		ExternalContact: o.ExternalContact,
		
		CreatedDate: CreatedDate,
		
		CustomerId: o.CustomerId,
		
		CustomerIdType: o.CustomerIdType,
		
		Session: o.Session,
		
		EventType: o.EventType,
		
		WebEvent: o.WebEvent,
		
		WebActionEvent: o.WebActionEvent,
		
		OutcomeAchievedEvent: o.OutcomeAchievedEvent,
		
		SegmentAssignmentEvent: o.SegmentAssignmentEvent,
		Alias:    (Alias)(o),
	})
}

func (o *Journeywebeventsnotificationwebeventsnotification) UnmarshalJSON(b []byte) error {
	var JourneywebeventsnotificationwebeventsnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebeventsnotificationwebeventsnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneywebeventsnotificationwebeventsnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if CorrelationId, ok := JourneywebeventsnotificationwebeventsnotificationMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if ExternalContact, ok := JourneywebeventsnotificationwebeventsnotificationMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if createdDateString, ok := JourneywebeventsnotificationwebeventsnotificationMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if CustomerId, ok := JourneywebeventsnotificationwebeventsnotificationMap["customerId"].(string); ok {
		o.CustomerId = &CustomerId
	}
    
	if CustomerIdType, ok := JourneywebeventsnotificationwebeventsnotificationMap["customerIdType"].(string); ok {
		o.CustomerIdType = &CustomerIdType
	}
    
	if Session, ok := JourneywebeventsnotificationwebeventsnotificationMap["session"].(map[string]interface{}); ok {
		SessionString, _ := json.Marshal(Session)
		json.Unmarshal(SessionString, &o.Session)
	}
	
	if EventType, ok := JourneywebeventsnotificationwebeventsnotificationMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if WebEvent, ok := JourneywebeventsnotificationwebeventsnotificationMap["webEvent"].(map[string]interface{}); ok {
		WebEventString, _ := json.Marshal(WebEvent)
		json.Unmarshal(WebEventString, &o.WebEvent)
	}
	
	if WebActionEvent, ok := JourneywebeventsnotificationwebeventsnotificationMap["webActionEvent"].(map[string]interface{}); ok {
		WebActionEventString, _ := json.Marshal(WebActionEvent)
		json.Unmarshal(WebActionEventString, &o.WebActionEvent)
	}
	
	if OutcomeAchievedEvent, ok := JourneywebeventsnotificationwebeventsnotificationMap["outcomeAchievedEvent"].(map[string]interface{}); ok {
		OutcomeAchievedEventString, _ := json.Marshal(OutcomeAchievedEvent)
		json.Unmarshal(OutcomeAchievedEventString, &o.OutcomeAchievedEvent)
	}
	
	if SegmentAssignmentEvent, ok := JourneywebeventsnotificationwebeventsnotificationMap["segmentAssignmentEvent"].(map[string]interface{}); ok {
		SegmentAssignmentEventString, _ := json.Marshal(SegmentAssignmentEvent)
		json.Unmarshal(SegmentAssignmentEventString, &o.SegmentAssignmentEvent)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationwebeventsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
