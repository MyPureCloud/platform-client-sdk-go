package platformclientv2
import (
	"time"
	"encoding/json"
)

// Workdayvaluestrenditem
type Workdayvaluestrenditem struct { 
	// DateWorkday - The workday for the metric value. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`


	// Value - The metric value
	Value *float64 `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Workdayvaluestrenditem) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
