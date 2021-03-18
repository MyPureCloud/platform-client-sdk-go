package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Campaignruleparameters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
