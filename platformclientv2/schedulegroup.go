package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulegroup - A group of schedules that define the operating hours of an organization.
type Schedulegroup struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`


	// Description - The resource's description.
	Description *string `json:"description,omitempty"`


	// Version - The current version of the resource.
	Version *int `json:"version,omitempty"`


	// DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - The ID of the user that last modified the resource.
	ModifiedBy *string `json:"modifiedBy,omitempty"`


	// CreatedBy - The ID of the user that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`


	// State - Indicates if the resource is active, inactive, or deleted.
	State *string `json:"state,omitempty"`


	// ModifiedByApp - The application that last modified the resource.
	ModifiedByApp *string `json:"modifiedByApp,omitempty"`


	// CreatedByApp - The application that created the resource.
	CreatedByApp *string `json:"createdByApp,omitempty"`


	// TimeZone - The timezone the schedules are a part of.  This is not a schedule property to allow a schedule to be used in multiple timezones.
	TimeZone *string `json:"timeZone,omitempty"`


	// OpenSchedules - The schedules defining the hours an organization is open.
	OpenSchedules *[]Domainentityref `json:"openSchedules,omitempty"`


	// ClosedSchedules - The schedules defining the hours an organization is closed.
	ClosedSchedules *[]Domainentityref `json:"closedSchedules,omitempty"`


	// HolidaySchedules - The schedules defining the hours an organization is closed for the holidays.
	HolidaySchedules *[]Domainentityref `json:"holidaySchedules,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Schedulegroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulegroup
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ModifiedByApp *string `json:"modifiedByApp,omitempty"`
		
		CreatedByApp *string `json:"createdByApp,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		OpenSchedules *[]Domainentityref `json:"openSchedules,omitempty"`
		
		ClosedSchedules *[]Domainentityref `json:"closedSchedules,omitempty"`
		
		HolidaySchedules *[]Domainentityref `json:"holidaySchedules,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		CreatedBy: o.CreatedBy,
		
		State: o.State,
		
		ModifiedByApp: o.ModifiedByApp,
		
		CreatedByApp: o.CreatedByApp,
		
		TimeZone: o.TimeZone,
		
		OpenSchedules: o.OpenSchedules,
		
		ClosedSchedules: o.ClosedSchedules,
		
		HolidaySchedules: o.HolidaySchedules,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Schedulegroup) UnmarshalJSON(b []byte) error {
	var SchedulegroupMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulegroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SchedulegroupMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := SchedulegroupMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := SchedulegroupMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := SchedulegroupMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := SchedulegroupMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := SchedulegroupMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := SchedulegroupMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := SchedulegroupMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := SchedulegroupMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := SchedulegroupMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := SchedulegroupMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if TimeZone, ok := SchedulegroupMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
	
	if OpenSchedules, ok := SchedulegroupMap["openSchedules"].([]interface{}); ok {
		OpenSchedulesString, _ := json.Marshal(OpenSchedules)
		json.Unmarshal(OpenSchedulesString, &o.OpenSchedules)
	}
	
	if ClosedSchedules, ok := SchedulegroupMap["closedSchedules"].([]interface{}); ok {
		ClosedSchedulesString, _ := json.Marshal(ClosedSchedules)
		json.Unmarshal(ClosedSchedulesString, &o.ClosedSchedules)
	}
	
	if HolidaySchedules, ok := SchedulegroupMap["holidaySchedules"].([]interface{}); ok {
		HolidaySchedulesString, _ := json.Marshal(HolidaySchedules)
		json.Unmarshal(HolidaySchedulesString, &o.HolidaySchedules)
	}
	
	if SelfUri, ok := SchedulegroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulegroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
