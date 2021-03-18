package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsroutingstatusrecord
type Analyticsroutingstatusrecord struct { 
	// StartTime - The start time of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - The end time of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// RoutingStatus - The user's ACD routing status
	RoutingStatus *string `json:"routingStatus,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsroutingstatusrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
