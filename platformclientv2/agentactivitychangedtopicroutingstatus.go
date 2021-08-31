package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentactivitychangedtopicroutingstatus
type Agentactivitychangedtopicroutingstatus struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`

}

func (o *Agentactivitychangedtopicroutingstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentactivitychangedtopicroutingstatus
	
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

func (o *Agentactivitychangedtopicroutingstatus) UnmarshalJSON(b []byte) error {
	var AgentactivitychangedtopicroutingstatusMap map[string]interface{}
	err := json.Unmarshal(b, &AgentactivitychangedtopicroutingstatusMap)
	if err != nil {
		return err
	}
	
	if Status, ok := AgentactivitychangedtopicroutingstatusMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if startTimeString, ok := AgentactivitychangedtopicroutingstatusMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentactivitychangedtopicroutingstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
