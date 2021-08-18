package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradayplanninggrouplisting - A list of IntradayPlanningGroup objects
type Wfmintradayplanninggrouplisting struct { 
	// Entities
	Entities *[]Forecastplanninggroupresponse `json:"entities,omitempty"`


	// NoDataReason - The reason there was no data for the request
	NoDataReason *string `json:"noDataReason,omitempty"`

}

func (u *Wfmintradayplanninggrouplisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradayplanninggrouplisting

	

	return json.Marshal(&struct { 
		Entities *[]Forecastplanninggroupresponse `json:"entities,omitempty"`
		
		NoDataReason *string `json:"noDataReason,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		
		NoDataReason: u.NoDataReason,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmintradayplanninggrouplisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
