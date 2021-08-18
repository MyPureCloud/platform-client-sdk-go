package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivityaggregatequeryresponsemetric
type Developmentactivityaggregatequeryresponsemetric struct { 
	// Metric - The metric this applies to
	Metric *string `json:"metric,omitempty"`


	// Stats - The aggregated values for this metric
	Stats *Developmentactivityaggregatequeryresponsestatistics `json:"stats,omitempty"`

}

func (u *Developmentactivityaggregatequeryresponsemetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Developmentactivityaggregatequeryresponsemetric

	

	return json.Marshal(&struct { 
		Metric *string `json:"metric,omitempty"`
		
		Stats *Developmentactivityaggregatequeryresponsestatistics `json:"stats,omitempty"`
		*Alias
	}{ 
		Metric: u.Metric,
		
		Stats: u.Stats,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsemetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
