package platformclientv2
import (
	"encoding/json"
)

// Dialerpreview
type Dialerpreview struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ContactId - The contact associated with this preview data pop
	ContactId *string `json:"contactId,omitempty"`


	// ContactListId - The contactList associated with this preview data pop.
	ContactListId *string `json:"contactListId,omitempty"`


	// CampaignId - The campaignId associated with this preview data pop.
	CampaignId *string `json:"campaignId,omitempty"`


	// PhoneNumberColumns - The phone number columns associated with this campaign
	PhoneNumberColumns *[]Phonenumbercolumn `json:"phoneNumberColumns,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerpreview) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
