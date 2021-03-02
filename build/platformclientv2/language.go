package platformclientv2
import (
	"time"
	"encoding/json"
)

// Language
type Language struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The language name.
	Name *string `json:"name,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Language) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
