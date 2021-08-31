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

func (o *Gdprsubject) MarshalJSON() ([]byte, error) {
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
		Name: o.Name,
		
		UserId: o.UserId,
		
		ExternalContactId: o.ExternalContactId,
		
		DialerContactId: o.DialerContactId,
		
		JourneyCustomer: o.JourneyCustomer,
		
		SocialHandle: o.SocialHandle,
		
		ExternalId: o.ExternalId,
		
		Addresses: o.Addresses,
		
		PhoneNumbers: o.PhoneNumbers,
		
		EmailAddresses: o.EmailAddresses,
		Alias:    (*Alias)(o),
	})
}

func (o *Gdprsubject) UnmarshalJSON(b []byte) error {
	var GdprsubjectMap map[string]interface{}
	err := json.Unmarshal(b, &GdprsubjectMap)
	if err != nil {
		return err
	}
	
	if Name, ok := GdprsubjectMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if UserId, ok := GdprsubjectMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	
	if ExternalContactId, ok := GdprsubjectMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
	
	if DialerContactId, ok := GdprsubjectMap["dialerContactId"].(map[string]interface{}); ok {
		DialerContactIdString, _ := json.Marshal(DialerContactId)
		json.Unmarshal(DialerContactIdString, &o.DialerContactId)
	}
	
	if JourneyCustomer, ok := GdprsubjectMap["journeyCustomer"].(map[string]interface{}); ok {
		JourneyCustomerString, _ := json.Marshal(JourneyCustomer)
		json.Unmarshal(JourneyCustomerString, &o.JourneyCustomer)
	}
	
	if SocialHandle, ok := GdprsubjectMap["socialHandle"].(map[string]interface{}); ok {
		SocialHandleString, _ := json.Marshal(SocialHandle)
		json.Unmarshal(SocialHandleString, &o.SocialHandle)
	}
	
	if ExternalId, ok := GdprsubjectMap["externalId"].(string); ok {
		o.ExternalId = &ExternalId
	}
	
	if Addresses, ok := GdprsubjectMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if PhoneNumbers, ok := GdprsubjectMap["phoneNumbers"].([]interface{}); ok {
		PhoneNumbersString, _ := json.Marshal(PhoneNumbers)
		json.Unmarshal(PhoneNumbersString, &o.PhoneNumbers)
	}
	
	if EmailAddresses, ok := GdprsubjectMap["emailAddresses"].([]interface{}); ok {
		EmailAddressesString, _ := json.Marshal(EmailAddresses)
		json.Unmarshal(EmailAddressesString, &o.EmailAddresses)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Gdprsubject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
