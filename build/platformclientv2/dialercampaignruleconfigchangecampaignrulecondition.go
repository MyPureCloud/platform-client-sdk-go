package platformclientv2
import (
	"encoding/json"
)

// Dialercampaignruleconfigchangecampaignrulecondition
type Dialercampaignruleconfigchangecampaignrulecondition struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Parameters
	Parameters *map[string]string `json:"parameters,omitempty"`


	// ConditionType
	ConditionType *string `json:"conditionType,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignrulecondition) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
