package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Inboundmessagerequest
type Inboundmessagerequest struct { 
	// QueueId - The ID of the queue to use for routing the email conversation. This field is mutually exclusive with flowId
	QueueId *string `json:"queueId,omitempty"`


	// FlowId - The ID of the flow to use for routing email conversation. This field is mutually exclusive with queueId
	FlowId *string `json:"flowId,omitempty"`


	// Provider - The name of the provider that is sourcing the email such as Oracle, Salesforce, etc.
	Provider *string `json:"provider,omitempty"`


	// SkillIds - The list of skill ID's to use for routing.
	SkillIds *[]string `json:"skillIds,omitempty"`


	// LanguageId - The ID of the language to use for routing.
	LanguageId *string `json:"languageId,omitempty"`


	// Priority - The priority to assign to the conversation for routing.
	Priority *int `json:"priority,omitempty"`


	// Attributes - The list of attributes to associate with the customer participant.
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ToAddress - The email address of the recipient of the email.
	ToAddress *string `json:"toAddress,omitempty"`


	// ToName - The name of the recipient of the email.
	ToName *string `json:"toName,omitempty"`


	// FromAddress - The email address of the sender of the email.
	FromAddress *string `json:"fromAddress,omitempty"`


	// FromName - The name of the sender of the email.
	FromName *string `json:"fromName,omitempty"`


	// Subject - The subject of the email
	Subject *string `json:"subject,omitempty"`

}

func (u *Inboundmessagerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Inboundmessagerequest

	

	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		FlowId *string `json:"flowId,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ToAddress *string `json:"toAddress,omitempty"`
		
		ToName *string `json:"toName,omitempty"`
		
		FromAddress *string `json:"fromAddress,omitempty"`
		
		FromName *string `json:"fromName,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		*Alias
	}{ 
		QueueId: u.QueueId,
		
		FlowId: u.FlowId,
		
		Provider: u.Provider,
		
		SkillIds: u.SkillIds,
		
		LanguageId: u.LanguageId,
		
		Priority: u.Priority,
		
		Attributes: u.Attributes,
		
		ToAddress: u.ToAddress,
		
		ToName: u.ToName,
		
		FromAddress: u.FromAddress,
		
		FromName: u.FromName,
		
		Subject: u.Subject,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Inboundmessagerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
