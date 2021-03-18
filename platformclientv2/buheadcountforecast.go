package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Buheadcountforecast
type Buheadcountforecast struct { 
	// Entities
	Entities *[]Buplanninggroupheadcountforecast `json:"entities,omitempty"`


	// ReferenceStartDate - Reference start date for the interval values in each forecast entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buheadcountforecast) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
