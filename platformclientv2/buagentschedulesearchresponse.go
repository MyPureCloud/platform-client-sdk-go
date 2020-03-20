package platformclientv2
import (
	"encoding/json"
)

// Buagentschedulesearchresponse
type Buagentschedulesearchresponse struct { 
	// User - The user to whom this agent schedule applies
	User *Userreference `json:"user,omitempty"`


	// Shifts - The shift definitions for this agent schedule
	Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`


	// FullDayTimeOffMarkers - Full day time off markers which apply to this agent schedule
	FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulesearchresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
