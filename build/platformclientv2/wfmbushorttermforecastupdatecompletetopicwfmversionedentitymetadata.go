package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadata
type Wfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadata struct { 
	// Version
	Version *int `json:"version,omitempty"`


	// ModifiedBy
	ModifiedBy *Wfmbushorttermforecastupdatecompletetopicuserreference `json:"modifiedBy,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
