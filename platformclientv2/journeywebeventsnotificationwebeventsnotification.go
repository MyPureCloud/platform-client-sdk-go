package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationwebeventsnotification
type Journeywebeventsnotificationwebeventsnotification struct { 
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

}

func (o *Journeywebeventsnotificationwebeventsnotification) MarshalJSON() ([]byte, error) {
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
		*Alias
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
		Alias:    (*Alias)(o),
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationwebeventsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
