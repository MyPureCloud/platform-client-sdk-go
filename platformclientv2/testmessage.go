package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// Time - The time when the message was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`

}

func (u *Testmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Testmessage

	
	Time := new(string)
	if u.Time != nil {
		
		*Time = timeutil.Strftime(u.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		To *[]Emailaddress `json:"to,omitempty"`
		
		From *Emailaddress `json:"from,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		TextBody *string `json:"textBody,omitempty"`
		
		HtmlBody *string `json:"htmlBody,omitempty"`
		
		Time *string `json:"time,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		To: u.To,
		
		From: u.From,
		
		Subject: u.Subject,
		
		TextBody: u.TextBody,
		
		HtmlBody: u.HtmlBody,
		
		Time: Time,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Testmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
