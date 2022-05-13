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

func (o *Coachingappointmentstatusresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingappointmentstatusresponse
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Appointment: o.Appointment,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Coachingappointmentstatusresponse) UnmarshalJSON(b []byte) error {
	var CoachingappointmentstatusresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CoachingappointmentstatusresponseMap)
	if err != nil {
		return err
	}
	
	if Appointment, ok := CoachingappointmentstatusresponseMap["appointment"].(map[string]interface{}); ok {
		AppointmentString, _ := json.Marshal(Appointment)
		json.Unmarshal(AppointmentString, &o.Appointment)
	}
	
	if CreatedBy, ok := CoachingappointmentstatusresponseMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := CoachingappointmentstatusresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if Status, ok := CoachingappointmentstatusresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Coachingappointmentstatusresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
