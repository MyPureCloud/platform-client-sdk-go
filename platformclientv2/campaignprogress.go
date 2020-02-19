package platformclientv2
import (
	"encoding/json"
)

// Campaignprogress
type Campaignprogress struct { 
	// Campaign - Identifier of the campaign
	Campaign *Domainentityref `json:"campaign,omitempty"`


	// ContactList - Identifier of the contact list
	ContactList *Domainentityref `json:"contactList,omitempty"`


	// NumberOfContactsCalled - Number of contacts processed during the campaign
	NumberOfContactsCalled *int64 `json:"numberOfContactsCalled,omitempty"`


	// TotalNumberOfContacts - Total number of contacts in the campaign
	TotalNumberOfContacts *int64 `json:"totalNumberOfContacts,omitempty"`


	// Percentage - Percentage of contacts processed during the campaign
	Percentage *int64 `json:"percentage,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaignprogress) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
