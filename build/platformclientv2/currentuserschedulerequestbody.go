package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Currentuserschedulerequestbody - POST request body for fetching the current user's schedule over a given range
type Currentuserschedulerequestbody struct { 
	// StartDate - Beginning of the range of schedules to fetch, in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End of the range of schedules to fetch, in ISO-8601 format
	EndDate *time.Time `json:"endDate,omitempty"`


	// LoadFullWeeks - Whether to load the full week's schedule (for the current user) of any week overlapping the start/end date query parameters, defaults to false
	LoadFullWeeks *bool `json:"loadFullWeeks,omitempty"`

}

// String returns a JSON representation of the model
func (o *Currentuserschedulerequestbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
