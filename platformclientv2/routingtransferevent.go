package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingtransferevent
type Routingtransferevent struct { 
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// TransferType - Indicates the desired type of transfer.
	TransferType *string `json:"transferType,omitempty"`


	// CommandId - The id (V4 UUID) used by the external platform to refer to the transfer in subsequent *Transfer events.
	CommandId *string `json:"commandId,omitempty"`


	// InitiatingCommunicationId - Indicates the desired type of transfer.
	InitiatingCommunicationId *string `json:"initiatingCommunicationId,omitempty"`


	// TargetCommunicationId - The id (V4 UUID) of the communication that is being transferred away from. In many cases this will be the same as the `initiatingCommunicationId`.
	TargetCommunicationId *string `json:"targetCommunicationId,omitempty"`


	// ObjectCommunicationId - The id (V4 UUID) of the communication that is being transferred.
	ObjectCommunicationId *string `json:"objectCommunicationId,omitempty"`


	// DestinationQueueId - The id (V4 UUID) of the desired destination queue that the object communication should be transferred to.
	DestinationQueueId *string `json:"destinationQueueId,omitempty"`


	// LanguageId - The unique identifier (V4 UUID) for the language that should be used to determine the destination for the conversation.
	LanguageId *string `json:"languageId,omitempty"`


	// SkillIds - The unique identifiers (V4 UUID) for the skills that should be used to determine the destination for the conversation.
	SkillIds *[]string `json:"skillIds,omitempty"`

}

func (o *Routingtransferevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingtransferevent
	
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
		
		DestinationQueueId *string `json:"destinationQueueId,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
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
		
		DestinationQueueId: o.DestinationQueueId,
		
		LanguageId: o.LanguageId,
		
		SkillIds: o.SkillIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Routingtransferevent) UnmarshalJSON(b []byte) error {
	var RoutingtransfereventMap map[string]interface{}
	err := json.Unmarshal(b, &RoutingtransfereventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := RoutingtransfereventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := RoutingtransfereventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := RoutingtransfereventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if TransferType, ok := RoutingtransfereventMap["transferType"].(string); ok {
		o.TransferType = &TransferType
	}
    
	if CommandId, ok := RoutingtransfereventMap["commandId"].(string); ok {
		o.CommandId = &CommandId
	}
    
	if InitiatingCommunicationId, ok := RoutingtransfereventMap["initiatingCommunicationId"].(string); ok {
		o.InitiatingCommunicationId = &InitiatingCommunicationId
	}
    
	if TargetCommunicationId, ok := RoutingtransfereventMap["targetCommunicationId"].(string); ok {
		o.TargetCommunicationId = &TargetCommunicationId
	}
    
	if ObjectCommunicationId, ok := RoutingtransfereventMap["objectCommunicationId"].(string); ok {
		o.ObjectCommunicationId = &ObjectCommunicationId
	}
    
	if DestinationQueueId, ok := RoutingtransfereventMap["destinationQueueId"].(string); ok {
		o.DestinationQueueId = &DestinationQueueId
	}
    
	if LanguageId, ok := RoutingtransfereventMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if SkillIds, ok := RoutingtransfereventMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routingtransferevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
