package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignruleconfigchangecampaignrulecondition
type Dialercampaignruleconfigchangecampaignrulecondition struct { 
	// Id - The globally unique identifier for the condition
	Id *string `json:"id,omitempty"`


	// Parameters - The parameters to match this condition
	Parameters *map[string]string `json:"parameters,omitempty"`


	// ConditionType - The type of this condition
	ConditionType *string `json:"conditionType,omitempty"`

}

func (o *Dialercampaignruleconfigchangecampaignrulecondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignruleconfigchangecampaignrulecondition
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Parameters *map[string]string `json:"parameters,omitempty"`
		
		ConditionType *string `json:"conditionType,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Parameters: o.Parameters,
		
		ConditionType: o.ConditionType,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignruleconfigchangecampaignrulecondition) UnmarshalJSON(b []byte) error {
	var DialercampaignruleconfigchangecampaignruleconditionMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignruleconfigchangecampaignruleconditionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialercampaignruleconfigchangecampaignruleconditionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Parameters, ok := DialercampaignruleconfigchangecampaignruleconditionMap["parameters"].(map[string]interface{}); ok {
		ParametersString, _ := json.Marshal(Parameters)
		json.Unmarshal(ParametersString, &o.Parameters)
	}
	
	if ConditionType, ok := DialercampaignruleconfigchangecampaignruleconditionMap["conditionType"].(string); ok {
		o.ConditionType = &ConditionType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignrulecondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
