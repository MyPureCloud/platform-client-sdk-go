package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingrule
type Routingrule struct { 
	// Operator - matching operator.  MEETS_THRESHOLD matches any agent with a score at or above the rule's threshold.  ANY matches all specified agents, regardless of score.
	Operator *string `json:"operator,omitempty"`


	// Threshold - threshold required for routing attempt (generally an agent score).  may be null for operator ANY.
	Threshold *int `json:"threshold,omitempty"`


	// WaitSeconds - seconds to wait in this rule before moving to the next
	WaitSeconds *float64 `json:"waitSeconds,omitempty"`

}

func (u *Routingrule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingrule

	

	return json.Marshal(&struct { 
		Operator *string `json:"operator,omitempty"`
		
		Threshold *int `json:"threshold,omitempty"`
		
		WaitSeconds *float64 `json:"waitSeconds,omitempty"`
		*Alias
	}{ 
		Operator: u.Operator,
		
		Threshold: u.Threshold,
		
		WaitSeconds: u.WaitSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Routingrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
