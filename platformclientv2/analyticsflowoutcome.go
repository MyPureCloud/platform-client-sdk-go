package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsflowoutcome
type Analyticsflowoutcome struct { 
	// FlowOutcome - Combination of unique flow outcome identifier and its value separated by colon
	FlowOutcome *string `json:"flowOutcome,omitempty"`


	// FlowOutcomeEndTimestamp - The outcome ending timestamp in ISO 8601 format. This may be null if the outcome did not succeed.
	FlowOutcomeEndTimestamp *time.Time `json:"flowOutcomeEndTimestamp,omitempty"`


	// FlowOutcomeId - Unique identifier of a flow outcome
	FlowOutcomeId *string `json:"flowOutcomeId,omitempty"`


	// FlowOutcomeStartTimestamp - The outcome starting timestamp in ISO 8601 format
	FlowOutcomeStartTimestamp *time.Time `json:"flowOutcomeStartTimestamp,omitempty"`


	// FlowOutcomeValue - Flow outcome value, e.g. SUCCESS
	FlowOutcomeValue *string `json:"flowOutcomeValue,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsflowoutcome) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
