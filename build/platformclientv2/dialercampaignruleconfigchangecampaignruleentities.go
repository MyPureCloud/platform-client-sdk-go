package platformclientv2
import (
	"encoding/json"
)

// Dialercampaignruleconfigchangecampaignruleentities
type Dialercampaignruleconfigchangecampaignruleentities struct { 
	// Campaigns
	Campaigns *[]Dialercampaignruleconfigchangeurireference `json:"campaigns,omitempty"`


	// Sequences
	Sequences *[]Dialercampaignruleconfigchangeurireference `json:"sequences,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignruleentities) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
