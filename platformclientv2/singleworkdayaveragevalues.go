package platformclientv2
import (
	"time"
	"encoding/json"
)

// Singleworkdayaveragevalues
type Singleworkdayaveragevalues struct { 
	// DateWorkday - The targeted workday for average value query. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`


	// Division - The targeted division for the metrics
	Division *Division `json:"division,omitempty"`


	// User - The targeted user for the metrics
	User *Userreference `json:"user,omitempty"`


	// Timezone - The time zone used for aggregating metric values
	Timezone *string `json:"timezone,omitempty"`


	// Results - The metric value averages
	Results *[]Workdayvaluesmetricitem `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Singleworkdayaveragevalues) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
