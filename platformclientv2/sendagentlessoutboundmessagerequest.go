package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Sendagentlessoutboundmessagerequest
type Sendagentlessoutboundmessagerequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// FromAddress - The messaging address of the sender of the message. For an SMS messenger type, this must be a currently provisioned SMS phone number. For a WhatsApp messenger type use the provisioned WhatsApp integration’s ID
	FromAddress *string `json:"fromAddress,omitempty"`

	// ToAddress - The messaging address of the recipient of the message. For an SMS messenger type, the phone number address must be in E.164 format. E.g. +13175555555 or +34234234234. For WhatsApp messenger type, use a WhatsApp ID of a phone number. E.g for a E.164 formatted phone number `+13175555555`, a WhatsApp ID would be 13175555555
	ToAddress *string `json:"toAddress,omitempty"`

	// ToAddressMessengerType - The recipient messaging address messenger type.
	ToAddressMessengerType *string `json:"toAddressMessengerType,omitempty"`

	// TextBody - The text of the message to send. This field is required in the case of SMS messenger type. Maximum character counts are: SMS - 765 characters, other channels - 2000 characters.
	TextBody *string `json:"textBody,omitempty"`

	// MessagingTemplate - The messaging template to use in the case of WhatsApp messenger type. This field is required when using WhatsApp messenger type
	MessagingTemplate *Sendmessagingtemplaterequest `json:"messagingTemplate,omitempty"`

	// UseExistingActiveConversation - Use an existing active conversation to send the agentless outbound message. Set this parameter to 'true' to use active conversation. Default value: false
	UseExistingActiveConversation *bool `json:"useExistingActiveConversation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Sendagentlessoutboundmessagerequest) SetField(field string, fieldValue interface{}) {
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

func (o Sendagentlessoutboundmessagerequest) MarshalJSON() ([]byte, error) {
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
	type Alias Sendagentlessoutboundmessagerequest
	
	return json.Marshal(&struct { 
		FromAddress *string `json:"fromAddress,omitempty"`
		
		ToAddress *string `json:"toAddress,omitempty"`
		
		ToAddressMessengerType *string `json:"toAddressMessengerType,omitempty"`
		
		TextBody *string `json:"textBody,omitempty"`
		
		MessagingTemplate *Sendmessagingtemplaterequest `json:"messagingTemplate,omitempty"`
		
		UseExistingActiveConversation *bool `json:"useExistingActiveConversation,omitempty"`
		Alias
	}{ 
		FromAddress: o.FromAddress,
		
		ToAddress: o.ToAddress,
		
		ToAddressMessengerType: o.ToAddressMessengerType,
		
		TextBody: o.TextBody,
		
		MessagingTemplate: o.MessagingTemplate,
		
		UseExistingActiveConversation: o.UseExistingActiveConversation,
		Alias:    (Alias)(o),
	})
}

func (o *Sendagentlessoutboundmessagerequest) UnmarshalJSON(b []byte) error {
	var SendagentlessoutboundmessagerequestMap map[string]interface{}
	err := json.Unmarshal(b, &SendagentlessoutboundmessagerequestMap)
	if err != nil {
		return err
	}
	
	if FromAddress, ok := SendagentlessoutboundmessagerequestMap["fromAddress"].(string); ok {
		o.FromAddress = &FromAddress
	}
    
	if ToAddress, ok := SendagentlessoutboundmessagerequestMap["toAddress"].(string); ok {
		o.ToAddress = &ToAddress
	}
    
	if ToAddressMessengerType, ok := SendagentlessoutboundmessagerequestMap["toAddressMessengerType"].(string); ok {
		o.ToAddressMessengerType = &ToAddressMessengerType
	}
    
	if TextBody, ok := SendagentlessoutboundmessagerequestMap["textBody"].(string); ok {
		o.TextBody = &TextBody
	}
    
	if MessagingTemplate, ok := SendagentlessoutboundmessagerequestMap["messagingTemplate"].(map[string]interface{}); ok {
		MessagingTemplateString, _ := json.Marshal(MessagingTemplate)
		json.Unmarshal(MessagingTemplateString, &o.MessagingTemplate)
	}
	
	if UseExistingActiveConversation, ok := SendagentlessoutboundmessagerequestMap["useExistingActiveConversation"].(bool); ok {
		o.UseExistingActiveConversation = &UseExistingActiveConversation
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Sendagentlessoutboundmessagerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
