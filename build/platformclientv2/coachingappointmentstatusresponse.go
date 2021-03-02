package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Coachingappointmentstatusresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
