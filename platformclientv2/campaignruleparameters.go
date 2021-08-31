package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignruleparameters
type Campaignruleparameters struct { 
	// Operator - The operator for comparison. Required for a CampaignRuleCondition.
	Operator *string `json:"operator,omitempty"`


	// Value - The value for comparison. Required for a CampaignRuleCondition.
	Value *string `json:"value,omitempty"`


	// Priority - The priority to set a campaign to. Required for the 'setCampaignPriority' action.
	Priority *string `json:"priority,omitempty"`


	// DialingMode - The dialing mode to set a campaign to. Required for the 'setCampaignDialingMode' action.
	DialingMode *string `json:"dialingMode,omitempty"`

}

func (o *Campaignruleparameters) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignruleparameters
	
	return json.Marshal(&struct { 
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Priority *string `json:"priority,omitempty"`
		
		DialingMode *string `json:"dialingMode,omitempty"`
		*Alias
	}{ 
		Operator: o.Operator,
		
		Value: o.Value,
		
		Priority: o.Priority,
		
		DialingMode: o.DialingMode,
		Alias:    (*Alias)(o),
	})
}

func (o *Campaignruleparameters) UnmarshalJSON(b []byte) error {
	var CampaignruleparametersMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignruleparametersMap)
	if err != nil {
		return err
	}
	
	if Operator, ok := CampaignruleparametersMap["operator"].(string); ok {
		o.Operator = &Operator
	}
	
	if Value, ok := CampaignruleparametersMap["value"].(string); ok {
		o.Value = &Value
	}
	
	if Priority, ok := CampaignruleparametersMap["priority"].(string); ok {
		o.Priority = &Priority
	}
	
	if DialingMode, ok := CampaignruleparametersMap["dialingMode"].(string); ok {
		o.DialingMode = &DialingMode
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignruleparameters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
