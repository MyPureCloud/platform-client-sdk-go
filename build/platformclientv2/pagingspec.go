package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Pagingspec
type Pagingspec struct { 
	// PageSize - How many results per page
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - How many pages in
	PageNumber *int `json:"pageNumber,omitempty"`

}

// String returns a JSON representation of the model
func (o *Pagingspec) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
