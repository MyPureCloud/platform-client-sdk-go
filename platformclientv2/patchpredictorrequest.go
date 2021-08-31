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

func (o *Patchpredictorrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchpredictorrequest
	
	return json.Marshal(&struct { 
		RoutingTimeoutSeconds *int `json:"routingTimeoutSeconds,omitempty"`
		
		Schedule *Predictorschedule `json:"schedule,omitempty"`
		
		WorkloadBalancingConfig *Predictorworkloadbalancing `json:"workloadBalancingConfig,omitempty"`
		*Alias
	}{ 
		RoutingTimeoutSeconds: o.RoutingTimeoutSeconds,
		
		Schedule: o.Schedule,
		
		WorkloadBalancingConfig: o.WorkloadBalancingConfig,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchpredictorrequest) UnmarshalJSON(b []byte) error {
	var PatchpredictorrequestMap map[string]interface{}
	err := json.Unmarshal(b, &PatchpredictorrequestMap)
	if err != nil {
		return err
	}
	
	if RoutingTimeoutSeconds, ok := PatchpredictorrequestMap["routingTimeoutSeconds"].(float64); ok {
		RoutingTimeoutSecondsInt := int(RoutingTimeoutSeconds)
		o.RoutingTimeoutSeconds = &RoutingTimeoutSecondsInt
	}
	
	if Schedule, ok := PatchpredictorrequestMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if WorkloadBalancingConfig, ok := PatchpredictorrequestMap["workloadBalancingConfig"].(map[string]interface{}); ok {
		WorkloadBalancingConfigString, _ := json.Marshal(WorkloadBalancingConfig)
		json.Unmarshal(WorkloadBalancingConfigString, &o.WorkloadBalancingConfig)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchpredictorrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
