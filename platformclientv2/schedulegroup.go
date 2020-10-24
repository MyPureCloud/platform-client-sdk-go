package platformclientv2
import (
	"time"
	"encoding/json"
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
	Version *int32 `json:"version,omitempty"`


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

// String returns a JSON representation of the model
func (o *Schedulegroup) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
