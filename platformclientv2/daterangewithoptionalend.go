package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Daterangewithoptionalend
type Daterangewithoptionalend struct { 
	// StartBusinessUnitDate - The start date for work plan rotation or an agent, interpreted in the business unit's time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartBusinessUnitDate *time.Time `json:"startBusinessUnitDate,omitempty"`


	// EndBusinessUnitDate - The end date for work plan rotation or an agent, interpreted in the business unit's time zone. Null denotes open ended date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	EndBusinessUnitDate *time.Time `json:"endBusinessUnitDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Daterangewithoptionalend) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
