package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Emailconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
