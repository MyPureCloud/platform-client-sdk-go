package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Createpredictorrequest
type Createpredictorrequest struct { 
	// QueueIds - The queue IDs associated with the predictor.
	QueueIds *[]string `json:"queueIds,omitempty"`


	// Kpi - The KPI that the predictor attempts to maximize/minimize.
	Kpi *string `json:"kpi,omitempty"`


	// RoutingTimeoutSeconds - Number of seconds allocated to predictive routing before attempting a different routing method. This is a value between 12 and 900 seconds.
	RoutingTimeoutSeconds *int `json:"routingTimeoutSeconds,omitempty"`


	// Schedule - The predictor schedule that determines when the predictor is used for routing interactions.
	Schedule *Predictorschedule `json:"schedule,omitempty"`


	// WorkloadBalancingConfig - The predictor balancing configuration to enable workload balancing
	WorkloadBalancingConfig *Predictorworkloadbalancing `json:"workloadBalancingConfig,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createpredictorrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
