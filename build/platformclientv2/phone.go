package platformclientv2
import (
	"time"
	"encoding/json"
)

// Phone
type Phone struct { 
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


	// Site - The site associated to the phone.
	Site *Domainentityref `json:"site,omitempty"`


	// PhoneBaseSettings - Phone Base Settings
	PhoneBaseSettings *Domainentityref `json:"phoneBaseSettings,omitempty"`


	// LineBaseSettings
	LineBaseSettings *Domainentityref `json:"lineBaseSettings,omitempty"`


	// PhoneMetaBase
	PhoneMetaBase *Domainentityref `json:"phoneMetaBase,omitempty"`


	// Lines - Lines
	Lines *[]Line `json:"lines,omitempty"`


	// Status - The status of the phone and lines from the primary Edge.
	Status *Phonestatus `json:"status,omitempty"`


	// SecondaryStatus - The status of the phone and lines from the secondary Edge.
	SecondaryStatus *Phonestatus `json:"secondaryStatus,omitempty"`


	// UserAgentInfo - User Agent Information for this phone. This includes model, firmware version, and manufacturer.
	UserAgentInfo *Useragentinfo `json:"userAgentInfo,omitempty"`


	// Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`


	// Capabilities
	Capabilities *Phonecapabilities `json:"capabilities,omitempty"`


	// WebRtcUser - This is the user associated with a WebRTC type phone.  It is required for all WebRTC phones.
	WebRtcUser *Domainentityref `json:"webRtcUser,omitempty"`


	// PrimaryEdge
	PrimaryEdge *Edge `json:"primaryEdge,omitempty"`


	// SecondaryEdge
	SecondaryEdge *Edge `json:"secondaryEdge,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phone) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
