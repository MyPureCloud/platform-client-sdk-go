package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Copybuforecastrequest
type Copybuforecastrequest struct { 
	// Description - The description for the forecast
	Description *string `json:"description,omitempty"`


	// WeekDate - The start date of the new forecast to create from the existing forecast. Must correspond to the start day of week for the business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Copybuforecastrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
