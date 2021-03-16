package platformclientv2
import (
	"time"
	"encoding/json"
)

// Userbestpointsitem
type Userbestpointsitem struct { 
	// GranularityType - Best points aggregation interval granularity
	GranularityType *string `json:"granularityType,omitempty"`


	// Points - Gamification points
	Points *int `json:"points,omitempty"`


	// DateStartWorkday - Start workday of the best points aggregation interval. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`


	// DateEndWorkday - End workday of the best points aggregation interval. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`


	// Rank - The rank of this user
	Rank *int `json:"rank,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userbestpointsitem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
