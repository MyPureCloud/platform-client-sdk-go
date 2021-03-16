package platformclientv2
import (
	"time"
	"encoding/json"
)

// Overallbestpointsitem
type Overallbestpointsitem struct { 
	// GranularityType - Best points aggregation interval granularity
	GranularityType *string `json:"granularityType,omitempty"`


	// Users - List of associated users with the equal points.
	Users *[]Userreference `json:"users,omitempty"`


	// Count - The count of the user IDs in the list
	Count *int `json:"count,omitempty"`


	// Points - Gamification points
	Points *int `json:"points,omitempty"`


	// DateStartWorkday - Start workday of the best points aggregation interval. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`


	// DateEndWorkday - End workday of the best points aggregation interval. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`

}

// String returns a JSON representation of the model
func (o *Overallbestpointsitem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
