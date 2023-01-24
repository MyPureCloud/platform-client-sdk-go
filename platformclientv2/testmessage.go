package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Testmessage
type Testmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Testmessage) SetField(field string, fieldValue interface{}) {
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

func (o Testmessage) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Testmessage
	
	Time := new(string)
	if o.Time != nil {
		
		*Time = timeutil.Strftime(o.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Alias
	}{ 
		Id: o.Id,
		
		To: o.To,
		
		From: o.From,
		
		Subject: o.Subject,
		
		TextBody: o.TextBody,
		
		HtmlBody: o.HtmlBody,
		
		Time: Time,
		Alias:    (Alias)(o),
	})
}

func (o *Testmessage) UnmarshalJSON(b []byte) error {
	var TestmessageMap map[string]interface{}
	err := json.Unmarshal(b, &TestmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TestmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if To, ok := TestmessageMap["to"].([]interface{}); ok {
		ToString, _ := json.Marshal(To)
		json.Unmarshal(ToString, &o.To)
	}
	
	if From, ok := TestmessageMap["from"].(map[string]interface{}); ok {
		FromString, _ := json.Marshal(From)
		json.Unmarshal(FromString, &o.From)
	}
	
	if Subject, ok := TestmessageMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    
	if TextBody, ok := TestmessageMap["textBody"].(string); ok {
		o.TextBody = &TextBody
	}
    
	if HtmlBody, ok := TestmessageMap["htmlBody"].(string); ok {
		o.HtmlBody = &HtmlBody
	}
    
	if timeString, ok := TestmessageMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Testmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
