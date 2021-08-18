package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchpredictorrequest
type Patchpredictorrequest struct { 
	// RoutingTimeoutSeconds - Number of seconds allocated to predictive routing before attempting a different routing method. This is a value between 12 and 900 seconds.
	RoutingTimeoutSeconds *int `json:"routingTimeoutSeconds,omitempty"`


	// Schedule - The predictor schedule that determines when the predictor is used for routing interactions.
	Schedule *Predictorschedule `json:"schedule,omitempty"`


	// WorkloadBalancingConfig - The predictor balancing configuration to enable workload balancing
	WorkloadBalancingConfig *Predictorworkloadbalancing `json:"workloadBalancingConfig,omitempty"`

}

func (u *Patchpredictorrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchpredictorrequest

	

	return json.Marshal(&struct { 
		RoutingTimeoutSeconds *int `json:"routingTimeoutSeconds,omitempty"`
		
		Schedule *Predictorschedule `json:"schedule,omitempty"`
		
		WorkloadBalancingConfig *Predictorworkloadbalancing `json:"workloadBalancingConfig,omitempty"`
		*Alias
	}{ 
		RoutingTimeoutSeconds: u.RoutingTimeoutSeconds,
		
		Schedule: u.Schedule,
		
		WorkloadBalancingConfig: u.WorkloadBalancingConfig,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Patchpredictorrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
