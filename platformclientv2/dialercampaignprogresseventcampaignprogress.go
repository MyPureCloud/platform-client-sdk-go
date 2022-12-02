package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignprogresseventcampaignprogress
type Dialercampaignprogresseventcampaignprogress struct { 
	// Campaign
	Campaign *Dialercampaignprogresseventurireference `json:"campaign,omitempty"`


	// NumberOfContactsCalled - The number of contacts that have been called so far
	NumberOfContactsCalled *float32 `json:"numberOfContactsCalled,omitempty"`


	// NumberOfContactsMessaged - The number of contacts that have been messaged so far
	NumberOfContactsMessaged *float32 `json:"numberOfContactsMessaged,omitempty"`


	// TotalNumberOfContacts - The total number of contacts in the contact list
	TotalNumberOfContacts *float32 `json:"totalNumberOfContacts,omitempty"`


	// Percentage - numberOfContactsContacted/totalNumberOfContacts*100
	Percentage *int `json:"percentage,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialercampaignprogresseventcampaignprogress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignprogresseventcampaignprogress
	
	return json.Marshal(&struct { 
		Campaign *Dialercampaignprogresseventurireference `json:"campaign,omitempty"`
		
		NumberOfContactsCalled *float32 `json:"numberOfContactsCalled,omitempty"`
		
		NumberOfContactsMessaged *float32 `json:"numberOfContactsMessaged,omitempty"`
		
		TotalNumberOfContacts *float32 `json:"totalNumberOfContacts,omitempty"`
		
		Percentage *int `json:"percentage,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Campaign: o.Campaign,
		
		NumberOfContactsCalled: o.NumberOfContactsCalled,
		
		NumberOfContactsMessaged: o.NumberOfContactsMessaged,
		
		TotalNumberOfContacts: o.TotalNumberOfContacts,
		
		Percentage: o.Percentage,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignprogresseventcampaignprogress) UnmarshalJSON(b []byte) error {
	var DialercampaignprogresseventcampaignprogressMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignprogresseventcampaignprogressMap)
	if err != nil {
		return err
	}
	
	if Campaign, ok := DialercampaignprogresseventcampaignprogressMap["campaign"].(map[string]interface{}); ok {
		CampaignString, _ := json.Marshal(Campaign)
		json.Unmarshal(CampaignString, &o.Campaign)
	}
	
	if NumberOfContactsCalled, ok := DialercampaignprogresseventcampaignprogressMap["numberOfContactsCalled"].(float64); ok {
		NumberOfContactsCalledFloat32 := float32(NumberOfContactsCalled)
		o.NumberOfContactsCalled = &NumberOfContactsCalledFloat32
	}
    
	if NumberOfContactsMessaged, ok := DialercampaignprogresseventcampaignprogressMap["numberOfContactsMessaged"].(float64); ok {
		NumberOfContactsMessagedFloat32 := float32(NumberOfContactsMessaged)
		o.NumberOfContactsMessaged = &NumberOfContactsMessagedFloat32
	}
    
	if TotalNumberOfContacts, ok := DialercampaignprogresseventcampaignprogressMap["totalNumberOfContacts"].(float64); ok {
		TotalNumberOfContactsFloat32 := float32(TotalNumberOfContacts)
		o.TotalNumberOfContacts = &TotalNumberOfContactsFloat32
	}
    
	if Percentage, ok := DialercampaignprogresseventcampaignprogressMap["percentage"].(float64); ok {
		PercentageInt := int(Percentage)
		o.Percentage = &PercentageInt
	}
	
	if AdditionalProperties, ok := DialercampaignprogresseventcampaignprogressMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignprogresseventcampaignprogress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
