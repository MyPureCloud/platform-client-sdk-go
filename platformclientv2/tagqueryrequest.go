package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Tagqueryrequest
type Tagqueryrequest struct { 
	// Query
	Query *string `json:"query,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`

}

// String returns a JSON representation of the model
func (o *Tagqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
