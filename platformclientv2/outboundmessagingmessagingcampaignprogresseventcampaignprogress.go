package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignprogresseventcampaignprogress
type Outboundmessagingmessagingcampaignprogresseventcampaignprogress struct { 
	// Campaign
	Campaign *Outboundmessagingmessagingcampaignprogresseventurireference `json:"campaign,omitempty"`


	// NumberOfContactsCalled
	NumberOfContactsCalled *float32 `json:"numberOfContactsCalled,omitempty"`


	// NumberOfContactsMessaged
	NumberOfContactsMessaged *float32 `json:"numberOfContactsMessaged,omitempty"`


	// TotalNumberOfContacts
	TotalNumberOfContacts *float32 `json:"totalNumberOfContacts,omitempty"`


	// Percentage
	Percentage *int `json:"percentage,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Outboundmessagingmessagingcampaignprogresseventcampaignprogress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignprogresseventcampaignprogress
	
	return json.Marshal(&struct { 
		Campaign *Outboundmessagingmessagingcampaignprogresseventurireference `json:"campaign,omitempty"`
		
		NumberOfContactsCalled *float32 `json:"numberOfContactsCalled,omitempty"`
		
		NumberOfContactsMessaged *float32 `json:"numberOfContactsMessaged,omitempty"`
		
		TotalNumberOfContacts *float32 `json:"totalNumberOfContacts,omitempty"`
		
		Percentage *int `json:"percentage,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
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

func (o *Outboundmessagingmessagingcampaignprogresseventcampaignprogress) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap)
	if err != nil {
		return err
	}
	
	if Campaign, ok := OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap["campaign"].(map[string]interface{}); ok {
		CampaignString, _ := json.Marshal(Campaign)
		json.Unmarshal(CampaignString, &o.Campaign)
	}
	
	if NumberOfContactsCalled, ok := OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap["numberOfContactsCalled"].(float64); ok {
		NumberOfContactsCalledFloat32 := float32(NumberOfContactsCalled)
		o.NumberOfContactsCalled = &NumberOfContactsCalledFloat32
	}
	
	if NumberOfContactsMessaged, ok := OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap["numberOfContactsMessaged"].(float64); ok {
		NumberOfContactsMessagedFloat32 := float32(NumberOfContactsMessaged)
		o.NumberOfContactsMessaged = &NumberOfContactsMessagedFloat32
	}
	
	if TotalNumberOfContacts, ok := OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap["totalNumberOfContacts"].(float64); ok {
		TotalNumberOfContactsFloat32 := float32(TotalNumberOfContacts)
		o.TotalNumberOfContacts = &TotalNumberOfContactsFloat32
	}
	
	if Percentage, ok := OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap["percentage"].(float64); ok {
		PercentageInt := int(Percentage)
		o.Percentage = &PercentageInt
	}
	
	if AdditionalProperties, ok := OutboundmessagingmessagingcampaignprogresseventcampaignprogressMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignprogresseventcampaignprogress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
