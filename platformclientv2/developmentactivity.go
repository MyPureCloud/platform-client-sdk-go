package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivity - Development Activity object
type Developmentactivity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// DateCompleted - Date that activity was completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// CreatedBy - User that created activity
	CreatedBy *Userreference `json:"createdBy,omitempty"`


	// DateCreated - Date activity was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// Name - The name of the activity
	Name *string `json:"name,omitempty"`


	// VarType - The type of activity
	VarType *string `json:"type,omitempty"`


	// Status - The status of the activity
	Status *string `json:"status,omitempty"`


	// DateDue - Due date for completion of the activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDue *time.Time `json:"dateDue,omitempty"`


	// Facilitator - Facilitator of the activity
	Facilitator *Userreference `json:"facilitator,omitempty"`


	// Attendees - List of users attending the activity
	Attendees *[]Userreference `json:"attendees,omitempty"`


	// IsOverdue - Indicates if the activity is overdue
	IsOverdue *bool `json:"isOverdue,omitempty"`

}

func (u *Developmentactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Developmentactivity

	
	DateCompleted := new(string)
	if u.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(u.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateDue := new(string)
	if u.DateDue != nil {
		
		*DateDue = timeutil.Strftime(u.DateDue, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDue = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		DateDue *string `json:"dateDue,omitempty"`
		
		Facilitator *Userreference `json:"facilitator,omitempty"`
		
		Attendees *[]Userreference `json:"attendees,omitempty"`
		
		IsOverdue *bool `json:"isOverdue,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		DateCompleted: DateCompleted,
		
		CreatedBy: u.CreatedBy,
		
		DateCreated: DateCreated,
		
		SelfUri: u.SelfUri,
		
		Name: u.Name,
		
		VarType: u.VarType,
		
		Status: u.Status,
		
		DateDue: DateDue,
		
		Facilitator: u.Facilitator,
		
		Attendees: u.Attendees,
		
		IsOverdue: u.IsOverdue,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Developmentactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
