package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmbuscheduletopicwfmversionedentitymetadata
type Wfmbuscheduletopicwfmversionedentitymetadata struct { 
	// Version
	Version *int `json:"version,omitempty"`


	// ModifiedBy
	ModifiedBy *Wfmbuscheduletopicuserreference `json:"modifiedBy,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicwfmversionedentitymetadata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
