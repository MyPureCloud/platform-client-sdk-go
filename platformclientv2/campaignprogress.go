package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Campaignprogress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignprogress

	

	return json.Marshal(&struct { 
		Campaign *Domainentityref `json:"campaign,omitempty"`
		
		ContactList *Domainentityref `json:"contactList,omitempty"`
		
		NumberOfContactsCalled *int `json:"numberOfContactsCalled,omitempty"`
		
		NumberOfContactsMessaged *int `json:"numberOfContactsMessaged,omitempty"`
		
		TotalNumberOfContacts *int `json:"totalNumberOfContacts,omitempty"`
		
		Percentage *int `json:"percentage,omitempty"`
		*Alias
	}{ 
		Campaign: u.Campaign,
		
		ContactList: u.ContactList,
		
		NumberOfContactsCalled: u.NumberOfContactsCalled,
		
		NumberOfContactsMessaged: u.NumberOfContactsMessaged,
		
		TotalNumberOfContacts: u.TotalNumberOfContacts,
		
		Percentage: u.Percentage,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Campaignprogress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
