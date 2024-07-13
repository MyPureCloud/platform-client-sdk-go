package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createcallrequest
type Createcallrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PhoneNumber - The phone number to dial.
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// CallerId - The caller id phone number for this outbound call.
	CallerId *string `json:"callerId,omitempty"`

	// CallerIdName - The caller id name for this outbound call.
	CallerIdName *string `json:"callerIdName,omitempty"`

	// CallFromQueueId - The queue ID to call on behalf of.
	CallFromQueueId *string `json:"callFromQueueId,omitempty"`

	// CallQueueId - The queue ID to call.
	CallQueueId *string `json:"callQueueId,omitempty"`

	// CallUserId - The user ID to call.
	CallUserId *string `json:"callUserId,omitempty"`

	// Priority - The priority to assign to this call (if calling a queue).
	Priority *int `json:"priority,omitempty"`

	// Attributes - The list of attributes to associate with the customer participant.
	Attributes *map[string]string `json:"attributes,omitempty"`

	// LanguageId - The language skill ID to use for routing this call (if calling a queue).
	LanguageId *string `json:"languageId,omitempty"`

	// RoutingSkillsIds - The skill ID's to use for routing this call (if calling a queue).
	RoutingSkillsIds *[]string `json:"routingSkillsIds,omitempty"`

	// ConversationIds - The list of existing call conversations to merge into a new ad-hoc conference.
	ConversationIds *[]string `json:"conversationIds,omitempty"`

	// Participants - The list of participants to call to create a new ad-hoc conference.
	Participants *[]Destination `json:"participants,omitempty"`

	// UuiData - User to User Information (UUI) data managed by SIP session application.
	UuiData *string `json:"uuiData,omitempty"`

	// ExternalContactId - The external contact with which to associate the call.
	ExternalContactId *string `json:"externalContactId,omitempty"`

	// Label - An optional label that categorizes the conversation.  Max-utilization settings can be configured at a per-label level
	Label *string `json:"label,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createcallrequest) SetField(field string, fieldValue interface{}) {
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

func (o Createcallrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createcallrequest
	
	return json.Marshal(&struct { 
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		CallerId *string `json:"callerId,omitempty"`
		
		CallerIdName *string `json:"callerIdName,omitempty"`
		
		CallFromQueueId *string `json:"callFromQueueId,omitempty"`
		
		CallQueueId *string `json:"callQueueId,omitempty"`
		
		CallUserId *string `json:"callUserId,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		RoutingSkillsIds *[]string `json:"routingSkillsIds,omitempty"`
		
		ConversationIds *[]string `json:"conversationIds,omitempty"`
		
		Participants *[]Destination `json:"participants,omitempty"`
		
		UuiData *string `json:"uuiData,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		Label *string `json:"label,omitempty"`
		Alias
	}{ 
		PhoneNumber: o.PhoneNumber,
		
		CallerId: o.CallerId,
		
		CallerIdName: o.CallerIdName,
		
		CallFromQueueId: o.CallFromQueueId,
		
		CallQueueId: o.CallQueueId,
		
		CallUserId: o.CallUserId,
		
		Priority: o.Priority,
		
		Attributes: o.Attributes,
		
		LanguageId: o.LanguageId,
		
		RoutingSkillsIds: o.RoutingSkillsIds,
		
		ConversationIds: o.ConversationIds,
		
		Participants: o.Participants,
		
		UuiData: o.UuiData,
		
		ExternalContactId: o.ExternalContactId,
		
		Label: o.Label,
		Alias:    (Alias)(o),
	})
}

func (o *Createcallrequest) UnmarshalJSON(b []byte) error {
	var CreatecallrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatecallrequestMap)
	if err != nil {
		return err
	}
	
	if PhoneNumber, ok := CreatecallrequestMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if CallerId, ok := CreatecallrequestMap["callerId"].(string); ok {
		o.CallerId = &CallerId
	}
    
	if CallerIdName, ok := CreatecallrequestMap["callerIdName"].(string); ok {
		o.CallerIdName = &CallerIdName
	}
    
	if CallFromQueueId, ok := CreatecallrequestMap["callFromQueueId"].(string); ok {
		o.CallFromQueueId = &CallFromQueueId
	}
    
	if CallQueueId, ok := CreatecallrequestMap["callQueueId"].(string); ok {
		o.CallQueueId = &CallQueueId
	}
    
	if CallUserId, ok := CreatecallrequestMap["callUserId"].(string); ok {
		o.CallUserId = &CallUserId
	}
    
	if Priority, ok := CreatecallrequestMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Attributes, ok := CreatecallrequestMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if LanguageId, ok := CreatecallrequestMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if RoutingSkillsIds, ok := CreatecallrequestMap["routingSkillsIds"].([]interface{}); ok {
		RoutingSkillsIdsString, _ := json.Marshal(RoutingSkillsIds)
		json.Unmarshal(RoutingSkillsIdsString, &o.RoutingSkillsIds)
	}
	
	if ConversationIds, ok := CreatecallrequestMap["conversationIds"].([]interface{}); ok {
		ConversationIdsString, _ := json.Marshal(ConversationIds)
		json.Unmarshal(ConversationIdsString, &o.ConversationIds)
	}
	
	if Participants, ok := CreatecallrequestMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if UuiData, ok := CreatecallrequestMap["uuiData"].(string); ok {
		o.UuiData = &UuiData
	}
    
	if ExternalContactId, ok := CreatecallrequestMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if Label, ok := CreatecallrequestMap["label"].(string); ok {
		o.Label = &Label
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createcallrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
