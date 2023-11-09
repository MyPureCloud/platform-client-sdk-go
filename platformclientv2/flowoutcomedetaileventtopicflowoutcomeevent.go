package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowoutcomedetaileventtopicflowoutcomeevent
type Flowoutcomedetaileventtopicflowoutcomeevent struct { 
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

	// AddressTo
	AddressTo *string `json:"addressTo,omitempty"`

	// AddressFrom
	AddressFrom *string `json:"addressFrom,omitempty"`

	// Subject
	Subject *string `json:"subject,omitempty"`

	// MessageType
	MessageType *string `json:"messageType,omitempty"`

	// FlowType
	FlowType *string `json:"flowType,omitempty"`

	// FlowId
	FlowId *string `json:"flowId,omitempty"`

	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`

	// FlowVersion
	FlowVersion *string `json:"flowVersion,omitempty"`

	// FlowOutcomeId
	FlowOutcomeId *string `json:"flowOutcomeId,omitempty"`

	// FlowOutcomeStartTime
	FlowOutcomeStartTime *int `json:"flowOutcomeStartTime,omitempty"`

	// FlowOutcomeEndTime
	FlowOutcomeEndTime *int `json:"flowOutcomeEndTime,omitempty"`

	// FlowOutcomeValue
	FlowOutcomeValue *string `json:"flowOutcomeValue,omitempty"`

	// FlowMilestones
	FlowMilestones *[]Flowoutcomedetaileventtopicflowmilestone `json:"flowMilestones,omitempty"`

	// ConversationExternalContactIds
	ConversationExternalContactIds *[]string `json:"conversationExternalContactIds,omitempty"`

	// ConversationExternalOrganizationIds
	ConversationExternalOrganizationIds *[]string `json:"conversationExternalOrganizationIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowoutcomedetaileventtopicflowoutcomeevent) SetField(field string, fieldValue interface{}) {
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

func (o Flowoutcomedetaileventtopicflowoutcomeevent) MarshalJSON() ([]byte, error) {
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
	type Alias Flowoutcomedetaileventtopicflowoutcomeevent
	
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
		
		AddressTo *string `json:"addressTo,omitempty"`
		
		AddressFrom *string `json:"addressFrom,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		FlowType *string `json:"flowType,omitempty"`
		
		FlowId *string `json:"flowId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		FlowVersion *string `json:"flowVersion,omitempty"`
		
		FlowOutcomeId *string `json:"flowOutcomeId,omitempty"`
		
		FlowOutcomeStartTime *int `json:"flowOutcomeStartTime,omitempty"`
		
		FlowOutcomeEndTime *int `json:"flowOutcomeEndTime,omitempty"`
		
		FlowOutcomeValue *string `json:"flowOutcomeValue,omitempty"`
		
		FlowMilestones *[]Flowoutcomedetaileventtopicflowmilestone `json:"flowMilestones,omitempty"`
		
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
		
		AddressTo: o.AddressTo,
		
		AddressFrom: o.AddressFrom,
		
		Subject: o.Subject,
		
		MessageType: o.MessageType,
		
		FlowType: o.FlowType,
		
		FlowId: o.FlowId,
		
		DivisionId: o.DivisionId,
		
		FlowVersion: o.FlowVersion,
		
		FlowOutcomeId: o.FlowOutcomeId,
		
		FlowOutcomeStartTime: o.FlowOutcomeStartTime,
		
		FlowOutcomeEndTime: o.FlowOutcomeEndTime,
		
		FlowOutcomeValue: o.FlowOutcomeValue,
		
		FlowMilestones: o.FlowMilestones,
		
		ConversationExternalContactIds: o.ConversationExternalContactIds,
		
		ConversationExternalOrganizationIds: o.ConversationExternalOrganizationIds,
		Alias:    (Alias)(o),
	})
}

func (o *Flowoutcomedetaileventtopicflowoutcomeevent) UnmarshalJSON(b []byte) error {
	var FlowoutcomedetaileventtopicflowoutcomeeventMap map[string]interface{}
	err := json.Unmarshal(b, &FlowoutcomedetaileventtopicflowoutcomeeventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if SessionId, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if MediaType, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Provider, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Direction, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Ani, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if AddressTo, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if AddressFrom, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if Subject, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if MessageType, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if FlowType, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["flowType"].(string); ok {
		o.FlowType = &FlowType
	}
    
	if FlowId, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["flowId"].(string); ok {
		o.FlowId = &FlowId
	}
    
	if DivisionId, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if FlowVersion, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["flowVersion"].(string); ok {
		o.FlowVersion = &FlowVersion
	}
    
	if FlowOutcomeId, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["flowOutcomeId"].(string); ok {
		o.FlowOutcomeId = &FlowOutcomeId
	}
    
	if FlowOutcomeStartTime, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["flowOutcomeStartTime"].(float64); ok {
		FlowOutcomeStartTimeInt := int(FlowOutcomeStartTime)
		o.FlowOutcomeStartTime = &FlowOutcomeStartTimeInt
	}
	
	if FlowOutcomeEndTime, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["flowOutcomeEndTime"].(float64); ok {
		FlowOutcomeEndTimeInt := int(FlowOutcomeEndTime)
		o.FlowOutcomeEndTime = &FlowOutcomeEndTimeInt
	}
	
	if FlowOutcomeValue, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["flowOutcomeValue"].(string); ok {
		o.FlowOutcomeValue = &FlowOutcomeValue
	}
    
	if FlowMilestones, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["flowMilestones"].([]interface{}); ok {
		FlowMilestonesString, _ := json.Marshal(FlowMilestones)
		json.Unmarshal(FlowMilestonesString, &o.FlowMilestones)
	}
	
	if ConversationExternalContactIds, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["conversationExternalContactIds"].([]interface{}); ok {
		ConversationExternalContactIdsString, _ := json.Marshal(ConversationExternalContactIds)
		json.Unmarshal(ConversationExternalContactIdsString, &o.ConversationExternalContactIds)
	}
	
	if ConversationExternalOrganizationIds, ok := FlowoutcomedetaileventtopicflowoutcomeeventMap["conversationExternalOrganizationIds"].([]interface{}); ok {
		ConversationExternalOrganizationIdsString, _ := json.Marshal(ConversationExternalOrganizationIds)
		json.Unmarshal(ConversationExternalOrganizationIdsString, &o.ConversationExternalOrganizationIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowoutcomedetaileventtopicflowoutcomeevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
