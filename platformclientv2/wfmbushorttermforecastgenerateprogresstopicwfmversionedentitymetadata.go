package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastgenerateprogresstopicwfmversionedentitymetadata
type Wfmbushorttermforecastgenerateprogresstopicwfmversionedentitymetadata struct { 
	// Version
	Version *int `json:"version,omitempty"`


	// ModifiedBy
	ModifiedBy *Wfmbushorttermforecastgenerateprogresstopicuserreference `json:"modifiedBy,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastgenerateprogresstopicwfmversionedentitymetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
