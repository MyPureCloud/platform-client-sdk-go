package platformclientv2
import (
	"encoding/json"
)

// Managementunitlisting
type Managementunitlisting struct { 
	// Entities
	Entities *[]Managementunit `json:"entities,omitempty"`


	// PageSize - Deprecated, paging is not supported
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - Deprecated, paging is not supported
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total - Deprecated, paging is not supported
	Total *int `json:"total,omitempty"`


	// FirstUri - Deprecated, paging is not supported
	FirstUri *string `json:"firstUri,omitempty"`


	// NextUri - Deprecated, paging is not supported
	NextUri *string `json:"nextUri,omitempty"`


	// PageCount - Deprecated, paging is not supported
	PageCount *int `json:"pageCount,omitempty"`


	// PreviousUri - Deprecated, paging is not supported
	PreviousUri *string `json:"previousUri,omitempty"`


	// LastUri - Deprecated, paging is not supported
	LastUri *string `json:"lastUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Managementunitlisting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
