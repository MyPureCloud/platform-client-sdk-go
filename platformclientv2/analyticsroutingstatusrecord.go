package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsroutingstatusrecord
type Analyticsroutingstatusrecord struct { 
	// StartTime - The start time of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - The end time of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// RoutingStatus - The user's ACD routing status
	RoutingStatus *string `json:"routingStatus,omitempty"`

}

func (o *Analyticsroutingstatusrecord) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsroutingstatusrecord
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	return json.Marshal(&struct { 
		StartTime *string `json:"startTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		RoutingStatus *string `json:"routingStatus,omitempty"`
		*Alias
	}{ 
		StartTime: StartTime,
		
		EndTime: EndTime,
		
		RoutingStatus: o.RoutingStatus,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsroutingstatusrecord) UnmarshalJSON(b []byte) error {
	var AnalyticsroutingstatusrecordMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsroutingstatusrecordMap)
	if err != nil {
		return err
	}
	
	if startTimeString, ok := AnalyticsroutingstatusrecordMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if endTimeString, ok := AnalyticsroutingstatusrecordMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if RoutingStatus, ok := AnalyticsroutingstatusrecordMap["routingStatus"].(string); ok {
		o.RoutingStatus = &RoutingStatus
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsroutingstatusrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
