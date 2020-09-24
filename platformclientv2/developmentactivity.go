package platformclientv2
import (
	"time"
	"encoding/json"
)

// Developmentactivity - Development Activity object
type Developmentactivity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// DateCompleted - Date that activity was completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// CreatedBy - User that created activity
	CreatedBy *Userreference `json:"createdBy,omitempty"`


	// DateCreated - Date activity was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// Name - The name of the activity
	Name *string `json:"name,omitempty"`


	// VarType - The type of activity
	VarType *string `json:"type,omitempty"`


	// Status - The status of the activity
	Status *string `json:"status,omitempty"`


	// DateDue - Due date for completion of the activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateDue *time.Time `json:"dateDue,omitempty"`


	// Facilitator - Facilitator of the activity
	Facilitator *Userreference `json:"facilitator,omitempty"`


	// Attendees - List of users attending the activity
	Attendees *[]Userreference `json:"attendees,omitempty"`


	// IsOverdue - Indicates if the activity is overdue
	IsOverdue *bool `json:"isOverdue,omitempty"`

}

// String returns a JSON representation of the model
func (o *Developmentactivity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
