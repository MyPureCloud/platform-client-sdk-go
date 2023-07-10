package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailmessagepreview
type Emailmessagepreview struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// TextBodyPreview - A truncated version of the textBody
	TextBodyPreview *string `json:"textBodyPreview,omitempty"`

	// Time - The time when the message was received or sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`

	// HistoryIncluded - Indicates whether the history of previous emails of the conversation is included within the email bodies of this message.
	HistoryIncluded *bool `json:"historyIncluded,omitempty"`

	// State - The state of the current draft.
	State *string `json:"state,omitempty"`

	// DraftType - The type of draft that need to be treated.
	DraftType *string `json:"draftType,omitempty"`

	// EmailSizeBytes - Indicates an estimation of the size of the current email as a whole, in its final, ready to be sent form.
	EmailSizeBytes *int `json:"emailSizeBytes,omitempty"`

	// MaxEmailSizeBytes - Indicates the maximum allowed size for an email to be send via SMTP server, based on the email domain configuration
	MaxEmailSizeBytes *int `json:"maxEmailSizeBytes,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Emailmessagepreview) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Emailmessagepreview) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Time", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailmessagepreview
	
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
		
		TextBodyPreview *string `json:"textBodyPreview,omitempty"`
		
		Time *string `json:"time,omitempty"`
		
		HistoryIncluded *bool `json:"historyIncluded,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DraftType *string `json:"draftType,omitempty"`
		
		EmailSizeBytes *int `json:"emailSizeBytes,omitempty"`
		
		MaxEmailSizeBytes *int `json:"maxEmailSizeBytes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
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
		
		TextBodyPreview: o.TextBodyPreview,
		
		Time: Time,
		
		HistoryIncluded: o.HistoryIncluded,
		
		State: o.State,
		
		DraftType: o.DraftType,
		
		EmailSizeBytes: o.EmailSizeBytes,
		
		MaxEmailSizeBytes: o.MaxEmailSizeBytes,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Emailmessagepreview) UnmarshalJSON(b []byte) error {
	var EmailmessagepreviewMap map[string]interface{}
	err := json.Unmarshal(b, &EmailmessagepreviewMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EmailmessagepreviewMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EmailmessagepreviewMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if To, ok := EmailmessagepreviewMap["to"].([]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if Cc, ok := EmailmessagepreviewMap["cc"].([]interface{}); ok {
		CcString, _ := json.Marshal(Cc)
		json.Unmarshal(CcString, &o.Cc)
	}
	
	if Bcc, ok := EmailmessagepreviewMap["bcc"].([]interface{}); ok {
		BccString, _ := json.Marshal(Bcc)
		json.Unmarshal(BccString, &o.Bcc)
	}
	
	if From, ok := EmailmessagepreviewMap["from"].(map[string]interface{}); ok {
		FromString, _ := json.Marshal(From)
		json.Unmarshal(FromString, &o.From)
	}
	
	if ReplyTo, ok := EmailmessagepreviewMap["replyTo"].(map[string]interface{}); ok {
		ReplyToString, _ := json.Marshal(ReplyTo)
		json.Unmarshal(ReplyToString, &o.ReplyTo)
	}
	
	if Subject, ok := EmailmessagepreviewMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if Attachments, ok := EmailmessagepreviewMap["attachments"].([]interface{}); ok {
		AttachmentsString, _ := json.Marshal(Attachments)
		json.Unmarshal(AttachmentsString, &o.Attachments)
	}
	
	if TextBodyPreview, ok := EmailmessagepreviewMap["textBodyPreview"].(string); ok {
		o.TextBodyPreview = &TextBodyPreview
	}
    
	if timeString, ok := EmailmessagepreviewMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	
	if HistoryIncluded, ok := EmailmessagepreviewMap["historyIncluded"].(bool); ok {
		o.HistoryIncluded = &HistoryIncluded
	}
    
	if State, ok := EmailmessagepreviewMap["state"].(string); ok {
		o.State = &State
	}
    
	if DraftType, ok := EmailmessagepreviewMap["draftType"].(string); ok {
		o.DraftType = &DraftType
	}
    
	if EmailSizeBytes, ok := EmailmessagepreviewMap["emailSizeBytes"].(float64); ok {
		EmailSizeBytesInt := int(EmailSizeBytes)
		o.EmailSizeBytes = &EmailSizeBytesInt
	}
	
	if MaxEmailSizeBytes, ok := EmailmessagepreviewMap["maxEmailSizeBytes"].(float64); ok {
		MaxEmailSizeBytesInt := int(MaxEmailSizeBytes)
		o.MaxEmailSizeBytes = &MaxEmailSizeBytesInt
	}
	
	if SelfUri, ok := EmailmessagepreviewMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Emailmessagepreview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
