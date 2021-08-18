package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Deleteretention
type Deleteretention struct { 
	// Days
	Days *int `json:"days,omitempty"`

}

func (u *Deleteretention) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Deleteretention

	

	return json.Marshal(&struct { 
		Days *int `json:"days,omitempty"`
		*Alias
	}{ 
		Days: u.Days,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Deleteretention) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
