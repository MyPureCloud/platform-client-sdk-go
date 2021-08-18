package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentsortitem
type Contentsortitem struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Ascending
	Ascending *bool `json:"ascending,omitempty"`

}

func (u *Contentsortitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentsortitem

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Ascending *bool `json:"ascending,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Ascending: u.Ascending,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contentsortitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
