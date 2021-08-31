package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Buagentschedulesearchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentschedulesearchresponse
	
	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`
		
		FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		Shifts: o.Shifts,
		
		FullDayTimeOffMarkers: o.FullDayTimeOffMarkers,
		Alias:    (*Alias)(o),
	})
}

func (o *Buagentschedulesearchresponse) UnmarshalJSON(b []byte) error {
	var BuagentschedulesearchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentschedulesearchresponseMap)
	if err != nil {
		return err
	}
	
	if User, ok := BuagentschedulesearchresponseMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Shifts, ok := BuagentschedulesearchresponseMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if FullDayTimeOffMarkers, ok := BuagentschedulesearchresponseMap["fullDayTimeOffMarkers"].([]interface{}); ok {
		FullDayTimeOffMarkersString, _ := json.Marshal(FullDayTimeOffMarkers)
		json.Unmarshal(FullDayTimeOffMarkersString, &o.FullDayTimeOffMarkers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentschedulesearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
