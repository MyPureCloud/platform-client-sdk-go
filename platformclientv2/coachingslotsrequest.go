package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingslotsrequest
type Coachingslotsrequest struct { 
	// Interval - Range of time to get slots for scheduling coaching appointments. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`


	// LengthInMinutes - The duration of coaching appointment to schedule in 15 minutes granularity up to maximum of 60 minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// AttendeeIds - List of attendees to determine coaching appointment slots
	AttendeeIds *[]string `json:"attendeeIds,omitempty"`


	// FacilitatorIds - List of facilitators to determine coaching appointment slots
	FacilitatorIds *[]string `json:"facilitatorIds,omitempty"`


	// InterruptibleAppointmentIds - List of appointment ids to exclude from consideration when determining blocked slots
	InterruptibleAppointmentIds *[]string `json:"interruptibleAppointmentIds,omitempty"`

}

func (o *Coachingslotsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingslotsrequest
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		AttendeeIds *[]string `json:"attendeeIds,omitempty"`
		
		FacilitatorIds *[]string `json:"facilitatorIds,omitempty"`
		
		InterruptibleAppointmentIds *[]string `json:"interruptibleAppointmentIds,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		LengthInMinutes: o.LengthInMinutes,
		
		AttendeeIds: o.AttendeeIds,
		
		FacilitatorIds: o.FacilitatorIds,
		
		InterruptibleAppointmentIds: o.InterruptibleAppointmentIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Coachingslotsrequest) UnmarshalJSON(b []byte) error {
	var CoachingslotsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CoachingslotsrequestMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := CoachingslotsrequestMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if LengthInMinutes, ok := CoachingslotsrequestMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if AttendeeIds, ok := CoachingslotsrequestMap["attendeeIds"].([]interface{}); ok {
		AttendeeIdsString, _ := json.Marshal(AttendeeIds)
		json.Unmarshal(AttendeeIdsString, &o.AttendeeIds)
	}
	
	if FacilitatorIds, ok := CoachingslotsrequestMap["facilitatorIds"].([]interface{}); ok {
		FacilitatorIdsString, _ := json.Marshal(FacilitatorIds)
		json.Unmarshal(FacilitatorIdsString, &o.FacilitatorIds)
	}
	
	if InterruptibleAppointmentIds, ok := CoachingslotsrequestMap["interruptibleAppointmentIds"].([]interface{}); ok {
		InterruptibleAppointmentIdsString, _ := json.Marshal(InterruptibleAppointmentIds)
		json.Unmarshal(InterruptibleAppointmentIdsString, &o.InterruptibleAppointmentIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Coachingslotsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
