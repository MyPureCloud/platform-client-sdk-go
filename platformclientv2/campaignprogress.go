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

func (o *Campaignprogress) MarshalJSON() ([]byte, error) {
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
		Campaign: o.Campaign,
		
		ContactList: o.ContactList,
		
		NumberOfContactsCalled: o.NumberOfContactsCalled,
		
		NumberOfContactsMessaged: o.NumberOfContactsMessaged,
		
		TotalNumberOfContacts: o.TotalNumberOfContacts,
		
		Percentage: o.Percentage,
		Alias:    (*Alias)(o),
	})
}

func (o *Campaignprogress) UnmarshalJSON(b []byte) error {
	var CampaignprogressMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignprogressMap)
	if err != nil {
		return err
	}
	
	if Campaign, ok := CampaignprogressMap["campaign"].(map[string]interface{}); ok {
		CampaignString, _ := json.Marshal(Campaign)
		json.Unmarshal(CampaignString, &o.Campaign)
	}
	
	if ContactList, ok := CampaignprogressMap["contactList"].(map[string]interface{}); ok {
		ContactListString, _ := json.Marshal(ContactList)
		json.Unmarshal(ContactListString, &o.ContactList)
	}
	
	if NumberOfContactsCalled, ok := CampaignprogressMap["numberOfContactsCalled"].(float64); ok {
		NumberOfContactsCalledInt := int(NumberOfContactsCalled)
		o.NumberOfContactsCalled = &NumberOfContactsCalledInt
	}
	
	if NumberOfContactsMessaged, ok := CampaignprogressMap["numberOfContactsMessaged"].(float64); ok {
		NumberOfContactsMessagedInt := int(NumberOfContactsMessaged)
		o.NumberOfContactsMessaged = &NumberOfContactsMessagedInt
	}
	
	if TotalNumberOfContacts, ok := CampaignprogressMap["totalNumberOfContacts"].(float64); ok {
		TotalNumberOfContactsInt := int(TotalNumberOfContacts)
		o.TotalNumberOfContacts = &TotalNumberOfContactsInt
	}
	
	if Percentage, ok := CampaignprogressMap["percentage"].(float64); ok {
		PercentageInt := int(Percentage)
		o.Percentage = &PercentageInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignprogress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
