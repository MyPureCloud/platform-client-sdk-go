package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventusertopicdatum
type Stateventusertopicdatum struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventusertopicmetric `json:"metrics,omitempty"`

}

func (u *Stateventusertopicdatum) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventusertopicdatum

	

	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Stateventusertopicmetric `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: u.Interval,
		
		Metrics: u.Metrics,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Stateventusertopicdatum) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
