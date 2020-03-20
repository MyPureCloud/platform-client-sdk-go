package platformclientv2
import (
	"encoding/json"
)

// Searchsort
type Searchsort struct { 
	// SortOrder - The sort order for results
	SortOrder *string `json:"sortOrder,omitempty"`


	// SortBy - The field in the resource that you want to sort the results by
	SortBy *string `json:"sortBy,omitempty"`

}

// String returns a JSON representation of the model
func (o *Searchsort) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
