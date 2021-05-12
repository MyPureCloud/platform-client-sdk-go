package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Daterange
type Daterange struct { 
	// StartDate - The inclusive start of a date range in yyyy-MM-dd format. Should be interpreted in the management unit's configured time zone.
	StartDate *string `json:"startDate,omitempty"`


	// EndDate - The inclusive end of a date range in yyyy-MM-dd format. Should be interpreted in the management unit's configured time zone.
	EndDate *string `json:"endDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Daterange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
