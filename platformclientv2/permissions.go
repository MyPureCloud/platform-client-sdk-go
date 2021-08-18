package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Permissions
type Permissions struct { 
	// Ids - List of permission ids.
	Ids *[]string `json:"ids,omitempty"`

}

func (u *Permissions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Permissions

	

	return json.Marshal(&struct { 
		Ids *[]string `json:"ids,omitempty"`
		*Alias
	}{ 
		Ids: u.Ids,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Permissions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
