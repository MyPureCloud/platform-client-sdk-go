package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialercampaignruleconfigchangecampaignrulecondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignruleconfigchangecampaignrulecondition

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Parameters *map[string]string `json:"parameters,omitempty"`
		
		ConditionType *string `json:"conditionType,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Parameters: u.Parameters,
		
		ConditionType: u.ConditionType,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercampaignruleconfigchangecampaignrulecondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
