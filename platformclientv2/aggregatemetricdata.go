package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregatemetricdata
type Aggregatemetricdata struct { 
	// Metric
	Metric *string `json:"metric,omitempty"`


	// Qualifier
	Qualifier *string `json:"qualifier,omitempty"`


	// Stats
	Stats *Statisticalsummary `json:"stats,omitempty"`

}

func (u *Aggregatemetricdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Aggregatemetricdata

	

	return json.Marshal(&struct { 
		Metric *string `json:"metric,omitempty"`
		
		Qualifier *string `json:"qualifier,omitempty"`
		
		Stats *Statisticalsummary `json:"stats,omitempty"`
		*Alias
	}{ 
		Metric: u.Metric,
		
		Qualifier: u.Qualifier,
		
		Stats: u.Stats,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Aggregatemetricdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
