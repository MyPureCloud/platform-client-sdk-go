package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buimportagentscheduleuploadschema
type Buimportagentscheduleuploadschema struct { 
	// UserId - The ID of the user to whom this agent schedule applies
	UserId *string `json:"userId,omitempty"`


	// WorkPlanId - The ID of the work plan for this user.  Mutually exclusive with workPlanIdsPerWeek
	WorkPlanId *Valuewrapperstring `json:"workPlanId,omitempty"`


	// WorkPlanIdsPerWeek - The IDs of the work plans per week for this user.  Mutually exclusive with workPlanId
	WorkPlanIdsPerWeek *Listwrapperstring `json:"workPlanIdsPerWeek,omitempty"`


	// Shifts - The shift definitions for this agent schedule
	Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`


	// FullDayTimeOffMarkers - Any full day time off markers that apply to this agent schedule
	FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`

}

func (o *Buimportagentscheduleuploadschema) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buimportagentscheduleuploadschema
	
	return json.Marshal(&struct { 
		UserId *string `json:"userId,omitempty"`
		
		WorkPlanId *Valuewrapperstring `json:"workPlanId,omitempty"`
		
		WorkPlanIdsPerWeek *Listwrapperstring `json:"workPlanIdsPerWeek,omitempty"`
		
		Shifts *[]Buagentscheduleshift `json:"shifts,omitempty"`
		
		FullDayTimeOffMarkers *[]Bufulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		*Alias
	}{ 
		UserId: o.UserId,
		
		WorkPlanId: o.WorkPlanId,
		
		WorkPlanIdsPerWeek: o.WorkPlanIdsPerWeek,
		
		Shifts: o.Shifts,
		
		FullDayTimeOffMarkers: o.FullDayTimeOffMarkers,
		Alias:    (*Alias)(o),
	})
}

func (o *Buimportagentscheduleuploadschema) UnmarshalJSON(b []byte) error {
	var BuimportagentscheduleuploadschemaMap map[string]interface{}
	err := json.Unmarshal(b, &BuimportagentscheduleuploadschemaMap)
	if err != nil {
		return err
	}
	
	if UserId, ok := BuimportagentscheduleuploadschemaMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if WorkPlanId, ok := BuimportagentscheduleuploadschemaMap["workPlanId"].(map[string]interface{}); ok {
		WorkPlanIdString, _ := json.Marshal(WorkPlanId)
		json.Unmarshal(WorkPlanIdString, &o.WorkPlanId)
	}
	
	if WorkPlanIdsPerWeek, ok := BuimportagentscheduleuploadschemaMap["workPlanIdsPerWeek"].(map[string]interface{}); ok {
		WorkPlanIdsPerWeekString, _ := json.Marshal(WorkPlanIdsPerWeek)
		json.Unmarshal(WorkPlanIdsPerWeekString, &o.WorkPlanIdsPerWeek)
	}
	
	if Shifts, ok := BuimportagentscheduleuploadschemaMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if FullDayTimeOffMarkers, ok := BuimportagentscheduleuploadschemaMap["fullDayTimeOffMarkers"].([]interface{}); ok {
		FullDayTimeOffMarkersString, _ := json.Marshal(FullDayTimeOffMarkers)
		json.Unmarshal(FullDayTimeOffMarkersString, &o.FullDayTimeOffMarkers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buimportagentscheduleuploadschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
