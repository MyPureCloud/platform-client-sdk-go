package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingusertransferevent
type Messagingusertransferevent struct { 
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Messagingusertransferevent) SetField(field string, fieldValue interface{}) {
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

func (o Messagingusertransferevent) MarshalJSON() ([]byte, error) {
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
	type Alias Messagingusertransferevent
	
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
		
		DestinationUserId: o.DestinationUserId,
		Alias:    (Alias)(o),
	})
}

func (o *Messagingusertransferevent) UnmarshalJSON(b []byte) error {
	var MessagingusertransfereventMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingusertransfereventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := MessagingusertransfereventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := MessagingusertransfereventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := MessagingusertransfereventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if TransferType, ok := MessagingusertransfereventMap["transferType"].(string); ok {
		o.TransferType = &TransferType
	}
    
	if CommandId, ok := MessagingusertransfereventMap["commandId"].(string); ok {
		o.CommandId = &CommandId
	}
    
	if InitiatingCommunicationId, ok := MessagingusertransfereventMap["initiatingCommunicationId"].(string); ok {
		o.InitiatingCommunicationId = &InitiatingCommunicationId
	}
    
	if TargetCommunicationId, ok := MessagingusertransfereventMap["targetCommunicationId"].(string); ok {
		o.TargetCommunicationId = &TargetCommunicationId
	}
    
	if ObjectCommunicationId, ok := MessagingusertransfereventMap["objectCommunicationId"].(string); ok {
		o.ObjectCommunicationId = &ObjectCommunicationId
	}
    
	if DestinationUserId, ok := MessagingusertransfereventMap["destinationUserId"].(string); ok {
		o.DestinationUserId = &DestinationUserId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingusertransferevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
