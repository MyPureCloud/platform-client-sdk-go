package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekschedulegenerationresult
type Weekschedulegenerationresult struct { 
	// Failed - Whether the schedule generation failed
	Failed *bool `json:"failed,omitempty"`


	// RunId - ID of the schedule run
	RunId *string `json:"runId,omitempty"`


	// AgentWarnings - Warning messages from the schedule run. This will be available only when requesting information for a single week schedule
	AgentWarnings *[]Schedulegenerationwarning `json:"agentWarnings,omitempty"`


	// AgentWarningCount - Count of warning messages from the schedule run. This will be available only when requesting multiple week schedules
	AgentWarningCount *int `json:"agentWarningCount,omitempty"`

}

func (o *Weekschedulegenerationresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Weekschedulegenerationresult
	
	return json.Marshal(&struct { 
		Failed *bool `json:"failed,omitempty"`
		
		RunId *string `json:"runId,omitempty"`
		
		AgentWarnings *[]Schedulegenerationwarning `json:"agentWarnings,omitempty"`
		
		AgentWarningCount *int `json:"agentWarningCount,omitempty"`
		*Alias
	}{ 
		Failed: o.Failed,
		
		RunId: o.RunId,
		
		AgentWarnings: o.AgentWarnings,
		
		AgentWarningCount: o.AgentWarningCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Weekschedulegenerationresult) UnmarshalJSON(b []byte) error {
	var WeekschedulegenerationresultMap map[string]interface{}
	err := json.Unmarshal(b, &WeekschedulegenerationresultMap)
	if err != nil {
		return err
	}
	
	if Failed, ok := WeekschedulegenerationresultMap["failed"].(bool); ok {
		o.Failed = &Failed
	}
	
	if RunId, ok := WeekschedulegenerationresultMap["runId"].(string); ok {
		o.RunId = &RunId
	}
	
	if AgentWarnings, ok := WeekschedulegenerationresultMap["agentWarnings"].([]interface{}); ok {
		AgentWarningsString, _ := json.Marshal(AgentWarnings)
		json.Unmarshal(AgentWarningsString, &o.AgentWarnings)
	}
	
	if AgentWarningCount, ok := WeekschedulegenerationresultMap["agentWarningCount"].(float64); ok {
		AgentWarningCountInt := int(AgentWarningCount)
		o.AgentWarningCount = &AgentWarningCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Weekschedulegenerationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
