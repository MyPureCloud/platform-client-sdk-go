package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventqueuetopicdatum
type Stateventqueuetopicdatum struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventqueuetopicmetric `json:"metrics,omitempty"`

}

func (u *Stateventqueuetopicdatum) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventqueuetopicdatum

	

	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Stateventqueuetopicmetric `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: u.Interval,
		
		Metrics: u.Metrics,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Stateventqueuetopicdatum) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
