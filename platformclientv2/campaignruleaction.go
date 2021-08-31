package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Campaignruleaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignruleaction
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Parameters *Campaignruleparameters `json:"parameters,omitempty"`
		
		ActionType *string `json:"actionType,omitempty"`
		
		CampaignRuleActionEntities *Campaignruleactionentities `json:"campaignRuleActionEntities,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Parameters: o.Parameters,
		
		ActionType: o.ActionType,
		
		CampaignRuleActionEntities: o.CampaignRuleActionEntities,
		Alias:    (*Alias)(o),
	})
}

func (o *Campaignruleaction) UnmarshalJSON(b []byte) error {
	var CampaignruleactionMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignruleactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CampaignruleactionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Parameters, ok := CampaignruleactionMap["parameters"].(map[string]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	
	if ActionType, ok := CampaignruleactionMap["actionType"].(string); ok {
		o.ActionType = &ActionType
	}
	
	if CampaignRuleActionEntities, ok := CampaignruleactionMap["campaignRuleActionEntities"].(map[string]interface{}); ok {
		CampaignRuleActionEntitiesString, _ := json.Marshal(CampaignRuleActionEntities)
		json.Unmarshal(CampaignRuleActionEntitiesString, &o.CampaignRuleActionEntities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignruleaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
