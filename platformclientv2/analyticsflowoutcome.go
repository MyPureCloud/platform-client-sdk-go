package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsflowoutcome
type Analyticsflowoutcome struct { 
	// FlowOutcomeId - Unique identifiers of a flow outcome
	FlowOutcomeId *string `json:"flowOutcomeId,omitempty"`


	// FlowOutcomeValue - Flow outcome value, e.g. SUCCESS
	FlowOutcomeValue *string `json:"flowOutcomeValue,omitempty"`


	// FlowOutcome - Colon-separated combinations of unique flow outcome identifier and value
	FlowOutcome *string `json:"flowOutcome,omitempty"`


	// FlowOutcomeStartTimestamp - Date/time the outcome started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	FlowOutcomeStartTimestamp *time.Time `json:"flowOutcomeStartTimestamp,omitempty"`


	// FlowOutcomeEndTimestamp - Date/time the outcome ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	FlowOutcomeEndTimestamp *time.Time `json:"flowOutcomeEndTimestamp,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsflowoutcome) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
