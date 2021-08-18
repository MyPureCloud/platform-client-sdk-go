package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Programmappings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Programmappings

	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Program *Baseprogramentity `json:"program,omitempty"`
		
		Queues *[]Addressableentityref `json:"queues,omitempty"`
		
		Flows *[]Addressableentityref `json:"flows,omitempty"`
		
		ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		*Alias
	}{ 
		Program: u.Program,
		
		Queues: u.Queues,
		
		Flows: u.Flows,
		
		ModifiedBy: u.ModifiedBy,
		
		DateModified: DateModified,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Programmappings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
