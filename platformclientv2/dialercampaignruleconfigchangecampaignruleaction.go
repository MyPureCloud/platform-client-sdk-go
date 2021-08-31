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

func (o *Dialercampaignruleconfigchangecampaignruleaction) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Parameters: o.Parameters,
		
		ActionType: o.ActionType,
		
		CampaignRuleActionEntities: o.CampaignRuleActionEntities,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignruleconfigchangecampaignruleaction) UnmarshalJSON(b []byte) error {
	var DialercampaignruleconfigchangecampaignruleactionMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignruleconfigchangecampaignruleactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialercampaignruleconfigchangecampaignruleactionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Parameters, ok := DialercampaignruleconfigchangecampaignruleactionMap["parameters"].(map[string]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	
	if ActionType, ok := DialercampaignruleconfigchangecampaignruleactionMap["actionType"].(string); ok {
		o.ActionType = &ActionType
	}
	
	if CampaignRuleActionEntities, ok := DialercampaignruleconfigchangecampaignruleactionMap["campaignRuleActionEntities"].(map[string]interface{}); ok {
		CampaignRuleActionEntitiesString, _ := json.Marshal(CampaignRuleActionEntities)
		json.Unmarshal(CampaignRuleActionEntitiesString, &o.CampaignRuleActionEntities)
	}
	
	if AdditionalProperties, ok := DialercampaignruleconfigchangecampaignruleactionMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignruleaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
