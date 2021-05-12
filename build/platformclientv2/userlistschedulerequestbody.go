package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Userlistschedulerequestbody - Request body for fetching the schedule for a group of users over a given time range
type Userlistschedulerequestbody struct { 
	// UserIds - The user ids for which to fetch schedules
	UserIds *[]string `json:"userIds,omitempty"`


	// StartDate - Beginning of the range of schedules to fetch, in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End of the range of schedules to fetch, in ISO-8601 format
	EndDate *time.Time `json:"endDate,omitempty"`


	// LoadFullWeeks - Whether to load the full week's schedule (for the requested users) of any week overlapping the start/end date query parameters, defaults to false
	LoadFullWeeks *bool `json:"loadFullWeeks,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userlistschedulerequestbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
