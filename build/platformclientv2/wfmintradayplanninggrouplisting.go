package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Wfmintradayplanninggrouplisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
