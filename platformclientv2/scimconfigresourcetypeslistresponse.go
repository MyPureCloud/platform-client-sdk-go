package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Scimconfigresourcetypeslistresponse - Defines a response for a list of SCIM resource types.
type Scimconfigresourcetypeslistresponse struct { 
	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`


	// TotalResults - The total number of results.
	TotalResults *int `json:"totalResults,omitempty"`


	// StartIndex - The 1-based index of the first result returned by this request. Add this to \"itemsPerPage\" when requesting the next page of results.
	StartIndex *int `json:"startIndex,omitempty"`


	// ItemsPerPage - The number of resources returned per page.
	ItemsPerPage *int `json:"itemsPerPage,omitempty"`


	// Resources - The list of requested resources.
	Resources *[]Scimconfigresourcetype `json:"Resources,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimconfigresourcetypeslistresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
