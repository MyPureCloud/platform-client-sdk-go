package platformclientv2
import (
	"encoding/json"
)

// Dialercampaignruleconfigchangecampaignruleaction
type Dialercampaignruleconfigchangecampaignruleaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Parameters
	Parameters *map[string]string `json:"parameters,omitempty"`


	// ActionType
	ActionType *string `json:"actionType,omitempty"`


	// CampaignRuleActionEntities
	CampaignRuleActionEntities *Dialercampaignruleconfigchangecampaignruleactionentities `json:"campaignRuleActionEntities,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignruleaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
