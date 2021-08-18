package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysegment
type Journeysegment struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// IsActive - Whether or not the segment is active.
	IsActive *bool `json:"isActive,omitempty"`


	// DisplayName - The display name of the segment.
	DisplayName *string `json:"displayName,omitempty"`


	// Version - The version of the segment.
	Version *int `json:"version,omitempty"`


	// Description - A description of the segment.
	Description *string `json:"description,omitempty"`


	// Color - The hexadecimal color value of the segment.
	Color *string `json:"color,omitempty"`


	// Scope - The target entity that a segment applies to.
	Scope *string `json:"scope,omitempty"`


	// ShouldDisplayToAgent - Whether or not the segment should be displayed to agent/supervisor users.
	ShouldDisplayToAgent *bool `json:"shouldDisplayToAgent,omitempty"`


	// Context - The context of the segment.
	Context *Context `json:"context,omitempty"`


	// Journey - The pattern of rules defining the segment.
	Journey *Journey `json:"journey,omitempty"`


	// ExternalSegment - Details of an entity corresponding to this segment in an external system.
	ExternalSegment *Externalsegment `json:"externalSegment,omitempty"`


	// AssignmentExpirationDays - Time, in days, from when the segment is assigned until it is automatically unassigned.
	AssignmentExpirationDays *int `json:"assignmentExpirationDays,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// CreatedDate - Timestamp indicating when the segment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Timestamp indicating when the the segment was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

func (u *Journeysegment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysegment

	
	CreatedDate := new(string)
	if u.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(u.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		IsActive *bool `json:"isActive,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Color *string `json:"color,omitempty"`
		
		Scope *string `json:"scope,omitempty"`
		
		ShouldDisplayToAgent *bool `json:"shouldDisplayToAgent,omitempty"`
		
		Context *Context `json:"context,omitempty"`
		
		Journey *Journey `json:"journey,omitempty"`
		
		ExternalSegment *Externalsegment `json:"externalSegment,omitempty"`
		
		AssignmentExpirationDays *int `json:"assignmentExpirationDays,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		IsActive: u.IsActive,
		
		DisplayName: u.DisplayName,
		
		Version: u.Version,
		
		Description: u.Description,
		
		Color: u.Color,
		
		Scope: u.Scope,
		
		ShouldDisplayToAgent: u.ShouldDisplayToAgent,
		
		Context: u.Context,
		
		Journey: u.Journey,
		
		ExternalSegment: u.ExternalSegment,
		
		AssignmentExpirationDays: u.AssignmentExpirationDays,
		
		SelfUri: u.SelfUri,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Journeysegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
