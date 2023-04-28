package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactdetaileventtopiccontactupdateevent
type Contactdetaileventtopiccontactupdateevent struct { 
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

	// ExternalOrganizationId
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`

	// ExternalContactId
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// Provider
	Provider *string `json:"provider,omitempty"`

	// Direction
	Direction *string `json:"direction,omitempty"`

	// Ani
	Ani *string `json:"ani,omitempty"`

	// Dnis
	Dnis *string `json:"dnis,omitempty"`

	// AddressTo
	AddressTo *string `json:"addressTo,omitempty"`

	// AddressFrom
	AddressFrom *string `json:"addressFrom,omitempty"`

	// CallbackUserName
	CallbackUserName *string `json:"callbackUserName,omitempty"`

	// CallbackNumbers
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`

	// CallbackScheduledTime
	CallbackScheduledTime *int `json:"callbackScheduledTime,omitempty"`

	// Subject
	Subject *string `json:"subject,omitempty"`

	// MessageType
	MessageType *string `json:"messageType,omitempty"`

	// ConversationExternalContactIds
	ConversationExternalContactIds *[]string `json:"conversationExternalContactIds,omitempty"`

	// ConversationExternalOrganizationIds
	ConversationExternalOrganizationIds *[]string `json:"conversationExternalOrganizationIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactdetaileventtopiccontactupdateevent) SetField(field string, fieldValue interface{}) {
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

func (o Contactdetaileventtopiccontactupdateevent) MarshalJSON() ([]byte, error) {
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
	type Alias Contactdetaileventtopiccontactupdateevent
	
	return json.Marshal(&struct { 
		EventTime *int `json:"eventTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		AddressTo *string `json:"addressTo,omitempty"`
		
		AddressFrom *string `json:"addressFrom,omitempty"`
		
		CallbackUserName *string `json:"callbackUserName,omitempty"`
		
		CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`
		
		CallbackScheduledTime *int `json:"callbackScheduledTime,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		ConversationExternalContactIds *[]string `json:"conversationExternalContactIds,omitempty"`
		
		ConversationExternalOrganizationIds *[]string `json:"conversationExternalOrganizationIds,omitempty"`
		Alias
	}{ 
		EventTime: o.EventTime,
		
		ConversationId: o.ConversationId,
		
		ParticipantId: o.ParticipantId,
		
		SessionId: o.SessionId,
		
		MediaType: o.MediaType,
		
		ExternalOrganizationId: o.ExternalOrganizationId,
		
		ExternalContactId: o.ExternalContactId,
		
		Provider: o.Provider,
		
		Direction: o.Direction,
		
		Ani: o.Ani,
		
		Dnis: o.Dnis,
		
		AddressTo: o.AddressTo,
		
		AddressFrom: o.AddressFrom,
		
		CallbackUserName: o.CallbackUserName,
		
		CallbackNumbers: o.CallbackNumbers,
		
		CallbackScheduledTime: o.CallbackScheduledTime,
		
		Subject: o.Subject,
		
		MessageType: o.MessageType,
		
		ConversationExternalContactIds: o.ConversationExternalContactIds,
		
		ConversationExternalOrganizationIds: o.ConversationExternalOrganizationIds,
		Alias:    (Alias)(o),
	})
}

func (o *Contactdetaileventtopiccontactupdateevent) UnmarshalJSON(b []byte) error {
	var ContactdetaileventtopiccontactupdateeventMap map[string]interface{}
	err := json.Unmarshal(b, &ContactdetaileventtopiccontactupdateeventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := ContactdetaileventtopiccontactupdateeventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := ContactdetaileventtopiccontactupdateeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := ContactdetaileventtopiccontactupdateeventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if SessionId, ok := ContactdetaileventtopiccontactupdateeventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if MediaType, ok := ContactdetaileventtopiccontactupdateeventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ExternalOrganizationId, ok := ContactdetaileventtopiccontactupdateeventMap["externalOrganizationId"].(string); ok {
		o.ExternalOrganizationId = &ExternalOrganizationId
	}
    
	if ExternalContactId, ok := ContactdetaileventtopiccontactupdateeventMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if Provider, ok := ContactdetaileventtopiccontactupdateeventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Direction, ok := ContactdetaileventtopiccontactupdateeventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Ani, ok := ContactdetaileventtopiccontactupdateeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := ContactdetaileventtopiccontactupdateeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if AddressTo, ok := ContactdetaileventtopiccontactupdateeventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if AddressFrom, ok := ContactdetaileventtopiccontactupdateeventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if CallbackUserName, ok := ContactdetaileventtopiccontactupdateeventMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
    
	if CallbackNumbers, ok := ContactdetaileventtopiccontactupdateeventMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if CallbackScheduledTime, ok := ContactdetaileventtopiccontactupdateeventMap["callbackScheduledTime"].(float64); ok {
		CallbackScheduledTimeInt := int(CallbackScheduledTime)
		o.CallbackScheduledTime = &CallbackScheduledTimeInt
	}
	
	if Subject, ok := ContactdetaileventtopiccontactupdateeventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if MessageType, ok := ContactdetaileventtopiccontactupdateeventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if ConversationExternalContactIds, ok := ContactdetaileventtopiccontactupdateeventMap["conversationExternalContactIds"].([]interface{}); ok {
		ConversationExternalContactIdsString, _ := json.Marshal(ConversationExternalContactIds)
		json.Unmarshal(ConversationExternalContactIdsString, &o.ConversationExternalContactIds)
	}
	
	if ConversationExternalOrganizationIds, ok := ContactdetaileventtopiccontactupdateeventMap["conversationExternalOrganizationIds"].([]interface{}); ok {
		ConversationExternalOrganizationIdsString, _ := json.Marshal(ConversationExternalOrganizationIds)
		json.Unmarshal(ConversationExternalOrganizationIdsString, &o.ConversationExternalOrganizationIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactdetaileventtopiccontactupdateevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
