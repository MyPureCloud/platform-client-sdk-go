package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Roledivisiongrants
type Roledivisiongrants struct { 
	// Grants - A list containing pairs of role and division IDs
	Grants *[]Roledivisionpair `json:"grants,omitempty"`

}

func (u *Roledivisiongrants) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Roledivisiongrants

	

	return json.Marshal(&struct { 
		Grants *[]Roledivisionpair `json:"grants,omitempty"`
		*Alias
	}{ 
		Grants: u.Grants,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Roledivisiongrants) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
