package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createoutboundmessagingconversationrequest
type Createoutboundmessagingconversationrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QueueId - The ID of the queue to be associated with the message. This will determine the fromAddress of the message.
	QueueId *string `json:"queueId,omitempty"`

	// ToAddress - The messaging address of the recipient of the message. For an SMS messenger type, the phone number address must be in E.164 format. E.g. +13175555555 or +34234234234.  For open messenger type, any string within the outbound.open.messaging.to.address.characters.max limit can be used. For whatsapp messenger type, use a Whatsapp ID of a phone number. E.g for a E.164 formatted phone number `+13175555555`, a Whatsapp ID would be 13175555555
	ToAddress *string `json:"toAddress,omitempty"`

	// ToAddressMessengerType - The messaging address messenger type.
	ToAddressMessengerType *string `json:"toAddressMessengerType,omitempty"`

	// UseExistingConversation - An override to use an existing conversation.  If set to true, an existing conversation will be used if there is one within the conversation window.  If set to false, create request fails if there is a conversation within the conversation window.
	UseExistingConversation *bool `json:"useExistingConversation,omitempty"`

	// ExternalContactId - The external contact with which the message will be associated.
	ExternalContactId *string `json:"externalContactId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createoutboundmessagingconversationrequest) SetField(field string, fieldValue interface{}) {
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

func (o Createoutboundmessagingconversationrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createoutboundmessagingconversationrequest
	
	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		ToAddress *string `json:"toAddress,omitempty"`
		
		ToAddressMessengerType *string `json:"toAddressMessengerType,omitempty"`
		
		UseExistingConversation *bool `json:"useExistingConversation,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		Alias
	}{ 
		QueueId: o.QueueId,
		
		ToAddress: o.ToAddress,
		
		ToAddressMessengerType: o.ToAddressMessengerType,
		
		UseExistingConversation: o.UseExistingConversation,
		
		ExternalContactId: o.ExternalContactId,
		Alias:    (Alias)(o),
	})
}

func (o *Createoutboundmessagingconversationrequest) UnmarshalJSON(b []byte) error {
	var CreateoutboundmessagingconversationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreateoutboundmessagingconversationrequestMap)
	if err != nil {
		return err
	}
	
	if QueueId, ok := CreateoutboundmessagingconversationrequestMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if ToAddress, ok := CreateoutboundmessagingconversationrequestMap["toAddress"].(string); ok {
		o.ToAddress = &ToAddress
	}
    
	if ToAddressMessengerType, ok := CreateoutboundmessagingconversationrequestMap["toAddressMessengerType"].(string); ok {
		o.ToAddressMessengerType = &ToAddressMessengerType
	}
    
	if UseExistingConversation, ok := CreateoutboundmessagingconversationrequestMap["useExistingConversation"].(bool); ok {
		o.UseExistingConversation = &UseExistingConversation
	}
    
	if ExternalContactId, ok := CreateoutboundmessagingconversationrequestMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createoutboundmessagingconversationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
