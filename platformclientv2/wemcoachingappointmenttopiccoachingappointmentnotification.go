package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wemcoachingappointmenttopiccoachingappointmentnotification
type Wemcoachingappointmenttopiccoachingappointmentnotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateStart
	DateStart *time.Time `json:"dateStart,omitempty"`


	// LengthInMinutes
	LengthInMinutes *int32 `json:"lengthInMinutes,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// Facilitator
	Facilitator *Wemcoachingappointmenttopicuserreference `json:"facilitator,omitempty"`


	// Attendees
	Attendees *[]Wemcoachingappointmenttopicuserreference `json:"attendees,omitempty"`


	// CreatedBy
	CreatedBy *Wemcoachingappointmenttopicuserreference `json:"createdBy,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ModifiedBy
	ModifiedBy *Wemcoachingappointmenttopicuserreference `json:"modifiedBy,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Conversations
	Conversations *[]Wemcoachingappointmenttopiccoachingappointmentconversation `json:"conversations,omitempty"`


	// Documents
	Documents *[]Wemcoachingappointmenttopiccoachingappointmentdocument `json:"documents,omitempty"`


	// ChangeType
	ChangeType *string `json:"changeType,omitempty"`


	// DateCompleted
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopiccoachingappointmentnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
