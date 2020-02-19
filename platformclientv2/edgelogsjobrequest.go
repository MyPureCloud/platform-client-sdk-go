package platformclientv2
import (
	"encoding/json"
)

// Edgelogsjobrequest
type Edgelogsjobrequest struct { 
	// Path - A relative directory to the root Edge log folder to query from.
	Path *string `json:"path,omitempty"`


	// Query - The pattern to use when searching for logs, which may include the wildcards {*, ?}.  Multiple search patterns may be combined using a pipe '|' as a delimiter.
	Query *string `json:"query,omitempty"`


	// Recurse - Boolean whether or not to recurse into directories.
	Recurse *bool `json:"recurse,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgelogsjobrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
