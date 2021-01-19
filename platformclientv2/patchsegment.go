package platformclientv2
import (
	"time"
	"encoding/json"
)

// Patchsegment
type Patchsegment struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// IsActive - Whether or not the segment is active.
	IsActive *bool `json:"isActive,omitempty"`


	// DisplayName - The display name of the segment.
	DisplayName *string `json:"displayName,omitempty"`


	// Version - The version of the segment.
	Version *int32 `json:"version,omitempty"`


	// Description - A description of the segment.
	Description *string `json:"description,omitempty"`


	// Color - The hexadecimal color value of the segment.
	Color *string `json:"color,omitempty"`


	// ShouldDisplayToAgent - Whether or not the segment should be displayed to agent/supervisor users.
	ShouldDisplayToAgent *bool `json:"shouldDisplayToAgent,omitempty"`


	// Context - The context of the segment.
	Context *Context `json:"context,omitempty"`


	// Journey - The pattern of rules defining the segment.
	Journey *Journey `json:"journey,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// CreatedDate - Timestamp indicating when the segment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Timestamp indicating when the the segment was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchsegment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
