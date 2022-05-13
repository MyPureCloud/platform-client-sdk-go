package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Availabletime
type Availabletime struct { 
	// DateStart - Start of the availability period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`


	// LengthInMinutes - Length of availability period in minutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// IsPaid - Indicates if this availability period is paid in Workforce Management schedule
	IsPaid *bool `json:"isPaid,omitempty"`


	// ActivityCategory - Workforce Management activity category for this availability period
	ActivityCategory *string `json:"activityCategory,omitempty"`


	// WfmSchedule - Workforce Management schedule information associated with the available time
	WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`

}

func (o *Availabletime) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Availabletime
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		IsPaid *bool `json:"isPaid,omitempty"`
		
		ActivityCategory *string `json:"activityCategory,omitempty"`
		
		WfmSchedule *Wfmschedulereference `json:"wfmSchedule,omitempty"`
		*Alias
	}{ 
		DateStart: DateStart,
		
		LengthInMinutes: o.LengthInMinutes,
		
		IsPaid: o.IsPaid,
		
		ActivityCategory: o.ActivityCategory,
		
		WfmSchedule: o.WfmSchedule,
		Alias:    (*Alias)(o),
	})
}

func (o *Availabletime) UnmarshalJSON(b []byte) error {
	var AvailabletimeMap map[string]interface{}
	err := json.Unmarshal(b, &AvailabletimeMap)
	if err != nil {
		return err
	}
	
	if dateStartString, ok := AvailabletimeMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if LengthInMinutes, ok := AvailabletimeMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if IsPaid, ok := AvailabletimeMap["isPaid"].(bool); ok {
		o.IsPaid = &IsPaid
	}
    
	if ActivityCategory, ok := AvailabletimeMap["activityCategory"].(string); ok {
		o.ActivityCategory = &ActivityCategory
	}
    
	if WfmSchedule, ok := AvailabletimeMap["wfmSchedule"].(map[string]interface{}); ok {
		WfmScheduleString, _ := json.Marshal(WfmSchedule)
		json.Unmarshal(WfmScheduleString, &o.WfmSchedule)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Availabletime) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
