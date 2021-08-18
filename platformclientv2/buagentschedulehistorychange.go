package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Buagentschedulehistorychange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentschedulehistorychange

	

	return json.Marshal(&struct { 
		Metadata *Buagentschedulehistorychangemetadata `json:"metadata,omitempty"`
		
		Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`
		
		FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		
		Deletes *Buagentschedulehistorydeletedchange `json:"deletes,omitempty"`
		*Alias
	}{ 
		Metadata: u.Metadata,
		
		Shifts: u.Shifts,
		
		FullDayTimeOffMarkers: u.FullDayTimeOffMarkers,
		
		Deletes: u.Deletes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorychange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
