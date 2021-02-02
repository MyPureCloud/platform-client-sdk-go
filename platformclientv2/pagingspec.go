package platformclientv2
import (
	"encoding/json"
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
	return string(j)
}
