package platformclientv2
import (
	"encoding/json"
)

// Campaignruleentities
type Campaignruleentities struct { 
	// Campaigns - The list of campaigns for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a campaign.
	Campaigns *[]Domainentityref `json:"campaigns,omitempty"`


	// Sequences - The list of sequences for a CampaignRule to monitor. Required if the CampaignRule has any conditions that run on a sequence.
	Sequences *[]Domainentityref `json:"sequences,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaignruleentities) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
