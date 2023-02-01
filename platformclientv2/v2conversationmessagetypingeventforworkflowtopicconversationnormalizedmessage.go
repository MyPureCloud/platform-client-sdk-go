package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage
type V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Channel
	Channel *V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannel `json:"channel,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// Text
	Text *string `json:"text,omitempty"`

	// Content
	Content *[]V2conversationmessagetypingeventforworkflowtopicconversationmessagecontent `json:"content,omitempty"`

	// Events
	Events *[]V2conversationmessagetypingeventforworkflowtopicconversationmessageevent `json:"events,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// Reasons
	Reasons *[]V2conversationmessagetypingeventforworkflowtopicconversationreason `json:"reasons,omitempty"`

	// OriginatingEntity
	OriginatingEntity *string `json:"originatingEntity,omitempty"`

	// IsFinalReceipt
	IsFinalReceipt *bool `json:"isFinalReceipt,omitempty"`

	// Direction
	Direction *string `json:"direction,omitempty"`

	// Metadata
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage) SetField(field string, fieldValue interface{}) {
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

func (o V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage) MarshalJSON() ([]byte, error) {
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
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Channel *V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannel `json:"channel,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Content *[]V2conversationmessagetypingeventforworkflowtopicconversationmessagecontent `json:"content,omitempty"`
		
		Events *[]V2conversationmessagetypingeventforworkflowtopicconversationmessageevent `json:"events,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Reasons *[]V2conversationmessagetypingeventforworkflowtopicconversationreason `json:"reasons,omitempty"`
		
		OriginatingEntity *string `json:"originatingEntity,omitempty"`
		
		IsFinalReceipt *bool `json:"isFinalReceipt,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Metadata *map[string]string `json:"metadata,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Channel: o.Channel,
		
		VarType: o.VarType,
		
		Text: o.Text,
		
		Content: o.Content,
		
		Events: o.Events,
		
		Status: o.Status,
		
		Reasons: o.Reasons,
		
		OriginatingEntity: o.OriginatingEntity,
		
		IsFinalReceipt: o.IsFinalReceipt,
		
		Direction: o.Direction,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Channel, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["channel"].(map[string]interface{}); ok {
		ChannelString, _ := json.Marshal(Channel)
		json.Unmarshal(ChannelString, &o.Channel)
	}
	
	if VarType, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Content, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	
	if Events, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["events"].([]interface{}); ok {
		EventsString, _ := json.Marshal(Events)
		json.Unmarshal(EventsString, &o.Events)
	}
	
	if Status, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Reasons, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["reasons"].([]interface{}); ok {
		ReasonsString, _ := json.Marshal(Reasons)
		json.Unmarshal(ReasonsString, &o.Reasons)
	}
	
	if OriginatingEntity, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["originatingEntity"].(string); ok {
		o.OriginatingEntity = &OriginatingEntity
	}
    
	if IsFinalReceipt, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["isFinalReceipt"].(bool); ok {
		o.IsFinalReceipt = &IsFinalReceipt
	}
    
	if Direction, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Metadata, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
