package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Emailmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
