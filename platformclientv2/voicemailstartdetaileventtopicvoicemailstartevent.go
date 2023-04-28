package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailstartdetaileventtopicvoicemailstartevent
type Voicemailstartdetaileventtopicvoicemailstartevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventTime
	EventTime *int `json:"eventTime,omitempty"`

	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// ParticipantId
	ParticipantId *string `json:"participantId,omitempty"`

	// SessionId
	SessionId *string `json:"sessionId,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// Provider
	Provider *string `json:"provider,omitempty"`

	// Direction
	Direction *string `json:"direction,omitempty"`

	// Ani
	Ani *string `json:"ani,omitempty"`

	// Dnis
	Dnis *string `json:"dnis,omitempty"`

	// UserId
	UserId *string `json:"userId,omitempty"`

	// QueueId
	QueueId *string `json:"queueId,omitempty"`

	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`

	// ConversationExternalContactIds
	ConversationExternalContactIds *[]string `json:"conversationExternalContactIds,omitempty"`

	// ConversationExternalOrganizationIds
	ConversationExternalOrganizationIds *[]string `json:"conversationExternalOrganizationIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Voicemailstartdetaileventtopicvoicemailstartevent) SetField(field string, fieldValue interface{}) {
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

func (o Voicemailstartdetaileventtopicvoicemailstartevent) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Voicemailstartdetaileventtopicvoicemailstartevent
	
	return json.Marshal(&struct { 
		EventTime *int `json:"eventTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		ConversationExternalContactIds *[]string `json:"conversationExternalContactIds,omitempty"`
		
		ConversationExternalOrganizationIds *[]string `json:"conversationExternalOrganizationIds,omitempty"`
		Alias
	}{ 
		EventTime: o.EventTime,
		
		ConversationId: o.ConversationId,
		
		ParticipantId: o.ParticipantId,
		
		SessionId: o.SessionId,
		
		MediaType: o.MediaType,
		
		Provider: o.Provider,
		
		Direction: o.Direction,
		
		Ani: o.Ani,
		
		Dnis: o.Dnis,
		
		UserId: o.UserId,
		
		QueueId: o.QueueId,
		
		DivisionId: o.DivisionId,
		
		ConversationExternalContactIds: o.ConversationExternalContactIds,
		
		ConversationExternalOrganizationIds: o.ConversationExternalOrganizationIds,
		Alias:    (Alias)(o),
	})
}

func (o *Voicemailstartdetaileventtopicvoicemailstartevent) UnmarshalJSON(b []byte) error {
	var VoicemailstartdetaileventtopicvoicemailstarteventMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailstartdetaileventtopicvoicemailstarteventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if SessionId, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if MediaType, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Provider, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Direction, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Ani, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if UserId, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if QueueId, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if DivisionId, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if ConversationExternalContactIds, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["conversationExternalContactIds"].([]interface{}); ok {
		ConversationExternalContactIdsString, _ := json.Marshal(ConversationExternalContactIds)
		json.Unmarshal(ConversationExternalContactIdsString, &o.ConversationExternalContactIds)
	}
	
	if ConversationExternalOrganizationIds, ok := VoicemailstartdetaileventtopicvoicemailstarteventMap["conversationExternalOrganizationIds"].([]interface{}); ok {
		ConversationExternalOrganizationIdsString, _ := json.Marshal(ConversationExternalOrganizationIds)
		json.Unmarshal(ConversationExternalOrganizationIdsString, &o.ConversationExternalOrganizationIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemailstartdetaileventtopicvoicemailstartevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
