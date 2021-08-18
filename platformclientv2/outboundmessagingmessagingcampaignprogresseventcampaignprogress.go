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

func (u *Outboundmessagingmessagingcampaignprogresseventcampaignprogress) MarshalJSON() ([]byte, error) {
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
		Campaign: u.Campaign,
		
		NumberOfContactsCalled: u.NumberOfContactsCalled,
		
		NumberOfContactsMessaged: u.NumberOfContactsMessaged,
		
		TotalNumberOfContacts: u.TotalNumberOfContacts,
		
		Percentage: u.Percentage,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignprogresseventcampaignprogress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
