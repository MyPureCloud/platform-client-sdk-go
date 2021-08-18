package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetrics
type Trunkmetrics struct { 
	// EventTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventTime *time.Time `json:"eventTime,omitempty"`


	// LogicalInterface
	LogicalInterface *Domainentityref `json:"logicalInterface,omitempty"`


	// Trunk
	Trunk *Domainentityref `json:"trunk,omitempty"`


	// Calls
	Calls *Trunkmetricscalls `json:"calls,omitempty"`


	// Qos
	Qos *Trunkmetricsqos `json:"qos,omitempty"`

}

func (u *Trunkmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetrics

	
	EventTime := new(string)
	if u.EventTime != nil {
		
		*EventTime = timeutil.Strftime(u.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	

	return json.Marshal(&struct { 
		EventTime *string `json:"eventTime,omitempty"`
		
		LogicalInterface *Domainentityref `json:"logicalInterface,omitempty"`
		
		Trunk *Domainentityref `json:"trunk,omitempty"`
		
		Calls *Trunkmetricscalls `json:"calls,omitempty"`
		
		Qos *Trunkmetricsqos `json:"qos,omitempty"`
		*Alias
	}{ 
		EventTime: EventTime,
		
		LogicalInterface: u.LogicalInterface,
		
		Trunk: u.Trunk,
		
		Calls: u.Calls,
		
		Qos: u.Qos,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
