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

func (o *Recordingemailmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingemailmessage
	
	Time := new(string)
	if o.Time != nil {
		
		*Time = timeutil.Strftime(o.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		HtmlBody: o.HtmlBody,
		
		TextBody: o.TextBody,
		
		Id: o.Id,
		
		To: o.To,
		
		Cc: o.Cc,
		
		Bcc: o.Bcc,
		
		From: o.From,
		
		Subject: o.Subject,
		
		Attachments: o.Attachments,
		
		Time: Time,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingemailmessage) UnmarshalJSON(b []byte) error {
	var RecordingemailmessageMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingemailmessageMap)
	if err != nil {
		return err
	}
	
	if HtmlBody, ok := RecordingemailmessageMap["htmlBody"].(string); ok {
		o.HtmlBody = &HtmlBody
	}
    
	if TextBody, ok := RecordingemailmessageMap["textBody"].(string); ok {
		o.TextBody = &TextBody
	}
    
	if Id, ok := RecordingemailmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if To, ok := RecordingemailmessageMap["to"].([]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if Cc, ok := RecordingemailmessageMap["cc"].([]interface{}); ok {
		CcString, _ := json.Marshal(Cc)
		json.Unmarshal(CcString, &o.Cc)
	}
	
	if Bcc, ok := RecordingemailmessageMap["bcc"].([]interface{}); ok {
		BccString, _ := json.Marshal(Bcc)
		json.Unmarshal(BccString, &o.Bcc)
	}
	
	if From, ok := RecordingemailmessageMap["from"].(map[string]interface{}); ok {
		FromString, _ := json.Marshal(From)
		json.Unmarshal(FromString, &o.From)
	}
	
	if Subject, ok := RecordingemailmessageMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if Attachments, ok := RecordingemailmessageMap["attachments"].([]interface{}); ok {
		AttachmentsString, _ := json.Marshal(Attachments)
		json.Unmarshal(AttachmentsString, &o.Attachments)
	}
	
	if timeString, ok := RecordingemailmessageMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingemailmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
