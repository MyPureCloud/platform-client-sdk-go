package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queryresponsedata
type Queryresponsedata struct { 
	// Interval - Interval with start and end represented as ISO-8601 string. i.e: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'/yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	Interval *string `json:"interval,omitempty"`


	// Metrics - A list of aggregated metrics
	Metrics *[]Queryresponsemetric `json:"metrics,omitempty"`

}

func (u *Queryresponsedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queryresponsedata

	

	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Queryresponsemetric `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: u.Interval,
		
		Metrics: u.Metrics,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queryresponsedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
