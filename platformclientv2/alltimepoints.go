package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Alltimepoints
type Alltimepoints struct { 
	// User - Queried user
	User *Userreference `json:"user,omitempty"`


	// DateEndWorkday - Queried end workday for all time points to be collected. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`


	// AllTimePoints - All time point collected bt the user
	AllTimePoints *int `json:"allTimePoints,omitempty"`

}

// String returns a JSON representation of the model
func (o *Alltimepoints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
