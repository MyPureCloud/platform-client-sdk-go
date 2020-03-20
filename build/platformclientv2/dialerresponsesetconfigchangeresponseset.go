package platformclientv2
import (
	"time"
	"encoding/json"
)

// Dialerresponsesetconfigchangeresponseset
type Dialerresponsesetconfigchangeresponseset struct { 
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


	// Responses
	Responses *map[string]Dialerresponsesetconfigchangereaction `json:"responses,omitempty"`


	// BeepDetectionEnabled
	BeepDetectionEnabled *bool `json:"beepDetectionEnabled,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerresponsesetconfigchangeresponseset) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
