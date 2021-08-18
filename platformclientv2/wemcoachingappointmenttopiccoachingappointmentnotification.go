package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


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


	// ExternalLinks
	ExternalLinks *[]Wemcoachingappointmenttopiccoachingappointmentexternallink `json:"externalLinks,omitempty"`

}

func (u *Wemcoachingappointmenttopiccoachingappointmentnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wemcoachingappointmenttopiccoachingappointmentnotification

	
	DateStart := new(string)
	if u.DateStart != nil {
		
		*DateStart = timeutil.Strftime(u.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateCompleted := new(string)
	if u.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(u.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Facilitator *Wemcoachingappointmenttopicuserreference `json:"facilitator,omitempty"`
		
		Attendees *[]Wemcoachingappointmenttopicuserreference `json:"attendees,omitempty"`
		
		CreatedBy *Wemcoachingappointmenttopicuserreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ModifiedBy *Wemcoachingappointmenttopicuserreference `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Conversations *[]Wemcoachingappointmenttopiccoachingappointmentconversation `json:"conversations,omitempty"`
		
		Documents *[]Wemcoachingappointmenttopiccoachingappointmentdocument `json:"documents,omitempty"`
		
		ChangeType *string `json:"changeType,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		ExternalLinks *[]Wemcoachingappointmenttopiccoachingappointmentexternallink `json:"externalLinks,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DateStart: DateStart,
		
		LengthInMinutes: u.LengthInMinutes,
		
		Status: u.Status,
		
		Facilitator: u.Facilitator,
		
		Attendees: u.Attendees,
		
		CreatedBy: u.CreatedBy,
		
		DateCreated: DateCreated,
		
		ModifiedBy: u.ModifiedBy,
		
		DateModified: DateModified,
		
		Conversations: u.Conversations,
		
		Documents: u.Documents,
		
		ChangeType: u.ChangeType,
		
		DateCompleted: DateCompleted,
		
		ExternalLinks: u.ExternalLinks,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopiccoachingappointmentnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
