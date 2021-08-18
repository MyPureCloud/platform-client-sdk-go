package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictor
type Predictor struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Queues - The queue IDs associated with the predictor.
	Queues *[]Addressableentityref `json:"queues,omitempty"`


	// Kpi - The KPI that the predictor attempts to maximize/minimize.
	Kpi *string `json:"kpi,omitempty"`


	// RoutingTimeoutSeconds - Number of seconds allocated to predictive routing before attempting a different routing method. This is a value between 12 and 900 seconds.
	RoutingTimeoutSeconds *int `json:"routingTimeoutSeconds,omitempty"`


	// Schedule - The predictor schedule that determines when the predictor is used for routing interactions.
	Schedule *Predictorschedule `json:"schedule,omitempty"`


	// State - The predictor state.
	State *string `json:"state,omitempty"`


	// DateCreated - DateTime indicating when the predictor was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - DateTime indicating when the predictor was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// WorkloadBalancingConfig - The predictor balancing configuration to enable workload balancing.
	WorkloadBalancingConfig *Predictorworkloadbalancing `json:"workloadBalancingConfig,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Predictor) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Predictor

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Queues *[]Addressableentityref `json:"queues,omitempty"`
		
		Kpi *string `json:"kpi,omitempty"`
		
		RoutingTimeoutSeconds *int `json:"routingTimeoutSeconds,omitempty"`
		
		Schedule *Predictorschedule `json:"schedule,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		WorkloadBalancingConfig *Predictorworkloadbalancing `json:"workloadBalancingConfig,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Queues: u.Queues,
		
		Kpi: u.Kpi,
		
		RoutingTimeoutSeconds: u.RoutingTimeoutSeconds,
		
		Schedule: u.Schedule,
		
		State: u.State,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		WorkloadBalancingConfig: u.WorkloadBalancingConfig,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Predictor) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
