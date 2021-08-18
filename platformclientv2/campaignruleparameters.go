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

func (u *Campaignruleparameters) MarshalJSON() ([]byte, error) {
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
		Operator: u.Operator,
		
		Value: u.Value,
		
		Priority: u.Priority,
		
		DialingMode: u.DialingMode,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Campaignruleparameters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
