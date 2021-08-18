package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Entry
type Entry struct { 
	// Value - A value included in this facet.
	Value *string `json:"value,omitempty"`


	// Count - The number of results with this value.
	Count *int `json:"count,omitempty"`

}

func (u *Entry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Entry

	

	return json.Marshal(&struct { 
		Value *string `json:"value,omitempty"`
		
		Count *int `json:"count,omitempty"`
		*Alias
	}{ 
		Value: u.Value,
		
		Count: u.Count,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Entry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
