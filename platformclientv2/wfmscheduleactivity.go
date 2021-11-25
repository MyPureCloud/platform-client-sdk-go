package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmscheduleactivity
type Wfmscheduleactivity struct { 
	// UserReference - ID of user that the schedule is for
	UserReference *Userreference `json:"userReference,omitempty"`


	// Activities - List of user's scheduled activities
	Activities *[]Scheduleactivity `json:"activities,omitempty"`


	// FullDayTimeOffMarkers - List of user's days off
	FullDayTimeOffMarkers *[]Fulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`

}

func (o *Wfmscheduleactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmscheduleactivity
	
	return json.Marshal(&struct { 
		UserReference *Userreference `json:"userReference,omitempty"`
		
		Activities *[]Scheduleactivity `json:"activities,omitempty"`
		
		FullDayTimeOffMarkers *[]Fulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		*Alias
	}{ 
		UserReference: o.UserReference,
		
		Activities: o.Activities,
		
		FullDayTimeOffMarkers: o.FullDayTimeOffMarkers,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmscheduleactivity) UnmarshalJSON(b []byte) error {
	var WfmscheduleactivityMap map[string]interface{}
	err := json.Unmarshal(b, &WfmscheduleactivityMap)
	if err != nil {
		return err
	}
	
	if UserReference, ok := WfmscheduleactivityMap["userReference"].(map[string]interface{}); ok {
		UserReferenceString, _ := json.Marshal(UserReference)
		json.Unmarshal(UserReferenceString, &o.UserReference)
	}
	
	if Activities, ok := WfmscheduleactivityMap["activities"].([]interface{}); ok {
		ActivitiesString, _ := json.Marshal(Activities)
		json.Unmarshal(ActivitiesString, &o.Activities)
	}
	
	if FullDayTimeOffMarkers, ok := WfmscheduleactivityMap["fullDayTimeOffMarkers"].([]interface{}); ok {
		FullDayTimeOffMarkersString, _ := json.Marshal(FullDayTimeOffMarkers)
		json.Unmarshal(FullDayTimeOffMarkersString, &o.FullDayTimeOffMarkers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmscheduleactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
