package platformclientv2
import (
	"time"
	"encoding/json"
)

// Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata
type Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata struct { 
	// Version
	Version *int `json:"version,omitempty"`


	// ModifiedBy
	ModifiedBy *Wfmbushorttermforecastimportcompletetopicuserreference `json:"modifiedBy,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicwfmversionedentitymetadata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
