package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialercampaignruleconfigchangecampaignruleaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignruleconfigchangecampaignruleaction

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Parameters *map[string]string `json:"parameters,omitempty"`
		
		ActionType *string `json:"actionType,omitempty"`
		
		CampaignRuleActionEntities *Dialercampaignruleconfigchangecampaignruleactionentities `json:"campaignRuleActionEntities,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Parameters: u.Parameters,
		
		ActionType: u.ActionType,
		
		CampaignRuleActionEntities: u.CampaignRuleActionEntities,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignruleaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
