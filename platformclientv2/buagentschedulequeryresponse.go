package platformclientv2
import (
	"encoding/json"
)

// Buagentschedulequeryresponse
type Buagentschedulequeryresponse struct { 
	// User - The user to whom this agent schedule applies
	User *Userreference `json:"user,omitempty"`


	// Shifts - The shift definitions for this agent schedule
	Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`


	// FullDayTimeOffMarkers - Full day time off markers which apply to this agent schedule
	FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`


	// WorkPlan - The work plan for this user
	WorkPlan *Workplanreference `json:"workPlan,omitempty"`


	// WorkPlansPerWeek - The work plans per week for this user from the work plan rotation. Null values in the list denotes that user is not part of any work plan for that week
	WorkPlansPerWeek *[]Workplanreference `json:"workPlansPerWeek,omitempty"`


	// Metadata - Versioned entity metadata for this agent schedule
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulequeryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
