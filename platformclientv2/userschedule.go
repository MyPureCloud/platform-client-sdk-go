package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Userschedule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userschedule

	

	return json.Marshal(&struct { 
		Shifts *[]Userscheduleshift `json:"shifts,omitempty"`
		
		FullDayTimeOffMarkers *[]Userschedulefulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		
		Delete *bool `json:"delete,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		WorkPlanId *string `json:"workPlanId,omitempty"`
		*Alias
	}{ 
		Shifts: u.Shifts,
		
		FullDayTimeOffMarkers: u.FullDayTimeOffMarkers,
		
		Delete: u.Delete,
		
		Metadata: u.Metadata,
		
		WorkPlanId: u.WorkPlanId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
