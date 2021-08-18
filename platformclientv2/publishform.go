package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Publishform
type Publishform struct { 
	// Published - Is this form published
	Published *bool `json:"published,omitempty"`


	// Id - Unique Id for this version of this form
	Id *string `json:"id,omitempty"`

}

func (u *Publishform) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Publishform

	

	return json.Marshal(&struct { 
		Published *bool `json:"published,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Published: u.Published,
		
		Id: u.Id,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Publishform) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
