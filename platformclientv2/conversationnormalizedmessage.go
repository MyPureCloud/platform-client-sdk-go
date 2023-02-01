package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationnormalizedmessage - General rich media message structure with normalized feature support across many messaging channels.
type Conversationnormalizedmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Unique ID of the message. Message receipts will have the same ID as the message they reference.
	Id *string `json:"id,omitempty"`

	// Channel - Channel-specific information that describes the message and the message channel/provider.
	Channel *Conversationmessagingchannel `json:"channel,omitempty"`

	// VarType - Message type.
	VarType *string `json:"type,omitempty"`

	// Text - Message text.
	Text *string `json:"text,omitempty"`

	// Content - List of content elements.
	Content *[]Conversationmessagecontent `json:"content,omitempty"`

	// Events - List of event elements.
	Events *[]Conversationmessageevent `json:"events,omitempty"`

	// Status - Message receipt status, only used with type Receipt.
	Status *string `json:"status,omitempty"`

	// Reasons - List of reasons for a message receipt that indicates the message has failed. Only used with Failed status.
	Reasons *[]Conversationreason `json:"reasons,omitempty"`

	// OriginatingEntity - Specifies if this message was sent by a human agent or bot. The platform may use this to apply appropriate provider policies.
	OriginatingEntity *string `json:"originatingEntity,omitempty"`

	// IsFinalReceipt - Indicates if this is the last message receipt for this message, or if another message receipt can be expected.
	IsFinalReceipt *bool `json:"isFinalReceipt,omitempty"`

	// Direction - The direction of the message.
	Direction *string `json:"direction,omitempty"`

	// Metadata - Additional metadata about this message.
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationnormalizedmessage) SetField(field string, fieldValue interface{}) {
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

func (o Conversationnormalizedmessage) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationnormalizedmessage
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Channel *Conversationmessagingchannel `json:"channel,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Content *[]Conversationmessagecontent `json:"content,omitempty"`
		
		Events *[]Conversationmessageevent `json:"events,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Reasons *[]Conversationreason `json:"reasons,omitempty"`
		
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

func (o *Conversationnormalizedmessage) UnmarshalJSON(b []byte) error {
	var ConversationnormalizedmessageMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationnormalizedmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationnormalizedmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Channel, ok := ConversationnormalizedmessageMap["channel"].(map[string]interface{}); ok {
		ChannelString, _ := json.Marshal(Channel)
		json.Unmarshal(ChannelString, &o.Channel)
	}
	
	if VarType, ok := ConversationnormalizedmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := ConversationnormalizedmessageMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Content, ok := ConversationnormalizedmessageMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	
	if Events, ok := ConversationnormalizedmessageMap["events"].([]interface{}); ok {
		EventsString, _ := json.Marshal(Events)
		json.Unmarshal(EventsString, &o.Events)
	}
	
	if Status, ok := ConversationnormalizedmessageMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Reasons, ok := ConversationnormalizedmessageMap["reasons"].([]interface{}); ok {
		ReasonsString, _ := json.Marshal(Reasons)
		json.Unmarshal(ReasonsString, &o.Reasons)
	}
	
	if OriginatingEntity, ok := ConversationnormalizedmessageMap["originatingEntity"].(string); ok {
		o.OriginatingEntity = &OriginatingEntity
	}
    
	if IsFinalReceipt, ok := ConversationnormalizedmessageMap["isFinalReceipt"].(bool); ok {
		o.IsFinalReceipt = &IsFinalReceipt
	}
    
	if Direction, ok := ConversationnormalizedmessageMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Metadata, ok := ConversationnormalizedmessageMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationnormalizedmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
