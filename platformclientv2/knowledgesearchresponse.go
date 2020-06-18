package platformclientv2
import (
	"encoding/json"
)

// Knowledgesearchresponse
type Knowledgesearchresponse struct { 
	// SearchId - Search Id
	SearchId *string `json:"searchId,omitempty"`


	// Total - Total number of records returned
	Total *int32 `json:"total,omitempty"`


	// PageCount - Number of pages returned in the result calculated according to the pageSize and the total
	PageCount *int32 `json:"pageCount,omitempty"`


	// PageSize - Number of records according to the page size
	PageSize *int32 `json:"pageSize,omitempty"`


	// PageNumber - Current page number for the returned records
	PageNumber *int32 `json:"pageNumber,omitempty"`


	// Results - Results associated to the search response
	Results *[]Knowledgesearchdocument `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Knowledgesearchresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
