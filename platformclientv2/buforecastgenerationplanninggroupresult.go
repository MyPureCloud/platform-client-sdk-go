package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buforecastgenerationplanninggroupresult
type Buforecastgenerationplanninggroupresult struct { 
	// PlanningGroupId - The ID of the planning group
	PlanningGroupId *string `json:"planningGroupId,omitempty"`


	// MetricResults - The generation results for the associated planning group
	MetricResults *[]Buforecasttimeseriesresult `json:"metricResults,omitempty"`

}

func (u *Buforecastgenerationplanninggroupresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buforecastgenerationplanninggroupresult

	

	return json.Marshal(&struct { 
		PlanningGroupId *string `json:"planningGroupId,omitempty"`
		
		MetricResults *[]Buforecasttimeseriesresult `json:"metricResults,omitempty"`
		*Alias
	}{ 
		PlanningGroupId: u.PlanningGroupId,
		
		MetricResults: u.MetricResults,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buforecastgenerationplanninggroupresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
