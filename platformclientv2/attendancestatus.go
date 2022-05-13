package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Attendancestatus
type Attendancestatus struct { 
	// DateWorkday - the workday date of this attendance status. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`


	// AttendanceStatusType - the attendance status
	AttendanceStatusType *string `json:"attendanceStatusType,omitempty"`

}

func (o *Attendancestatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Attendancestatus
	
	DateWorkday := new(string)
	if o.DateWorkday != nil {
		*DateWorkday = timeutil.Strftime(o.DateWorkday, "%Y-%m-%d")
	} else {
		DateWorkday = nil
	}
	
	return json.Marshal(&struct { 
		DateWorkday *string `json:"dateWorkday,omitempty"`
		
		AttendanceStatusType *string `json:"attendanceStatusType,omitempty"`
		*Alias
	}{ 
		DateWorkday: DateWorkday,
		
		AttendanceStatusType: o.AttendanceStatusType,
		Alias:    (*Alias)(o),
	})
}

func (o *Attendancestatus) UnmarshalJSON(b []byte) error {
	var AttendancestatusMap map[string]interface{}
	err := json.Unmarshal(b, &AttendancestatusMap)
	if err != nil {
		return err
	}
	
	if dateWorkdayString, ok := AttendancestatusMap["dateWorkday"].(string); ok {
		DateWorkday, _ := time.Parse("2006-01-02", dateWorkdayString)
		o.DateWorkday = &DateWorkday
	}
	
	if AttendanceStatusType, ok := AttendancestatusMap["attendanceStatusType"].(string); ok {
		o.AttendanceStatusType = &AttendanceStatusType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Attendancestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
