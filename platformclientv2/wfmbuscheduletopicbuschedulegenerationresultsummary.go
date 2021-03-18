package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicbuschedulegenerationresultsummary
type Wfmbuscheduletopicbuschedulegenerationresultsummary struct { 
	// Failed
	Failed *bool `json:"failed,omitempty"`


	// RunId
	RunId *string `json:"runId,omitempty"`


	// MessageCount
	MessageCount *int `json:"messageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbuschedulegenerationresultsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
