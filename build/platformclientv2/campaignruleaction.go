package platformclientv2
import (
	"encoding/json"
)

// Campaignruleaction
type Campaignruleaction struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Parameters - The parameters for the CampaignRuleAction. Required for certain actionTypes.
	Parameters *Campaignruleparameters `json:"parameters,omitempty"`


	// ActionType - The action to take on the campaignRuleActionEntities.
	ActionType *string `json:"actionType,omitempty"`


	// CampaignRuleActionEntities - The list of entities that this action will apply to.
	CampaignRuleActionEntities *Campaignruleactionentities `json:"campaignRuleActionEntities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaignruleaction) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
