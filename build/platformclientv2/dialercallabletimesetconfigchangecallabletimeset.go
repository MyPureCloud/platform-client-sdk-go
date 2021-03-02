package platformclientv2
import (
	"time"
	"encoding/json"
)

// Dialercallabletimesetconfigchangecallabletimeset
type Dialercallabletimesetconfigchangecallabletimeset struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// CallableTimes
	CallableTimes *[]Dialercallabletimesetconfigchangecallabletime `json:"callableTimes,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercallabletimesetconfigchangecallabletimeset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
