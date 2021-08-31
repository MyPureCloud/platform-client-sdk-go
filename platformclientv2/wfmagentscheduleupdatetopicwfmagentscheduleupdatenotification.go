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

func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		User: o.User,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		Shifts: o.Shifts,
		
		FullDayTimeOffMarkers: o.FullDayTimeOffMarkers,
		
		Updates: o.Updates,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification) UnmarshalJSON(b []byte) error {
	var WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap)
	if err != nil {
		return err
	}
	
	if User, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if startDateString, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if Shifts, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if FullDayTimeOffMarkers, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap["fullDayTimeOffMarkers"].([]interface{}); ok {
		FullDayTimeOffMarkersString, _ := json.Marshal(FullDayTimeOffMarkers)
		json.Unmarshal(FullDayTimeOffMarkersString, &o.FullDayTimeOffMarkers)
	}
	
	if Updates, ok := WfmagentscheduleupdatetopicwfmagentscheduleupdatenotificationMap["updates"].([]interface{}); ok {
		UpdatesString, _ := json.Marshal(Updates)
		json.Unmarshal(UpdatesString, &o.Updates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmagentscheduleupdatenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
