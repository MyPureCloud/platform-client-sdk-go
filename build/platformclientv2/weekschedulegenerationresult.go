package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Weekschedulegenerationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
