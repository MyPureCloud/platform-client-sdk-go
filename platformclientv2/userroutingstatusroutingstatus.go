package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingstatusroutingstatus
type Userroutingstatusroutingstatus struct { 
	// Status - Indicates the Routing State of the agent.
	Status *string `json:"status,omitempty"`


	// StartTime - The timestamp when the agent went into this state.
	StartTime *time.Time `json:"startTime,omitempty"`

}

func (o *Userroutingstatusroutingstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userroutingstatusroutingstatus
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		StartTime: StartTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Userroutingstatusroutingstatus) UnmarshalJSON(b []byte) error {
	var UserroutingstatusroutingstatusMap map[string]interface{}
	err := json.Unmarshal(b, &UserroutingstatusroutingstatusMap)
	if err != nil {
		return err
	}
	
	if Status, ok := UserroutingstatusroutingstatusMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if startTimeString, ok := UserroutingstatusroutingstatusMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userroutingstatusroutingstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
