package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Smsconfig
type Smsconfig struct { 
	// MessageColumn - The Contact List column specifying the message to send to the contact.
	MessageColumn *string `json:"messageColumn,omitempty"`


	// PhoneColumn - The Contact List column specifying the phone number to send a message to.
	PhoneColumn *string `json:"phoneColumn,omitempty"`


	// SenderSmsPhoneNumber - A reference to the SMS Phone Number that will be used as the sender of a message.
	SenderSmsPhoneNumber *Smsphonenumberref `json:"senderSmsPhoneNumber,omitempty"`


	// ContentTemplate - The content template used to formulate the message to send to the contact.
	ContentTemplate *Domainentityref `json:"contentTemplate,omitempty"`

}

func (o *Smsconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Smsconfig
	
	return json.Marshal(&struct { 
		MessageColumn *string `json:"messageColumn,omitempty"`
		
		PhoneColumn *string `json:"phoneColumn,omitempty"`
		
		SenderSmsPhoneNumber *Smsphonenumberref `json:"senderSmsPhoneNumber,omitempty"`
		
		ContentTemplate *Domainentityref `json:"contentTemplate,omitempty"`
		*Alias
	}{ 
		MessageColumn: o.MessageColumn,
		
		PhoneColumn: o.PhoneColumn,
		
		SenderSmsPhoneNumber: o.SenderSmsPhoneNumber,
		
		ContentTemplate: o.ContentTemplate,
		Alias:    (*Alias)(o),
	})
}

func (o *Smsconfig) UnmarshalJSON(b []byte) error {
	var SmsconfigMap map[string]interface{}
	err := json.Unmarshal(b, &SmsconfigMap)
	if err != nil {
		return err
	}
	
	if MessageColumn, ok := SmsconfigMap["messageColumn"].(string); ok {
		o.MessageColumn = &MessageColumn
	}
    
	if PhoneColumn, ok := SmsconfigMap["phoneColumn"].(string); ok {
		o.PhoneColumn = &PhoneColumn
	}
    
	if SenderSmsPhoneNumber, ok := SmsconfigMap["senderSmsPhoneNumber"].(map[string]interface{}); ok {
		SenderSmsPhoneNumberString, _ := json.Marshal(SenderSmsPhoneNumber)
		json.Unmarshal(SenderSmsPhoneNumberString, &o.SenderSmsPhoneNumber)
	}
	
	if ContentTemplate, ok := SmsconfigMap["contentTemplate"].(map[string]interface{}); ok {
		ContentTemplateString, _ := json.Marshal(ContentTemplate)
		json.Unmarshal(ContentTemplateString, &o.ContentTemplate)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Smsconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
