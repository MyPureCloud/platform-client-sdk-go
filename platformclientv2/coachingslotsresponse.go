package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingslotsresponse
type Coachingslotsresponse struct { 
	// SuggestedSlots - List of slots where coaching appointment can be scheduled
	SuggestedSlots *[]Coachingslot `json:"suggestedSlots,omitempty"`


	// AttendeeSchedules - Periods of availability for attendees to schedule coaching appointment
	AttendeeSchedules *[]Useravailabletimes `json:"attendeeSchedules,omitempty"`


	// FacilitatorSchedules - Periods of availability for facilitators to schedule coaching appointment
	FacilitatorSchedules *[]Useravailabletimes `json:"facilitatorSchedules,omitempty"`


	// WfmScheduleActivities - Detailed data for WFM scheduled activities
	WfmScheduleActivities *[]Wfmscheduleactivity `json:"wfmScheduleActivities,omitempty"`

}

func (o *Coachingslotsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingslotsresponse
	
	return json.Marshal(&struct { 
		SuggestedSlots *[]Coachingslot `json:"suggestedSlots,omitempty"`
		
		AttendeeSchedules *[]Useravailabletimes `json:"attendeeSchedules,omitempty"`
		
		FacilitatorSchedules *[]Useravailabletimes `json:"facilitatorSchedules,omitempty"`
		
		WfmScheduleActivities *[]Wfmscheduleactivity `json:"wfmScheduleActivities,omitempty"`
		*Alias
	}{ 
		SuggestedSlots: o.SuggestedSlots,
		
		AttendeeSchedules: o.AttendeeSchedules,
		
		FacilitatorSchedules: o.FacilitatorSchedules,
		
		WfmScheduleActivities: o.WfmScheduleActivities,
		Alias:    (*Alias)(o),
	})
}

func (o *Coachingslotsresponse) UnmarshalJSON(b []byte) error {
	var CoachingslotsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CoachingslotsresponseMap)
	if err != nil {
		return err
	}
	
	if SuggestedSlots, ok := CoachingslotsresponseMap["suggestedSlots"].([]interface{}); ok {
		SuggestedSlotsString, _ := json.Marshal(SuggestedSlots)
		json.Unmarshal(SuggestedSlotsString, &o.SuggestedSlots)
	}
	
	if AttendeeSchedules, ok := CoachingslotsresponseMap["attendeeSchedules"].([]interface{}); ok {
		AttendeeSchedulesString, _ := json.Marshal(AttendeeSchedules)
		json.Unmarshal(AttendeeSchedulesString, &o.AttendeeSchedules)
	}
	
	if FacilitatorSchedules, ok := CoachingslotsresponseMap["facilitatorSchedules"].([]interface{}); ok {
		FacilitatorSchedulesString, _ := json.Marshal(FacilitatorSchedules)
		json.Unmarshal(FacilitatorSchedulesString, &o.FacilitatorSchedules)
	}
	
	if WfmScheduleActivities, ok := CoachingslotsresponseMap["wfmScheduleActivities"].([]interface{}); ok {
		WfmScheduleActivitiesString, _ := json.Marshal(WfmScheduleActivities)
		json.Unmarshal(WfmScheduleActivitiesString, &o.WfmScheduleActivities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Coachingslotsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
