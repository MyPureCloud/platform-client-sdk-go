package platformclientv2
import (
	"encoding/json"
)

// Scimv2schemalistresponse - SCIM Resource Types list response
type Scimv2schemalistresponse struct { 
	// TotalResults - The total number of results.
	TotalResults *int64 `json:"totalResults,omitempty"`


	// StartIndex - The 1-based index of the first result returned by this request. Add this to \"itemsPerPage\" when requesting the next page of results.
	StartIndex *int64 `json:"startIndex,omitempty"`


	// ItemsPerPage - The number of resources returned per page.
	ItemsPerPage *int64 `json:"itemsPerPage,omitempty"`


	// Resources - Resources
	Resources *[]Scimv2schemadefinition `json:"Resources,omitempty"`


	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimv2schemalistresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
