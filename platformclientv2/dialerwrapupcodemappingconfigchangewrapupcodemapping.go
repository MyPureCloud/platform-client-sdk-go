package platformclientv2
import (
	"time"
	"encoding/json"
)

// Dialerwrapupcodemappingconfigchangewrapupcodemapping
type Dialerwrapupcodemappingconfigchangewrapupcodemapping struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int32 `json:"version,omitempty"`


	// DefaultSet
	DefaultSet *[]string `json:"defaultSet,omitempty"`


	// Mapping
	Mapping *map[string][]string `json:"mapping,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerwrapupcodemappingconfigchangewrapupcodemapping) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
