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

func (o *Emailmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailmessage
	
	Time := new(string)
	if o.Time != nil {
		
		*Time = timeutil.Strftime(o.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		To: o.To,
		
		Cc: o.Cc,
		
		Bcc: o.Bcc,
		
		From: o.From,
		
		ReplyTo: o.ReplyTo,
		
		Subject: o.Subject,
		
		Attachments: o.Attachments,
		
		TextBody: o.TextBody,
		
		HtmlBody: o.HtmlBody,
		
		Time: Time,
		
		HistoryIncluded: o.HistoryIncluded,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Emailmessage) UnmarshalJSON(b []byte) error {
	var EmailmessageMap map[string]interface{}
	err := json.Unmarshal(b, &EmailmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EmailmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EmailmessageMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if To, ok := EmailmessageMap["to"].([]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if Cc, ok := EmailmessageMap["cc"].([]interface{}); ok {
		CcString, _ := json.Marshal(Cc)
		json.Unmarshal(CcString, &o.Cc)
	}
	
	if Bcc, ok := EmailmessageMap["bcc"].([]interface{}); ok {
		BccString, _ := json.Marshal(Bcc)
		json.Unmarshal(BccString, &o.Bcc)
	}
	
	if From, ok := EmailmessageMap["from"].(map[string]interface{}); ok {
		FromString, _ := json.Marshal(From)
		json.Unmarshal(FromString, &o.From)
	}
	
	if ReplyTo, ok := EmailmessageMap["replyTo"].(map[string]interface{}); ok {
		ReplyToString, _ := json.Marshal(ReplyTo)
		json.Unmarshal(ReplyToString, &o.ReplyTo)
	}
	
	if Subject, ok := EmailmessageMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if Attachments, ok := EmailmessageMap["attachments"].([]interface{}); ok {
		AttachmentsString, _ := json.Marshal(Attachments)
		json.Unmarshal(AttachmentsString, &o.Attachments)
	}
	
	if TextBody, ok := EmailmessageMap["textBody"].(string); ok {
		o.TextBody = &TextBody
	}
    
	if HtmlBody, ok := EmailmessageMap["htmlBody"].(string); ok {
		o.HtmlBody = &HtmlBody
	}
    
	if timeString, ok := EmailmessageMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	
	if HistoryIncluded, ok := EmailmessageMap["historyIncluded"].(bool); ok {
		o.HistoryIncluded = &HistoryIncluded
	}
    
	if SelfUri, ok := EmailmessageMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Emailmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
