package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentlessemailsendresponsedto
type Agentlessemailsendresponsedto struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// ConversationId - The identifier of the conversation.
	ConversationId *string `json:"conversationId,omitempty"`

	// SenderType - The identifier of the external participant of the given conversation.
	SenderType *string `json:"senderType,omitempty"`

	// FromAddress - The sender of the message.
	FromAddress *Emailaddress `json:"fromAddress,omitempty"`

	// ToAddresses - The recipient(s) of the message.
	ToAddresses *[]Emailaddress `json:"toAddresses,omitempty"`

	// ReplyToAddress - The address to use for reply.
	ReplyToAddress *Emailaddress `json:"replyToAddress,omitempty"`

	// Subject - The subject of the message.
	Subject *string `json:"subject,omitempty"`

	// DateCreated - The message creation timestamp. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentlessemailsendresponsedto) SetField(field string, fieldValue interface{}) {
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

func (o Agentlessemailsendresponsedto) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
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
	type Alias Agentlessemailsendresponsedto
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		SenderType *string `json:"senderType,omitempty"`
		
		FromAddress *Emailaddress `json:"fromAddress,omitempty"`
		
		ToAddresses *[]Emailaddress `json:"toAddresses,omitempty"`
		
		ReplyToAddress *Emailaddress `json:"replyToAddress,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ConversationId: o.ConversationId,
		
		SenderType: o.SenderType,
		
		FromAddress: o.FromAddress,
		
		ToAddresses: o.ToAddresses,
		
		ReplyToAddress: o.ReplyToAddress,
		
		Subject: o.Subject,
		
		DateCreated: DateCreated,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Agentlessemailsendresponsedto) UnmarshalJSON(b []byte) error {
	var AgentlessemailsendresponsedtoMap map[string]interface{}
	err := json.Unmarshal(b, &AgentlessemailsendresponsedtoMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgentlessemailsendresponsedtoMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ConversationId, ok := AgentlessemailsendresponsedtoMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if SenderType, ok := AgentlessemailsendresponsedtoMap["senderType"].(string); ok {
		o.SenderType = &SenderType
	}
    
	if FromAddress, ok := AgentlessemailsendresponsedtoMap["fromAddress"].(map[string]interface{}); ok {
		FromAddressString, _ := json.Marshal(FromAddress)
		json.Unmarshal(FromAddressString, &o.FromAddress)
	}
	
	if ToAddresses, ok := AgentlessemailsendresponsedtoMap["toAddresses"].([]interface{}); ok {
		ToAddressesString, _ := json.Marshal(ToAddresses)
		json.Unmarshal(ToAddressesString, &o.ToAddresses)
	}
	
	if ReplyToAddress, ok := AgentlessemailsendresponsedtoMap["replyToAddress"].(map[string]interface{}); ok {
		ReplyToAddressString, _ := json.Marshal(ReplyToAddress)
		json.Unmarshal(ReplyToAddressString, &o.ReplyToAddress)
	}
	
	if Subject, ok := AgentlessemailsendresponsedtoMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if dateCreatedString, ok := AgentlessemailsendresponsedtoMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if SelfUri, ok := AgentlessemailsendresponsedtoMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentlessemailsendresponsedto) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
