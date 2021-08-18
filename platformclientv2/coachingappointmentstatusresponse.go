package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingappointmentstatusresponse
type Coachingappointmentstatusresponse struct { 
	// Appointment - The coaching appointment this status belongs to
	Appointment *Coachingappointmentreference `json:"appointment,omitempty"`


	// CreatedBy - User who updated the status
	CreatedBy *Userreference `json:"createdBy,omitempty"`


	// DateCreated - Creation time of the status. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// Status - The status of the coaching appointment
	Status *string `json:"status,omitempty"`

}

func (u *Coachingappointmentstatusresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingappointmentstatusresponse

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	

	return json.Marshal(&struct { 
		Appointment *Coachingappointmentreference `json:"appointment,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Appointment: u.Appointment,
		
		CreatedBy: u.CreatedBy,
		
		DateCreated: DateCreated,
		
		Status: u.Status,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Coachingappointmentstatusresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
