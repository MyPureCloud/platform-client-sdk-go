package platformclientv2
import (
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

}

// String returns a JSON representation of the model
func (o *Smsconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
