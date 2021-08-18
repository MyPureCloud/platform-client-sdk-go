package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// Time - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`

}

func (u *Recordingemailmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingemailmessage

	
	Time := new(string)
	if u.Time != nil {
		
		*Time = timeutil.Strftime(u.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	

	return json.Marshal(&struct { 
		HtmlBody *string `json:"htmlBody,omitempty"`
		
		TextBody *string `json:"textBody,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		To *[]Emailaddress `json:"to,omitempty"`
		
		Cc *[]Emailaddress `json:"cc,omitempty"`
		
		Bcc *[]Emailaddress `json:"bcc,omitempty"`
		
		From *Emailaddress `json:"from,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		Attachments *[]Emailattachment `json:"attachments,omitempty"`
		
		Time *string `json:"time,omitempty"`
		*Alias
	}{ 
		HtmlBody: u.HtmlBody,
		
		TextBody: u.TextBody,
		
		Id: u.Id,
		
		To: u.To,
		
		Cc: u.Cc,
		
		Bcc: u.Bcc,
		
		From: u.From,
		
		Subject: u.Subject,
		
		Attachments: u.Attachments,
		
		Time: Time,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Recordingemailmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
