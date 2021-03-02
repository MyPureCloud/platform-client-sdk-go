package platformclientv2
import (
	"encoding/json"
)

// Buagentschedulehistorychange
type Buagentschedulehistorychange struct { 
	// Metadata - The metadata of the change, including who and when the change was made
	Metadata *Buagentschedulehistorychangemetadata `json:"metadata,omitempty"`


	// Shifts - The list of changed shifts
	Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`


	// FullDayTimeOffMarkers - The list of changed full day time off markers
	FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`


	// Deletes - The deleted shifts, full day time off markers, or the entire agent schedule
	Deletes *Buagentschedulehistorydeletedchange `json:"deletes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorychange) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
