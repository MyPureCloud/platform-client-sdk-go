package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmversionedentitymetadata - Metadata to associate with a given entity
type Wfmversionedentitymetadata struct { 
	// Version - The version of the associated entity.  Used to prevent conflicts on concurrent edits
	Version *int `json:"version,omitempty"`


	// ModifiedBy - The user who last modified the associated entity
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`


	// DateModified - The date the associated entity was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmversionedentitymetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
