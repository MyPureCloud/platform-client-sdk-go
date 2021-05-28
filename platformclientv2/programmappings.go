package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Programmappings
type Programmappings struct { 
	// Program
	Program *Baseprogramentity `json:"program,omitempty"`


	// Queues
	Queues *[]Addressableentityref `json:"queues,omitempty"`


	// Flows
	Flows *[]Addressableentityref `json:"flows,omitempty"`


	// ModifiedBy
	ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

}

// String returns a JSON representation of the model
func (o *Programmappings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
