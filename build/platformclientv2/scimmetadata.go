package platformclientv2
import (
	"time"
	"encoding/json"
)

// Scimmetadata - Defines the SCIM metadata.
type Scimmetadata struct { 
	// ResourceType - The type of SCIM resource.
	ResourceType *string `json:"resourceType,omitempty"`


	// LastModified - The last time that the resource was modified. Date time is represented as an ISO-8601 string, for example, yyyy-MM-ddTHH:mm:ss.SSSZ.
	LastModified *time.Time `json:"lastModified,omitempty"`


	// Location - The URI of the resource.
	Location *string `json:"location,omitempty"`


	// Version - The version of the resource. Matches the ETag HTTP response header.
	Version *string `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimmetadata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
