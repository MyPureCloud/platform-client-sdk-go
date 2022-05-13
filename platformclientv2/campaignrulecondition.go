package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Campaignrulecondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignrulecondition
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Parameters *Campaignruleparameters `json:"parameters,omitempty"`
		
		ConditionType *string `json:"conditionType,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Parameters: o.Parameters,
		
		ConditionType: o.ConditionType,
		Alias:    (*Alias)(o),
	})
}

func (o *Campaignrulecondition) UnmarshalJSON(b []byte) error {
	var CampaignruleconditionMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignruleconditionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CampaignruleconditionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Parameters, ok := CampaignruleconditionMap["parameters"].(map[string]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	
	if ConditionType, ok := CampaignruleconditionMap["conditionType"].(string); ok {
		o.ConditionType = &ConditionType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignrulecondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
