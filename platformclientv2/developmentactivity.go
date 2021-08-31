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

func (o *Developmentactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Developmentactivity
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateDue := new(string)
	if o.DateDue != nil {
		
		*DateDue = timeutil.Strftime(o.DateDue, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		DateCompleted: DateCompleted,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		SelfUri: o.SelfUri,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		Status: o.Status,
		
		DateDue: DateDue,
		
		Facilitator: o.Facilitator,
		
		Attendees: o.Attendees,
		
		IsOverdue: o.IsOverdue,
		Alias:    (*Alias)(o),
	})
}

func (o *Developmentactivity) UnmarshalJSON(b []byte) error {
	var DevelopmentactivityMap map[string]interface{}
	err := json.Unmarshal(b, &DevelopmentactivityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DevelopmentactivityMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if dateCompletedString, ok := DevelopmentactivityMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if CreatedBy, ok := DevelopmentactivityMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := DevelopmentactivityMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if SelfUri, ok := DevelopmentactivityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if Name, ok := DevelopmentactivityMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if VarType, ok := DevelopmentactivityMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Status, ok := DevelopmentactivityMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if dateDueString, ok := DevelopmentactivityMap["dateDue"].(string); ok {
		DateDue, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDueString)
		o.DateDue = &DateDue
	}
	
	if Facilitator, ok := DevelopmentactivityMap["facilitator"].(map[string]interface{}); ok {
		FacilitatorString, _ := json.Marshal(Facilitator)
		json.Unmarshal(FacilitatorString, &o.Facilitator)
	}
	
	if Attendees, ok := DevelopmentactivityMap["attendees"].([]interface{}); ok {
		AttendeesString, _ := json.Marshal(Attendees)
		json.Unmarshal(AttendeesString, &o.Attendees)
	}
	
	if IsOverdue, ok := DevelopmentactivityMap["isOverdue"].(bool); ok {
		o.IsOverdue = &IsOverdue
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Developmentactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
