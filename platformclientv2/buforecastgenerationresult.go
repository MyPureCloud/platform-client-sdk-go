package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buforecastgenerationresult
type Buforecastgenerationresult struct { 
	// PlanningGroupResults - Generation results, broken down by planning group
	PlanningGroupResults *[]Buforecastgenerationplanninggroupresult `json:"planningGroupResults,omitempty"`

}

func (u *Buforecastgenerationresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buforecastgenerationresult

	

	return json.Marshal(&struct { 
		PlanningGroupResults *[]Buforecastgenerationplanninggroupresult `json:"planningGroupResults,omitempty"`
		*Alias
	}{ 
		PlanningGroupResults: u.PlanningGroupResults,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buforecastgenerationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
