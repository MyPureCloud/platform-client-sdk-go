package platformclientv2
import (
	"encoding/json"
)

// Campaignrulecondition
type Campaignrulecondition struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Parameters - The parameters for the CampaignRuleCondition.
	Parameters *Campaignruleparameters `json:"parameters,omitempty"`


	// ConditionType - The type of condition to evaluate.
	ConditionType *string `json:"conditionType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Campaignrulecondition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
