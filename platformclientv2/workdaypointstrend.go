package platformclientv2
import (
	"time"
	"encoding/json"
)

// Workdaypointstrend
type Workdaypointstrend struct { 
	// DateStartWorkday - The start workday for the query range for the gamification points trend. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`


	// DateEndWorkday - The end workday for the query range for the gamification points trend. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`


	// User - The targeted user for the query
	User *Userreference `json:"user,omitempty"`


	// DayOfWeek - Aggregated for same day comparison
	DayOfWeek *string `json:"dayOfWeek,omitempty"`


	// AveragePoints - The total average points
	AveragePoints *float64 `json:"averagePoints,omitempty"`


	// Trend - Daily points trends
	Trend *[]Workdaypointstrenditem `json:"trend,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workdaypointstrend) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
