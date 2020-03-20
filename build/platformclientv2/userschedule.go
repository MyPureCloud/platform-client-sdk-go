package platformclientv2
import (
	"encoding/json"
)

// Userschedule - A schedule for a single user over a given time range
type Userschedule struct { 
	// Shifts - The shifts that belong to this schedule
	Shifts *[]Userscheduleshift `json:"shifts,omitempty"`


	// FullDayTimeOffMarkers - Markers to indicate a full day time off request, relative to the management unit time zone
	FullDayTimeOffMarkers *[]Userschedulefulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`


	// Delete - If marked true for updating an existing user schedule, it will be deleted
	Delete *bool `json:"delete,omitempty"`


	// Metadata - Version metadata for this schedule
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// WorkPlanId - ID of the work plan associated with the user during schedule creation
	WorkPlanId *string `json:"workPlanId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userschedule) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
