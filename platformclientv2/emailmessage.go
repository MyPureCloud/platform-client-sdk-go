package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailmessage
type Emailmessage struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// To - The recipients of the email message.
	To *[]Emailaddress `json:"to,omitempty"`


	// Cc - The recipients that were copied on the email message.
	Cc *[]Emailaddress `json:"cc,omitempty"`


	// Bcc - The recipients that were blind copied on the email message.
	Bcc *[]Emailaddress `json:"bcc,omitempty"`


	// From - The sender of the email message.
	From *Emailaddress `json:"from,omitempty"`


	// ReplyTo - The receiver of the reply email message.
	ReplyTo *Emailaddress `json:"replyTo,omitempty"`


	// Subject - The subject of the email message.
	Subject *string `json:"subject,omitempty"`


	// Attachments - The attachments of the email message.
	Attachments *[]Attachment `json:"attachments,omitempty"`


	// TextBody - The text body of the email message.
	TextBody *string `json:"textBody,omitempty"`


	// HtmlBody - The html body of the email message.
	HtmlBody *string `json:"htmlBody,omitempty"`


	// Time - The time when the message was received or sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`


	// HistoryIncluded - Indicates whether the history of previous emails of the conversation is included within the email bodies of this message.
	HistoryIncluded *bool `json:"historyIncluded,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Emailmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailmessage

	
	Time := new(string)
	if u.Time != nil {
		
		*Time = timeutil.Strftime(u.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		To *[]Emailaddress `json:"to,omitempty"`
		
		Cc *[]Emailaddress `json:"cc,omitempty"`
		
		Bcc *[]Emailaddress `json:"bcc,omitempty"`
		
		From *Emailaddress `json:"from,omitempty"`
		
		ReplyTo *Emailaddress `json:"replyTo,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		Attachments *[]Attachment `json:"attachments,omitempty"`
		
		TextBody *string `json:"textBody,omitempty"`
		
		HtmlBody *string `json:"htmlBody,omitempty"`
		
		Time *string `json:"time,omitempty"`
		
		HistoryIncluded *bool `json:"historyIncluded,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		To: u.To,
		
		Cc: u.Cc,
		
		Bcc: u.Bcc,
		
		From: u.From,
		
		ReplyTo: u.ReplyTo,
		
		Subject: u.Subject,
		
		Attachments: u.Attachments,
		
		TextBody: u.TextBody,
		
		HtmlBody: u.HtmlBody,
		
		Time: Time,
		
		HistoryIncluded: u.HistoryIncluded,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Emailmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
