package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Publishdraftinput - Draft to be published
type Publishdraftinput struct { 
	// Version - The current draft version.
	Version *int `json:"version,omitempty"`

}

func (u *Publishdraftinput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Publishdraftinput

	

	return json.Marshal(&struct { 
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		Version: u.Version,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Publishdraftinput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
