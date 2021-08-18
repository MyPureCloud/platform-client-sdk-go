package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Groupgreetingeventgreetingowner
type Groupgreetingeventgreetingowner struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (u *Groupgreetingeventgreetingowner) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Groupgreetingeventgreetingowner

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Groupgreetingeventgreetingowner) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
