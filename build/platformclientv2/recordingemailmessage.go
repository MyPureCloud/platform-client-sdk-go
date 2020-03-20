package platformclientv2
import (
	"time"
	"encoding/json"
)

// Recordingemailmessage
type Recordingemailmessage struct { 
	// HtmlBody
	HtmlBody *string `json:"htmlBody,omitempty"`


	// TextBody
	TextBody *string `json:"textBody,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// To
	To *[]Emailaddress `json:"to,omitempty"`


	// Cc
	Cc *[]Emailaddress `json:"cc,omitempty"`


	// Bcc
	Bcc *[]Emailaddress `json:"bcc,omitempty"`


	// From
	From *Emailaddress `json:"from,omitempty"`


	// Subject
	Subject *string `json:"subject,omitempty"`


	// Attachments
	Attachments *[]Emailattachment `json:"attachments,omitempty"`


	// Time - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	Time *time.Time `json:"time,omitempty"`

}

// String returns a JSON representation of the model
func (o *Recordingemailmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
