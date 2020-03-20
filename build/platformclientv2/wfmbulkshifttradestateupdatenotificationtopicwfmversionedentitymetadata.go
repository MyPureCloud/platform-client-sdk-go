package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmbulkshifttradestateupdatenotificationtopicwfmversionedentitymetadata
type Wfmbulkshifttradestateupdatenotificationtopicwfmversionedentitymetadata struct { 
	// Version
	Version *int32 `json:"version,omitempty"`


	// ModifiedBy
	ModifiedBy *Wfmbulkshifttradestateupdatenotificationtopicuserreference `json:"modifiedBy,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopicwfmversionedentitymetadata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
