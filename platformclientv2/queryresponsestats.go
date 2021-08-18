package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queryresponsestats
type Queryresponsestats struct { 
	// Count - The count for this metric
	Count *int `json:"count,omitempty"`

}

func (u *Queryresponsestats) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queryresponsestats

	

	return json.Marshal(&struct { 
		Count *int `json:"count,omitempty"`
		*Alias
	}{ 
		Count: u.Count,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queryresponsestats) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
