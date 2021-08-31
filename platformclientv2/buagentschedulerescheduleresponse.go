package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentschedulerescheduleresponse
type Buagentschedulerescheduleresponse struct { 
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

}

func (o *Buagentschedulerescheduleresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentschedulerescheduleresponse
	
	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`
		
		FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		
		WorkPlan *Workplanreference `json:"workPlan,omitempty"`
		
		WorkPlansPerWeek *[]Workplanreference `json:"workPlansPerWeek,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		Shifts: o.Shifts,
		
		FullDayTimeOffMarkers: o.FullDayTimeOffMarkers,
		
		WorkPlan: o.WorkPlan,
		
		WorkPlansPerWeek: o.WorkPlansPerWeek,
		Alias:    (*Alias)(o),
	})
}

func (o *Buagentschedulerescheduleresponse) UnmarshalJSON(b []byte) error {
	var BuagentschedulerescheduleresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentschedulerescheduleresponseMap)
	if err != nil {
		return err
	}
	
	if User, ok := BuagentschedulerescheduleresponseMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Shifts, ok := BuagentschedulerescheduleresponseMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if FullDayTimeOffMarkers, ok := BuagentschedulerescheduleresponseMap["fullDayTimeOffMarkers"].([]interface{}); ok {
		FullDayTimeOffMarkersString, _ := json.Marshal(FullDayTimeOffMarkers)
		json.Unmarshal(FullDayTimeOffMarkersString, &o.FullDayTimeOffMarkers)
	}
	
	if WorkPlan, ok := BuagentschedulerescheduleresponseMap["workPlan"].(map[string]interface{}); ok {
		WorkPlanString, _ := json.Marshal(WorkPlan)
		json.Unmarshal(WorkPlanString, &o.WorkPlan)
	}
	
	if WorkPlansPerWeek, ok := BuagentschedulerescheduleresponseMap["workPlansPerWeek"].([]interface{}); ok {
		WorkPlansPerWeekString, _ := json.Marshal(WorkPlansPerWeek)
		json.Unmarshal(WorkPlansPerWeekString, &o.WorkPlansPerWeek)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentschedulerescheduleresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
