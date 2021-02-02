package platformclientv2
import (
	"encoding/json"
)

// Scimconfigresourcetypeslistresponse - Defines a response for a list of SCIM resource types.
type Scimconfigresourcetypeslistresponse struct { 
	// TotalResults - The total number of results.
	TotalResults *int `json:"totalResults,omitempty"`


	// StartIndex - The 1-based index of the first result returned by this request. Add this to \"itemsPerPage\" when requesting the next page of results.
	StartIndex *int `json:"startIndex,omitempty"`


	// ItemsPerPage - The number of resources returned per page.
	ItemsPerPage *int `json:"itemsPerPage,omitempty"`


	// Resources - The list of requested resources.
	Resources *[]Scimconfigresourcetype `json:"Resources,omitempty"`


	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimconfigresourcetypeslistresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
