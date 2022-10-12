package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentlessemailsendresponsedto
type Agentlessemailsendresponsedto struct { 
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

func (o *Agentlessemailsendresponsedto) MarshalJSON() ([]byte, error) {
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
		*Alias
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
		Alias:    (*Alias)(o),
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
