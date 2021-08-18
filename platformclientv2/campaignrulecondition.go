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

func (u *Campaignrulecondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Campaignrulecondition

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Parameters *Campaignruleparameters `json:"parameters,omitempty"`
		
		ConditionType *string `json:"conditionType,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Parameters: u.Parameters,
		
		ConditionType: u.ConditionType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Campaignrulecondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
