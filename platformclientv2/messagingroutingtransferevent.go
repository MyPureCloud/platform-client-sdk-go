package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingroutingtransferevent
type Messagingroutingtransferevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Messagingroutingtransferevent) SetField(field string, fieldValue interface{}) {
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

func (o Messagingroutingtransferevent) MarshalJSON() ([]byte, error) {
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
	type Alias Messagingroutingtransferevent
	
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
		Alias
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
		Alias:    (Alias)(o),
	})
}

func (o *Messagingroutingtransferevent) UnmarshalJSON(b []byte) error {
	var MessagingroutingtransfereventMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingroutingtransfereventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := MessagingroutingtransfereventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := MessagingroutingtransfereventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := MessagingroutingtransfereventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if TransferType, ok := MessagingroutingtransfereventMap["transferType"].(string); ok {
		o.TransferType = &TransferType
	}
    
	if CommandId, ok := MessagingroutingtransfereventMap["commandId"].(string); ok {
		o.CommandId = &CommandId
	}
    
	if InitiatingCommunicationId, ok := MessagingroutingtransfereventMap["initiatingCommunicationId"].(string); ok {
		o.InitiatingCommunicationId = &InitiatingCommunicationId
	}
    
	if TargetCommunicationId, ok := MessagingroutingtransfereventMap["targetCommunicationId"].(string); ok {
		o.TargetCommunicationId = &TargetCommunicationId
	}
    
	if ObjectCommunicationId, ok := MessagingroutingtransfereventMap["objectCommunicationId"].(string); ok {
		o.ObjectCommunicationId = &ObjectCommunicationId
	}
    
	if DestinationQueueId, ok := MessagingroutingtransfereventMap["destinationQueueId"].(string); ok {
		o.DestinationQueueId = &DestinationQueueId
	}
    
	if LanguageId, ok := MessagingroutingtransfereventMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if SkillIds, ok := MessagingroutingtransfereventMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingroutingtransferevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
