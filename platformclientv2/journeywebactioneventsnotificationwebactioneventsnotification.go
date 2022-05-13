package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebactioneventsnotificationwebactioneventsnotification
type Journeywebactioneventsnotificationwebactioneventsnotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`


	// ExternalContact
	ExternalContact *Journeywebactioneventsnotificationexternalcontact `json:"externalContact,omitempty"`


	// CreatedDate
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// CustomerId
	CustomerId *string `json:"customerId,omitempty"`


	// CustomerIdType
	CustomerIdType *string `json:"customerIdType,omitempty"`


	// Session
	Session *Journeywebactioneventsnotificationsession `json:"session,omitempty"`


	// EventType
	EventType *string `json:"eventType,omitempty"`


	// WebActionEvent
	WebActionEvent *Journeywebactioneventsnotificationwebactionmessage `json:"webActionEvent,omitempty"`


	// BlockedWebActionOfferEvent
	BlockedWebActionOfferEvent *Journeywebactioneventsnotificationblockedwebactionoffermessage `json:"blockedWebActionOfferEvent,omitempty"`

}

func (o *Journeywebactioneventsnotificationwebactioneventsnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebactioneventsnotificationwebactioneventsnotification
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		ExternalContact *Journeywebactioneventsnotificationexternalcontact `json:"externalContact,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		CustomerId *string `json:"customerId,omitempty"`
		
		CustomerIdType *string `json:"customerIdType,omitempty"`
		
		Session *Journeywebactioneventsnotificationsession `json:"session,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		
		WebActionEvent *Journeywebactioneventsnotificationwebactionmessage `json:"webActionEvent,omitempty"`
		
		BlockedWebActionOfferEvent *Journeywebactioneventsnotificationblockedwebactionoffermessage `json:"blockedWebActionOfferEvent,omitempty"`
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
		
		WebActionEvent: o.WebActionEvent,
		
		BlockedWebActionOfferEvent: o.BlockedWebActionOfferEvent,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebactioneventsnotificationwebactioneventsnotification) UnmarshalJSON(b []byte) error {
	var JourneywebactioneventsnotificationwebactioneventsnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebactioneventsnotificationwebactioneventsnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneywebactioneventsnotificationwebactioneventsnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if CorrelationId, ok := JourneywebactioneventsnotificationwebactioneventsnotificationMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if ExternalContact, ok := JourneywebactioneventsnotificationwebactioneventsnotificationMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if createdDateString, ok := JourneywebactioneventsnotificationwebactioneventsnotificationMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if CustomerId, ok := JourneywebactioneventsnotificationwebactioneventsnotificationMap["customerId"].(string); ok {
		o.CustomerId = &CustomerId
	}
    
	if CustomerIdType, ok := JourneywebactioneventsnotificationwebactioneventsnotificationMap["customerIdType"].(string); ok {
		o.CustomerIdType = &CustomerIdType
	}
    
	if Session, ok := JourneywebactioneventsnotificationwebactioneventsnotificationMap["session"].(map[string]interface{}); ok {
		SessionString, _ := json.Marshal(Session)
		json.Unmarshal(SessionString, &o.Session)
	}
	
	if EventType, ok := JourneywebactioneventsnotificationwebactioneventsnotificationMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if WebActionEvent, ok := JourneywebactioneventsnotificationwebactioneventsnotificationMap["webActionEvent"].(map[string]interface{}); ok {
		WebActionEventString, _ := json.Marshal(WebActionEvent)
		json.Unmarshal(WebActionEventString, &o.WebActionEvent)
	}
	
	if BlockedWebActionOfferEvent, ok := JourneywebactioneventsnotificationwebactioneventsnotificationMap["blockedWebActionOfferEvent"].(map[string]interface{}); ok {
		BlockedWebActionOfferEventString, _ := json.Marshal(BlockedWebActionOfferEvent)
		json.Unmarshal(BlockedWebActionOfferEventString, &o.BlockedWebActionOfferEvent)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebactioneventsnotificationwebactioneventsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
