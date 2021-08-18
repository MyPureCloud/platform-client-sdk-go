package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastplanninggroupsresponse
type Forecastplanninggroupsresponse struct { 
	// Entities
	Entities *[]Forecastplanninggroupresponse `json:"entities,omitempty"`

}

func (u *Forecastplanninggroupsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Forecastplanninggroupsresponse

	

	return json.Marshal(&struct { 
		Entities *[]Forecastplanninggroupresponse `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Forecastplanninggroupsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
