package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentlessemailsendrequestdto
type Agentlessemailsendrequestdto struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SenderType - The direction of the message.
	SenderType *string `json:"senderType,omitempty"`

	// ConversationId - The identifier of the conversation. This must be an email interaction.
	ConversationId *string `json:"conversationId,omitempty"`

	// FromAddress - The sender of the message.
	FromAddress *Emailaddress `json:"fromAddress,omitempty"`

	// ToAddresses - The recipient of the message. We currently support one recipient only.
	ToAddresses *[]Emailaddress `json:"toAddresses,omitempty"`

	// ReplyToAddress - The address to use for reply.
	ReplyToAddress *Emailaddress `json:"replyToAddress,omitempty"`

	// Subject - The subject of the message.
	Subject *string `json:"subject,omitempty"`

	// TextBody - The Content of the message, in plain text.
	TextBody *string `json:"textBody,omitempty"`

	// HtmlBody - The Content of the message, in HTML. Links, images and styles are allowed
	HtmlBody *string `json:"htmlBody,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentlessemailsendrequestdto) SetField(field string, fieldValue interface{}) {
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

func (o Agentlessemailsendrequestdto) MarshalJSON() ([]byte, error) {
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
	type Alias Agentlessemailsendrequestdto
	
	return json.Marshal(&struct { 
		SenderType *string `json:"senderType,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		FromAddress *Emailaddress `json:"fromAddress,omitempty"`
		
		ToAddresses *[]Emailaddress `json:"toAddresses,omitempty"`
		
		ReplyToAddress *Emailaddress `json:"replyToAddress,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		TextBody *string `json:"textBody,omitempty"`
		
		HtmlBody *string `json:"htmlBody,omitempty"`
		Alias
	}{ 
		SenderType: o.SenderType,
		
		ConversationId: o.ConversationId,
		
		FromAddress: o.FromAddress,
		
		ToAddresses: o.ToAddresses,
		
		ReplyToAddress: o.ReplyToAddress,
		
		Subject: o.Subject,
		
		TextBody: o.TextBody,
		
		HtmlBody: o.HtmlBody,
		Alias:    (Alias)(o),
	})
}

func (o *Agentlessemailsendrequestdto) UnmarshalJSON(b []byte) error {
	var AgentlessemailsendrequestdtoMap map[string]interface{}
	err := json.Unmarshal(b, &AgentlessemailsendrequestdtoMap)
	if err != nil {
		return err
	}
	
	if SenderType, ok := AgentlessemailsendrequestdtoMap["senderType"].(string); ok {
		o.SenderType = &SenderType
	}
    
	if ConversationId, ok := AgentlessemailsendrequestdtoMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if FromAddress, ok := AgentlessemailsendrequestdtoMap["fromAddress"].(map[string]interface{}); ok {
		FromAddressString, _ := json.Marshal(FromAddress)
		json.Unmarshal(FromAddressString, &o.FromAddress)
	}
	
	if ToAddresses, ok := AgentlessemailsendrequestdtoMap["toAddresses"].([]interface{}); ok {
		ToAddressesString, _ := json.Marshal(ToAddresses)
		json.Unmarshal(ToAddressesString, &o.ToAddresses)
	}
	
	if ReplyToAddress, ok := AgentlessemailsendrequestdtoMap["replyToAddress"].(map[string]interface{}); ok {
		ReplyToAddressString, _ := json.Marshal(ReplyToAddress)
		json.Unmarshal(ReplyToAddressString, &o.ReplyToAddress)
	}
	
	if Subject, ok := AgentlessemailsendrequestdtoMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if TextBody, ok := AgentlessemailsendrequestdtoMap["textBody"].(string); ok {
		o.TextBody = &TextBody
	}
    
	if HtmlBody, ok := AgentlessemailsendrequestdtoMap["htmlBody"].(string); ok {
		o.HtmlBody = &HtmlBody
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentlessemailsendrequestdto) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
