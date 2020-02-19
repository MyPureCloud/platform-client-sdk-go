package platformclientv2
import (
	"time"
	"encoding/json"
)

// Presenceeventuserpresence
type Presenceeventuserpresence struct { 
	// Source
	Source *string `json:"source,omitempty"`


	// PresenceDefinition
	PresenceDefinition *Presenceeventorganizationpresence `json:"presenceDefinition,omitempty"`


	// Primary
	Primary *bool `json:"primary,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// ModifiedDate
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Presenceeventuserpresence) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
