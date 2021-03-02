package platformclientv2
import (
	"encoding/json"
)

// Dialercampaignruleconfigchangecampaignruleactionentities
type Dialercampaignruleconfigchangecampaignruleactionentities struct { 
	// Campaigns
	Campaigns *[]Dialercampaignruleconfigchangeurireference `json:"campaigns,omitempty"`


	// Sequences
	Sequences *[]Dialercampaignruleconfigchangeurireference `json:"sequences,omitempty"`


	// UseTriggeringEntity
	UseTriggeringEntity *bool `json:"useTriggeringEntity,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignruleactionentities) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
