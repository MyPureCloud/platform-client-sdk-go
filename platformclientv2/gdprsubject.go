package platformclientv2
import (
	"github.com/leekchan/timeutil"
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


	// ExternalId
	ExternalId *string `json:"externalId,omitempty"`


	// Addresses
	Addresses *[]string `json:"addresses,omitempty"`


	// PhoneNumbers
	PhoneNumbers *[]string `json:"phoneNumbers,omitempty"`


	// EmailAddresses
	EmailAddresses *[]string `json:"emailAddresses,omitempty"`

}

func (u *Gdprsubject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Gdprsubject

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		DialerContactId *Dialercontactid `json:"dialerContactId,omitempty"`
		
		JourneyCustomer *Gdprjourneycustomer `json:"journeyCustomer,omitempty"`
		
		SocialHandle *Socialhandle `json:"socialHandle,omitempty"`
		
		ExternalId *string `json:"externalId,omitempty"`
		
		Addresses *[]string `json:"addresses,omitempty"`
		
		PhoneNumbers *[]string `json:"phoneNumbers,omitempty"`
		
		EmailAddresses *[]string `json:"emailAddresses,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		UserId: u.UserId,
		
		ExternalContactId: u.ExternalContactId,
		
		DialerContactId: u.DialerContactId,
		
		JourneyCustomer: u.JourneyCustomer,
		
		SocialHandle: u.SocialHandle,
		
		ExternalId: u.ExternalId,
		
		Addresses: u.Addresses,
		
		PhoneNumbers: u.PhoneNumbers,
		
		EmailAddresses: u.EmailAddresses,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Gdprsubject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
