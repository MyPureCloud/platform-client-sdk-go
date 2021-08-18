package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregateviewdata
type Aggregateviewdata struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Stats
	Stats *Statisticalsummary `json:"stats,omitempty"`

}

func (u *Aggregateviewdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Aggregateviewdata

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Stats *Statisticalsummary `json:"stats,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Stats: u.Stats,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Aggregateviewdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
