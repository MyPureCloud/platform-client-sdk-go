package platformclientv2
import (
	"time"
	"encoding/json"
)

// Busearchagentschedulesrequest
type Busearchagentschedulesrequest struct { 
	// StartDate - Start date of the range to search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End date of the range to search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	EndDate *time.Time `json:"endDate,omitempty"`


	// UserIds - IDs of the users for whose schedules to search
	UserIds *[]string `json:"userIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Busearchagentschedulesrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
