package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailconfig
type Emailconfig struct { 
	// EmailColumns - The contact list columns specifying the email address(es) of the contact.
	EmailColumns *[]string `json:"emailColumns,omitempty"`


	// ContentTemplate - The content template used to formulate the email to send to the contact.
	ContentTemplate *Domainentityref `json:"contentTemplate,omitempty"`


	// FromAddress - The email address that will be used as the sender of the email.
	FromAddress *Fromemailaddress `json:"fromAddress,omitempty"`


	// ReplyToAddress - The email address from which any reply will be sent.
	ReplyToAddress *Replytoemailaddress `json:"replyToAddress,omitempty"`

}

func (o *Emailconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Emailconfig
	
	return json.Marshal(&struct { 
		EmailColumns *[]string `json:"emailColumns,omitempty"`
		
		ContentTemplate *Domainentityref `json:"contentTemplate,omitempty"`
		
		FromAddress *Fromemailaddress `json:"fromAddress,omitempty"`
		
		ReplyToAddress *Replytoemailaddress `json:"replyToAddress,omitempty"`
		*Alias
	}{ 
		EmailColumns: o.EmailColumns,
		
		ContentTemplate: o.ContentTemplate,
		
		FromAddress: o.FromAddress,
		
		ReplyToAddress: o.ReplyToAddress,
		Alias:    (*Alias)(o),
	})
}

func (o *Emailconfig) UnmarshalJSON(b []byte) error {
	var EmailconfigMap map[string]interface{}
	err := json.Unmarshal(b, &EmailconfigMap)
	if err != nil {
		return err
	}
	
	if EmailColumns, ok := EmailconfigMap["emailColumns"].([]interface{}); ok {
		EmailColumnsString, _ := json.Marshal(EmailColumns)
		json.Unmarshal(EmailColumnsString, &o.EmailColumns)
	}
	
	if ContentTemplate, ok := EmailconfigMap["contentTemplate"].(map[string]interface{}); ok {
		ContentTemplateString, _ := json.Marshal(ContentTemplate)
		json.Unmarshal(ContentTemplateString, &o.ContentTemplate)
	}
	
	if FromAddress, ok := EmailconfigMap["fromAddress"].(map[string]interface{}); ok {
		FromAddressString, _ := json.Marshal(FromAddress)
		json.Unmarshal(FromAddressString, &o.FromAddress)
	}
	
	if ReplyToAddress, ok := EmailconfigMap["replyToAddress"].(map[string]interface{}); ok {
		ReplyToAddressString, _ := json.Marshal(ReplyToAddress)
		json.Unmarshal(ReplyToAddressString, &o.ReplyToAddress)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Emailconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
