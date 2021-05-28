package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Programmappingsrequest
type Programmappingsrequest struct { 
	// QueueIds - The program queues
	QueueIds *[]string `json:"queueIds,omitempty"`


	// FlowIds - The program flows
	FlowIds *[]string `json:"flowIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Programmappingsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
