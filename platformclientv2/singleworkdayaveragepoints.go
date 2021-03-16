package platformclientv2
import (
	"time"
	"encoding/json"
)

// Singleworkdayaveragepoints
type Singleworkdayaveragepoints struct { 
	// DateWorkday - Queried target workday. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`


	// Division - The targeted division for the average points
	Division *Division `json:"division,omitempty"`


	// AveragePoints - The average points per agent earned within the division
	AveragePoints *float64 `json:"averagePoints,omitempty"`

}

// String returns a JSON representation of the model
func (o *Singleworkdayaveragepoints) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
