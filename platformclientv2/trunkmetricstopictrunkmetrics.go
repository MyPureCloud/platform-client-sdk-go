package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricstopictrunkmetrics
type Trunkmetricstopictrunkmetrics struct { 
	// Calls
	Calls *Trunkmetricstopictrunkmetricscalls `json:"calls,omitempty"`


	// EventTime
	EventTime *time.Time `json:"eventTime,omitempty"`


	// Qos
	Qos *Trunkmetricstopictrunkmetricsqos `json:"qos,omitempty"`


	// Trunk
	Trunk *Trunkmetricstopicurireference `json:"trunk,omitempty"`

}

func (o *Trunkmetricstopictrunkmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricstopictrunkmetrics
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	return json.Marshal(&struct { 
		Calls *Trunkmetricstopictrunkmetricscalls `json:"calls,omitempty"`
		
		EventTime *string `json:"eventTime,omitempty"`
		
		Qos *Trunkmetricstopictrunkmetricsqos `json:"qos,omitempty"`
		
		Trunk *Trunkmetricstopicurireference `json:"trunk,omitempty"`
		*Alias
	}{ 
		Calls: o.Calls,
		
		EventTime: EventTime,
		
		Qos: o.Qos,
		
		Trunk: o.Trunk,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkmetricstopictrunkmetrics) UnmarshalJSON(b []byte) error {
	var TrunkmetricstopictrunkmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkmetricstopictrunkmetricsMap)
	if err != nil {
		return err
	}
	
	if Calls, ok := TrunkmetricstopictrunkmetricsMap["calls"].(map[string]interface{}); ok {
		CallsString, _ := json.Marshal(Calls)
		json.Unmarshal(CallsString, &o.Calls)
	}
	
	if eventTimeString, ok := TrunkmetricstopictrunkmetricsMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if Qos, ok := TrunkmetricstopictrunkmetricsMap["qos"].(map[string]interface{}); ok {
		QosString, _ := json.Marshal(Qos)
		json.Unmarshal(QosString, &o.Qos)
	}
	
	if Trunk, ok := TrunkmetricstopictrunkmetricsMap["trunk"].(map[string]interface{}); ok {
		TrunkString, _ := json.Marshal(Trunk)
		json.Unmarshal(TrunkString, &o.Trunk)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkmetricstopictrunkmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
