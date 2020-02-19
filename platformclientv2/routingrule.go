package platformclientv2
import (
	"encoding/json"
)

// Routingrule
type Routingrule struct { 
	// Operator - matching operator.  MEETS_THRESHOLD matches any agent with a score at or above the rule's threshold.  ANY matches all specified agents, regardless of score.
	Operator *string `json:"operator,omitempty"`


	// Threshold - threshold required for routing attempt (generally an agent score).  may be null for operator ANY.
	Threshold *int32 `json:"threshold,omitempty"`


	// WaitSeconds - seconds to wait in this rule before moving to the next
	WaitSeconds *float64 `json:"waitSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Routingrule) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
