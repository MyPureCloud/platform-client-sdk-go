package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Busearchagentschedulesrequest
type Busearchagentschedulesrequest struct { 
	// StartDate - Start date of the range to search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End date of the range to search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`


	// UserIds - IDs of the users for whose schedules to search
	UserIds *[]string `json:"userIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Busearchagentschedulesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
