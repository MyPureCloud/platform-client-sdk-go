package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatguestmediarequestentitylist
type Webchatguestmediarequestentitylist struct { 
	// Entities
	Entities *[]Webchatguestmediarequest `json:"entities,omitempty"`

}

func (u *Webchatguestmediarequestentitylist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatguestmediarequestentitylist

	

	return json.Marshal(&struct { 
		Entities *[]Webchatguestmediarequest `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webchatguestmediarequestentitylist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
