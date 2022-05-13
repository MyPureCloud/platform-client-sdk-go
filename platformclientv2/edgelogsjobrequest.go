package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Edgelogsjobrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgelogsjobrequest
	
	return json.Marshal(&struct { 
		Path *string `json:"path,omitempty"`
		
		Query *string `json:"query,omitempty"`
		
		Recurse *bool `json:"recurse,omitempty"`
		*Alias
	}{ 
		Path: o.Path,
		
		Query: o.Query,
		
		Recurse: o.Recurse,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgelogsjobrequest) UnmarshalJSON(b []byte) error {
	var EdgelogsjobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &EdgelogsjobrequestMap)
	if err != nil {
		return err
	}
	
	if Path, ok := EdgelogsjobrequestMap["path"].(string); ok {
		o.Path = &Path
	}
    
	if Query, ok := EdgelogsjobrequestMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if Recurse, ok := EdgelogsjobrequestMap["recurse"].(bool); ok {
		o.Recurse = &Recurse
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgelogsjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
