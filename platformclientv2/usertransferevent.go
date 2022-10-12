package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Usertransferevent
type Usertransferevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// TransferType - Indicates the desired type of transfer.
	TransferType *string `json:"transferType,omitempty"`


	// CommandId - The id (V4 UUID) used by the external platform to refer to the transfer in subsequent Transfer events.
	CommandId *string `json:"commandId,omitempty"`


	// InitiatingCommunicationId - The id (V4 UUID) of the communication representing the participant that is initiating the transfer.
	InitiatingCommunicationId *string `json:"initiatingCommunicationId,omitempty"`


	// TargetCommunicationId - The id (V4 UUID) of the communication that is being transferred away from. In many cases this will be the same as the `initiatingCommunicationId`.
	TargetCommunicationId *string `json:"targetCommunicationId,omitempty"`


	// ObjectCommunicationId - The id (V4 UUID) of the communication that is being transferred.
	ObjectCommunicationId *string `json:"objectCommunicationId,omitempty"`


	// DestinationUserId - The id (V4 UUID) of the desired destination user that the object communication should be transferred to.
	DestinationUserId *string `json:"destinationUserId,omitempty"`

}

func (o *Usertransferevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Usertransferevent
	
	EventDateTime := new(string)
	if o.EventDateTime != nil {
		
		*EventDateTime = timeutil.Strftime(o.EventDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventDateTime = nil
	}
	
	return json.Marshal(&struct { 
		EventId *string `json:"eventId,omitempty"`
		
		EventDateTime *string `json:"eventDateTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		TransferType *string `json:"transferType,omitempty"`
		
		CommandId *string `json:"commandId,omitempty"`
		
		InitiatingCommunicationId *string `json:"initiatingCommunicationId,omitempty"`
		
		TargetCommunicationId *string `json:"targetCommunicationId,omitempty"`
		
		ObjectCommunicationId *string `json:"objectCommunicationId,omitempty"`
		
		DestinationUserId *string `json:"destinationUserId,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		TransferType: o.TransferType,
		
		CommandId: o.CommandId,
		
		InitiatingCommunicationId: o.InitiatingCommunicationId,
		
		TargetCommunicationId: o.TargetCommunicationId,
		
		ObjectCommunicationId: o.ObjectCommunicationId,
		
		DestinationUserId: o.DestinationUserId,
		Alias:    (*Alias)(o),
	})
}

func (o *Usertransferevent) UnmarshalJSON(b []byte) error {
	var UsertransfereventMap map[string]interface{}
	err := json.Unmarshal(b, &UsertransfereventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := UsertransfereventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := UsertransfereventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := UsertransfereventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if TransferType, ok := UsertransfereventMap["transferType"].(string); ok {
		o.TransferType = &TransferType
	}
    
	if CommandId, ok := UsertransfereventMap["commandId"].(string); ok {
		o.CommandId = &CommandId
	}
    
	if InitiatingCommunicationId, ok := UsertransfereventMap["initiatingCommunicationId"].(string); ok {
		o.InitiatingCommunicationId = &InitiatingCommunicationId
	}
    
	if TargetCommunicationId, ok := UsertransfereventMap["targetCommunicationId"].(string); ok {
		o.TargetCommunicationId = &TargetCommunicationId
	}
    
	if ObjectCommunicationId, ok := UsertransfereventMap["objectCommunicationId"].(string); ok {
		o.ObjectCommunicationId = &ObjectCommunicationId
	}
    
	if DestinationUserId, ok := UsertransfereventMap["destinationUserId"].(string); ok {
		o.DestinationUserId = &DestinationUserId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Usertransferevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
