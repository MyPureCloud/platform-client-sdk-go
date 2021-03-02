package platformclientv2
import (
	"encoding/json"
)

// Scimuserlistresponse - Defines a response for a list of SCIM users.
type Scimuserlistresponse struct { 
	// TotalResults - The total number of results.
	TotalResults *int `json:"totalResults,omitempty"`


	// StartIndex - The 1-based index of the first result returned by this request. Add this to \"itemsPerPage\" when requesting the next page of results.
	StartIndex *int `json:"startIndex,omitempty"`


	// ItemsPerPage - The number of resources returned per page.
	ItemsPerPage *int `json:"itemsPerPage,omitempty"`


	// Resources - The list of requested resources. If \"count\" is 0, then the list will be empty.
	Resources *[]Scimv2user `json:"Resources,omitempty"`


	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimuserlistresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
