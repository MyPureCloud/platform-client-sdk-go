package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsuserpresencerecord
type Analyticsuserpresencerecord struct { 
	// StartTime - The start time of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - The end time of the record. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// SystemPresence - The user's system presence
	SystemPresence *string `json:"systemPresence,omitempty"`


	// OrganizationPresenceId - The identifier for the user's organization presence
	OrganizationPresenceId *string `json:"organizationPresenceId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsuserpresencerecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
