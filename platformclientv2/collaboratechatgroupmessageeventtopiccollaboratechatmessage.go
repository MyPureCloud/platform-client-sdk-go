package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Collaboratechatgroupmessageeventtopiccollaboratechatmessage
type Collaboratechatgroupmessageeventtopiccollaboratechatmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MessageId
	MessageId *string `json:"messageId,omitempty"`

	// Created
	Created *string `json:"created,omitempty"`

	// ActionType
	ActionType *string `json:"actionType,omitempty"`

	// Body
	Body *string `json:"body,omitempty"`

	// From
	From *Collaboratechatgroupmessageeventtopiccollaboratechatentity `json:"from,omitempty"`

	// To
	To *Collaboratechatgroupmessageeventtopiccollaboratechatentity `json:"to,omitempty"`

	// Mentions
	Mentions *[]Collaboratechatgroupmessageeventtopiccollaboratechatentity `json:"mentions,omitempty"`

	// NotifyAll
	NotifyAll *bool `json:"notifyAll,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Collaboratechatgroupmessageeventtopiccollaboratechatmessage) SetField(field string, fieldValue interface{}) {
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

func (o Collaboratechatgroupmessageeventtopiccollaboratechatmessage) MarshalJSON() ([]byte, error) {
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
	type Alias Collaboratechatgroupmessageeventtopiccollaboratechatmessage
	
	return json.Marshal(&struct { 
		MessageId *string `json:"messageId,omitempty"`
		
		Created *string `json:"created,omitempty"`
		
		ActionType *string `json:"actionType,omitempty"`
		
		Body *string `json:"body,omitempty"`
		
		From *Collaboratechatgroupmessageeventtopiccollaboratechatentity `json:"from,omitempty"`
		
		To *Collaboratechatgroupmessageeventtopiccollaboratechatentity `json:"to,omitempty"`
		
		Mentions *[]Collaboratechatgroupmessageeventtopiccollaboratechatentity `json:"mentions,omitempty"`
		
		NotifyAll *bool `json:"notifyAll,omitempty"`
		Alias
	}{ 
		MessageId: o.MessageId,
		
		Created: o.Created,
		
		ActionType: o.ActionType,
		
		Body: o.Body,
		
		From: o.From,
		
		To: o.To,
		
		Mentions: o.Mentions,
		
		NotifyAll: o.NotifyAll,
		Alias:    (Alias)(o),
	})
}

func (o *Collaboratechatgroupmessageeventtopiccollaboratechatmessage) UnmarshalJSON(b []byte) error {
	var CollaboratechatgroupmessageeventtopiccollaboratechatmessageMap map[string]interface{}
	err := json.Unmarshal(b, &CollaboratechatgroupmessageeventtopiccollaboratechatmessageMap)
	if err != nil {
		return err
	}
	
	if MessageId, ok := CollaboratechatgroupmessageeventtopiccollaboratechatmessageMap["messageId"].(string); ok {
		o.MessageId = &MessageId
	}
    
	if Created, ok := CollaboratechatgroupmessageeventtopiccollaboratechatmessageMap["created"].(string); ok {
		o.Created = &Created
	}
    
	if ActionType, ok := CollaboratechatgroupmessageeventtopiccollaboratechatmessageMap["actionType"].(string); ok {
		o.ActionType = &ActionType
	}
    
	if Body, ok := CollaboratechatgroupmessageeventtopiccollaboratechatmessageMap["body"].(string); ok {
		o.Body = &Body
	}
    
	if From, ok := CollaboratechatgroupmessageeventtopiccollaboratechatmessageMap["from"].(map[string]interface{}); ok {
		FromString, _ := json.Marshal(From)
		json.Unmarshal(FromString, &o.From)
	}
	
	if To, ok := CollaboratechatgroupmessageeventtopiccollaboratechatmessageMap["to"].(map[string]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if Mentions, ok := CollaboratechatgroupmessageeventtopiccollaboratechatmessageMap["mentions"].([]interface{}); ok {
		MentionsString, _ := json.Marshal(Mentions)
		json.Unmarshal(MentionsString, &o.Mentions)
	}
	
	if NotifyAll, ok := CollaboratechatgroupmessageeventtopiccollaboratechatmessageMap["notifyAll"].(bool); ok {
		o.NotifyAll = &NotifyAll
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Collaboratechatgroupmessageeventtopiccollaboratechatmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
