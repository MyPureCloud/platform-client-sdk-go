package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Gdprsubject
type Gdprsubject struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// ExternalContactId
	ExternalContactId *string `json:"externalContactId,omitempty"`


	// DialerContactId
	DialerContactId *Dialercontactid `json:"dialerContactId,omitempty"`


	// JourneyCustomer
	JourneyCustomer *Gdprjourneycustomer `json:"journeyCustomer,omitempty"`


	// SocialHandle
	SocialHandle *Socialhandle `json:"socialHandle,omitempty"`


	// Addresses
	Addresses *[]string `json:"addresses,omitempty"`


	// PhoneNumbers
	PhoneNumbers *[]string `json:"phoneNumbers,omitempty"`


	// EmailAddresses
	EmailAddresses *[]string `json:"emailAddresses,omitempty"`

}

// String returns a JSON representation of the model
func (o *Gdprsubject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
