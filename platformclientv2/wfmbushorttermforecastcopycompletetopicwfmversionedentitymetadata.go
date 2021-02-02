package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmbushorttermforecastcopycompletetopicwfmversionedentitymetadata
type Wfmbushorttermforecastcopycompletetopicwfmversionedentitymetadata struct { 
	// Version
	Version *int `json:"version,omitempty"`


	// ModifiedBy
	ModifiedBy *Wfmbushorttermforecastcopycompletetopicuserreference `json:"modifiedBy,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastcopycompletetopicwfmversionedentitymetadata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
