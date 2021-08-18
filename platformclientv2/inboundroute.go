package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Inboundroute
type Inboundroute struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Pattern - The search pattern that the mailbox name should match.
	Pattern *string `json:"pattern,omitempty"`


	// Queue - The queue to route the emails to.
	Queue *Domainentityref `json:"queue,omitempty"`


	// Priority - The priority to use for routing.
	Priority *int `json:"priority,omitempty"`


	// Skills - The skills to use for routing.
	Skills *[]Domainentityref `json:"skills,omitempty"`


	// Language - The language to use for routing.
	Language *Domainentityref `json:"language,omitempty"`


	// FromName - The sender name to use for outgoing replies.
	FromName *string `json:"fromName,omitempty"`


	// FromEmail - The sender email to use for outgoing replies.
	FromEmail *string `json:"fromEmail,omitempty"`


	// Flow - The flow to use for processing the email.
	Flow *Domainentityref `json:"flow,omitempty"`


	// ReplyEmailAddress - The route to use for email replies.
	ReplyEmailAddress **Queueemailaddress `json:"replyEmailAddress,omitempty"`


	// AutoBcc - The recipients that should be  automatically blind copied on outbound emails associated with this InboundRoute.
	AutoBcc *[]Emailaddress `json:"autoBcc,omitempty"`


	// SpamFlow - The flow to use for processing inbound emails that have been marked as spam.
	SpamFlow *Domainentityref `json:"spamFlow,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Inboundroute) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Inboundroute

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Pattern *string `json:"pattern,omitempty"`
		
		Queue *Domainentityref `json:"queue,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Domainentityref `json:"skills,omitempty"`
		
		Language *Domainentityref `json:"language,omitempty"`
		
		FromName *string `json:"fromName,omitempty"`
		
		FromEmail *string `json:"fromEmail,omitempty"`
		
		Flow *Domainentityref `json:"flow,omitempty"`
		
		ReplyEmailAddress **Queueemailaddress `json:"replyEmailAddress,omitempty"`
		
		AutoBcc *[]Emailaddress `json:"autoBcc,omitempty"`
		
		SpamFlow *Domainentityref `json:"spamFlow,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Pattern: u.Pattern,
		
		Queue: u.Queue,
		
		Priority: u.Priority,
		
		Skills: u.Skills,
		
		Language: u.Language,
		
		FromName: u.FromName,
		
		FromEmail: u.FromEmail,
		
		Flow: u.Flow,
		
		ReplyEmailAddress: u.ReplyEmailAddress,
		
		AutoBcc: u.AutoBcc,
		
		SpamFlow: u.SpamFlow,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Inboundroute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
