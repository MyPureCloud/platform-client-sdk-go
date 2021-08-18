package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lineuserid - Channel-specific User ID for Line accounts
type Lineuserid struct { 
	// UserId - The unique channel-specific userId for the user
	UserId *string `json:"userId,omitempty"`

}

func (u *Lineuserid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lineuserid

	

	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		*Alias
	}{ 
		UserId: u.UserId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Lineuserid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
