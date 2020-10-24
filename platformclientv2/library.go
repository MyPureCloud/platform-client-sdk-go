package platformclientv2
import (
	"time"
	"encoding/json"
)

// Library
type Library struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The library name.
	Name *string `json:"name,omitempty"`


	// Version - Current version for this resource.
	Version *int32 `json:"version,omitempty"`


	// CreatedBy - User that created the library.
	CreatedBy *User `json:"createdBy,omitempty"`


	// DateCreated - The date and time the response was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ResponseType - This value is deprecated. Responses representing message templates may be added to any library.
	ResponseType *string `json:"responseType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Library) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
