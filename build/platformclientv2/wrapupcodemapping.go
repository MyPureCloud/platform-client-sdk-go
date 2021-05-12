package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Wrapupcodemapping
type Wrapupcodemapping struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`


	// DefaultSet - The default set of wrap-up flags. These will be used if there is no entry for a given wrap-up code in the mapping.
	DefaultSet *[]string `json:"defaultSet,omitempty"`


	// Mapping - A map from wrap-up code identifiers to a set of wrap-up flags.
	Mapping *map[string][]string `json:"mapping,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wrapupcodemapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
