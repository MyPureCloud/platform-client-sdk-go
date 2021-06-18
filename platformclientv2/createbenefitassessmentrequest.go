package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Createbenefitassessmentrequest
type Createbenefitassessmentrequest struct { 
	// QueueIds - The list of queue ids that are to be assessed for Predictive Routing benefit.
	QueueIds *[]string `json:"queueIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createbenefitassessmentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
