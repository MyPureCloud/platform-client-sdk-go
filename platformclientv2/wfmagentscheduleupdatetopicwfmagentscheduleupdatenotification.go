package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification
type Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification struct { 
	// User
	User *Wfmagentscheduleupdatetopicuserreference `json:"user,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`


	// Shifts
	Shifts *[]Wfmagentscheduleupdatetopicwfmscheduleshift `json:"shifts,omitempty"`


	// FullDayTimeOffMarkers
	FullDayTimeOffMarkers *[]Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`


	// Updates
	Updates *[]Wfmagentscheduleupdatetopicwfmagentscheduleupdate `json:"updates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
