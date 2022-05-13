package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingslot
type Coachingslot struct { 
	// DateStart - Start date and time of scheduled coaching appointment slot. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`


	// LengthInMinutes - Length of coaching appointment slot in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// StaffingDifference - Difference between scheduled and forecast headcount for this slot after scheduling the coaching appointment
	StaffingDifference *float64 `json:"staffingDifference,omitempty"`


	// DifferenceRating - Rating based on the staffing difference for scheduled slot
	DifferenceRating *string `json:"differenceRating,omitempty"`


	// WfmSchedule - Workforce Management schedule information associated with the slot
	WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`

}

func (o *Coachingslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingslot
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		StaffingDifference *float64 `json:"staffingDifference,omitempty"`
		
		DifferenceRating *string `json:"differenceRating,omitempty"`
		
		WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`
		*Alias
	}{ 
		DateStart: DateStart,
		
		LengthInMinutes: o.LengthInMinutes,
		
		StaffingDifference: o.StaffingDifference,
		
		DifferenceRating: o.DifferenceRating,
		
		WfmSchedule: o.WfmSchedule,
		Alias:    (*Alias)(o),
	})
}

func (o *Coachingslot) UnmarshalJSON(b []byte) error {
	var CoachingslotMap map[string]interface{}
	err := json.Unmarshal(b, &CoachingslotMap)
	if err != nil {
		return err
	}
	
	if dateStartString, ok := CoachingslotMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if LengthInMinutes, ok := CoachingslotMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if StaffingDifference, ok := CoachingslotMap["staffingDifference"].(float64); ok {
		o.StaffingDifference = &StaffingDifference
	}
    
	if DifferenceRating, ok := CoachingslotMap["differenceRating"].(string); ok {
		o.DifferenceRating = &DifferenceRating
	}
    
	if WfmSchedule, ok := CoachingslotMap["wfmSchedule"].(map[string]interface{}); ok {
		WfmScheduleString, _ := json.Marshal(WfmSchedule)
		json.Unmarshal(WfmScheduleString, &o.WfmSchedule)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Coachingslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
