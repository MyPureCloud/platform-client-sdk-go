package platformclientv2
import (
	"encoding/json"
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
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignprogresseventcampaignprogress) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
