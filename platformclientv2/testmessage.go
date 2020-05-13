package platformclientv2
import (
	"time"
	"encoding/json"
)

// Testmessage
type Testmessage struct { 
	// Id - After the message has been sent, this is the value of the Message-ID email header.
	Id *string `json:"id,omitempty"`


	// To - The recipients of the email message.
	To *[]Emailaddress `json:"to,omitempty"`


	// From - The sender of the email message.
	From *Emailaddress `json:"from,omitempty"`


	// Subject - The subject of the email message.
	Subject *string `json:"subject,omitempty"`


	// TextBody - The text body of the email message.
	TextBody *string `json:"textBody,omitempty"`


	// HtmlBody - The html body of the email message
	HtmlBody *string `json:"htmlBody,omitempty"`


	// Time - The time when the message was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	Time *time.Time `json:"time,omitempty"`

}

// String returns a JSON representation of the model
func (o *Testmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
