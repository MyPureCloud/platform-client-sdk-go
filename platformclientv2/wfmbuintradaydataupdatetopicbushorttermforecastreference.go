package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuintradaydataupdatetopicbushorttermforecastreference
type Wfmbuintradaydataupdatetopicbushorttermforecastreference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// WeekDate
	WeekDate *string `json:"weekDate,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbushorttermforecastreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
