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

func (o *Trunkmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetrics
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		LogicalInterface: o.LogicalInterface,
		
		Trunk: o.Trunk,
		
		Calls: o.Calls,
		
		Qos: o.Qos,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkmetrics) UnmarshalJSON(b []byte) error {
	var TrunkmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkmetricsMap)
	if err != nil {
		return err
	}
	
	if eventTimeString, ok := TrunkmetricsMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if LogicalInterface, ok := TrunkmetricsMap["logicalInterface"].(map[string]interface{}); ok {
		LogicalInterfaceString, _ := json.Marshal(LogicalInterface)
		json.Unmarshal(LogicalInterfaceString, &o.LogicalInterface)
	}
	
	if Trunk, ok := TrunkmetricsMap["trunk"].(map[string]interface{}); ok {
		TrunkString, _ := json.Marshal(Trunk)
		json.Unmarshal(TrunkString, &o.Trunk)
	}
	
	if Calls, ok := TrunkmetricsMap["calls"].(map[string]interface{}); ok {
		CallsString, _ := json.Marshal(Calls)
		json.Unmarshal(CallsString, &o.Calls)
	}
	
	if Qos, ok := TrunkmetricsMap["qos"].(map[string]interface{}); ok {
		QosString, _ := json.Marshal(Qos)
		json.Unmarshal(QosString, &o.Qos)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
