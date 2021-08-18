package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Programmappingsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Programmappingsrequest

	

	return json.Marshal(&struct { 
		QueueIds *[]string `json:"queueIds,omitempty"`
		
		FlowIds *[]string `json:"flowIds,omitempty"`
		*Alias
	}{ 
		QueueIds: u.QueueIds,
		
		FlowIds: u.FlowIds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Programmappingsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
