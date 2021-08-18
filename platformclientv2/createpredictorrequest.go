package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Createpredictorrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createpredictorrequest

	

	return json.Marshal(&struct { 
		QueueIds *[]string `json:"queueIds,omitempty"`
		
		Kpi *string `json:"kpi,omitempty"`
		
		RoutingTimeoutSeconds *int `json:"routingTimeoutSeconds,omitempty"`
		
		Schedule *Predictorschedule `json:"schedule,omitempty"`
		
		WorkloadBalancingConfig *Predictorworkloadbalancing `json:"workloadBalancingConfig,omitempty"`
		*Alias
	}{ 
		QueueIds: u.QueueIds,
		
		Kpi: u.Kpi,
		
		RoutingTimeoutSeconds: u.RoutingTimeoutSeconds,
		
		Schedule: u.Schedule,
		
		WorkloadBalancingConfig: u.WorkloadBalancingConfig,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createpredictorrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
