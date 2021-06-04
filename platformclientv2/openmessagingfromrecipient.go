package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Openmessagingfromrecipient - Information about the recipient the message is received from.
type Openmessagingfromrecipient struct { 
	// Nickname - Nickname or display name of the recipient.
	Nickname *string `json:"nickname,omitempty"`


	// Id - The recipient ID specific to the provider.
	Id *string `json:"id,omitempty"`


	// IdType - The recipient ID type. This is used to indicate the format used for the ID.
	IdType *string `json:"idType,omitempty"`


	// FirstName - First name of the recipient.
	FirstName *string `json:"firstName,omitempty"`


	// LastName - Last name of the recipient.
	LastName *string `json:"lastName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Openmessagingfromrecipient) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
