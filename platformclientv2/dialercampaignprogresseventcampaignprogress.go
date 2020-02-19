package platformclientv2
import (
	"encoding/json"
)

// Dialercampaignprogresseventcampaignprogress
type Dialercampaignprogresseventcampaignprogress struct { 
	// Campaign
	Campaign *Dialercampaignprogresseventurireference `json:"campaign,omitempty"`


	// NumberOfContactsCalled
	NumberOfContactsCalled *float32 `json:"numberOfContactsCalled,omitempty"`


	// TotalNumberOfContacts
	TotalNumberOfContacts *float32 `json:"totalNumberOfContacts,omitempty"`


	// Percentage
	Percentage *int32 `json:"percentage,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignprogresseventcampaignprogress) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
