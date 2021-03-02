package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmschedulereference
type Wfmschedulereference struct { 
	// Id - The ID of the WFM schedule
	Id *string `json:"id,omitempty"`


	// BusinessUnit - A reference to a Workforce Management Business Unit
	BusinessUnit *Wfmbusinessunitreference `json:"businessUnit,omitempty"`


	// WeekDate - The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	WeekDate *time.Time `json:"weekDate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmschedulereference) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
