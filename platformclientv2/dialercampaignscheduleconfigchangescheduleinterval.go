package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignscheduleconfigchangescheduleinterval
type Dialercampaignscheduleconfigchangescheduleinterval struct { 
	// Start
	Start *string `json:"start,omitempty"`


	// End
	End *string `json:"end,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialercampaignscheduleconfigchangescheduleinterval) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignscheduleconfigchangescheduleinterval

	

	return json.Marshal(&struct { 
		Start *string `json:"start,omitempty"`
		
		End *string `json:"end,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Start: u.Start,
		
		End: u.End,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercampaignscheduleconfigchangescheduleinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
