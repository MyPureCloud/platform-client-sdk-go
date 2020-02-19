package platformclientv2
import (
	"encoding/json"
)

// Userschedulecontainer - Container object to hold a map of user schedules
type Userschedulecontainer struct { 
	// ManagementUnitTimeZone - The reference time zone used for the management unit
	ManagementUnitTimeZone *string `json:"managementUnitTimeZone,omitempty"`


	// PublishedSchedules - References to all published week schedules overlapping the start/end date query parameters
	PublishedSchedules *[]Weekschedulereference `json:"publishedSchedules,omitempty"`


	// UserSchedules - Map of user id to user schedule
	UserSchedules *map[string]Userschedule `json:"userSchedules,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userschedulecontainer) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
