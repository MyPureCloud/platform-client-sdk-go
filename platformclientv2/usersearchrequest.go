package platformclientv2
import (
	"encoding/json"
)

// Usersearchrequest
type Usersearchrequest struct { 
	// SortOrder - The sort order for results
	SortOrder *string `json:"sortOrder,omitempty"`


	// SortBy - The field in the resource that you want to sort the results by
	SortBy *string `json:"sortBy,omitempty"`


	// PageSize - The number of results per page
	PageSize *int32 `json:"pageSize,omitempty"`


	// PageNumber - The page of resources you want to retrieve
	PageNumber *int32 `json:"pageNumber,omitempty"`


	// Sort - Multi-value sort order, list of multiple sort values
	Sort *[]Searchsort `json:"sort,omitempty"`


	// Expand - Provides more details about a specified resource
	Expand *[]string `json:"expand,omitempty"`


	// Query
	Query *[]Usersearchcriteria `json:"query,omitempty"`


	// IntegrationPresenceSource - Gets an integration presence for users instead of their defaults. This parameter will only be used when presence is provided as an \"expand\". When using this parameter the maximum number of users that can be returned is 10.
	IntegrationPresenceSource *string `json:"integrationPresenceSource,omitempty"`

}

// String returns a JSON representation of the model
func (o *Usersearchrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
