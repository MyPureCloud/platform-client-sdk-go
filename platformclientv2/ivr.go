package platformclientv2
import (
	"time"
	"encoding/json"
)

// Ivr - Defines the phone numbers, operating hours, and the Architect flows to execute for an IVR.
type Ivr struct { 
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


	// Dnis - The phone number(s) to contact the IVR by.  Each phone number must be unique and not in use by another resource.  For example, a user and an iVR cannot have the same phone number.
	Dnis *[]string `json:"dnis,omitempty"`


	// OpenHoursFlow - The Architect flow to execute during the hours an organization is open.
	OpenHoursFlow *Domainentityref `json:"openHoursFlow,omitempty"`


	// ClosedHoursFlow - The Architect flow to execute during the hours an organization is closed.
	ClosedHoursFlow *Domainentityref `json:"closedHoursFlow,omitempty"`


	// HolidayHoursFlow - The Architect flow to execute during an organization's holiday hours.
	HolidayHoursFlow *Domainentityref `json:"holidayHoursFlow,omitempty"`


	// ScheduleGroup - The schedule group defining the open and closed hours for an organization.  If this is provided, an open flow and a closed flow must be specified as well.
	ScheduleGroup *Domainentityref `json:"scheduleGroup,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ivr) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
