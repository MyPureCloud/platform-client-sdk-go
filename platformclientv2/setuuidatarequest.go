package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Setuuidatarequest
type Setuuidatarequest struct { 
	// UuiData - The value of the uuiData to set.
	UuiData *string `json:"uuiData,omitempty"`

}

func (u *Setuuidatarequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Setuuidatarequest

	

	return json.Marshal(&struct { 
		UuiData *string `json:"uuiData,omitempty"`
		*Alias
	}{ 
		UuiData: u.UuiData,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Setuuidatarequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
