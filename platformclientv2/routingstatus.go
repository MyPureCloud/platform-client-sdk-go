package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingstatus
type Routingstatus struct { 
	// UserId - The userId of the agent
	UserId *string `json:"userId,omitempty"`


	// Status - Indicates the Routing State of the agent.  A value of OFF_QUEUE will be returned if the specified user does not exist.
	Status *string `json:"status,omitempty"`


	// StartTime - The timestamp when the agent went into this state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`

}

func (o *Routingstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingstatus
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		*Alias
	}{ 
		UserId: o.UserId,
		
		Status: o.Status,
		
		StartTime: StartTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Routingstatus) UnmarshalJSON(b []byte) error {
	var RoutingstatusMap map[string]interface{}
	err := json.Unmarshal(b, &RoutingstatusMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := RoutingstatusMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if Status, ok := RoutingstatusMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if startTimeString, ok := RoutingstatusMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routingstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
