package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignprogress
type Campaignprogress struct { 
	// Campaign - Identifier of the campaign
	Campaign *Domainentityref `json:"campaign,omitempty"`


	// ContactList - Identifier of the contact list
	ContactList *Domainentityref `json:"contactList,omitempty"`


	// NumberOfContactsCalled - Number of contacts called during the campaign
	NumberOfContactsCalled *int `json:"numberOfContactsCalled,omitempty"`


	// NumberOfContactsMessaged - Number of contacts messaged during the campaign
	NumberOfContactsMessaged *int `json:"numberOfContactsMessaged,omitempty"`


	// TotalNumberOfContacts - Total number of contacts in the campaign
	TotalNumberOfContacts *int `json:"totalNumberOfContacts,omitempty"`


	// Percentage - Percentage of contacts processed during the campaign
	Percentage *int `json:"percentage,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaignprogress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
