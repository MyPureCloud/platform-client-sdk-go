package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if u.EndDate != nil {
		
		*EndDate = timeutil.Strftime(u.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	

	return json.Marshal(&struct { 
		User *Wfmagentscheduleupdatetopicuserreference `json:"user,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		Shifts *[]Wfmagentscheduleupdatetopicwfmscheduleshift `json:"shifts,omitempty"`
		
		FullDayTimeOffMarkers *[]Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		
		Updates *[]Wfmagentscheduleupdatetopicwfmagentscheduleupdate `json:"updates,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		Shifts: u.Shifts,
		
		FullDayTimeOffMarkers: u.FullDayTimeOffMarkers,
		
		Updates: u.Updates,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
