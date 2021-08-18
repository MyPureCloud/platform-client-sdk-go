package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Usersearchrequest
type Usersearchrequest struct { 
	// SortOrder - The sort order for results
	SortOrder *string `json:"sortOrder,omitempty"`


	// SortBy - The field in the resource that you want to sort the results by
	SortBy *string `json:"sortBy,omitempty"`


	// PageSize - The number of results per page
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - The page of resources you want to retrieve
	PageNumber *int `json:"pageNumber,omitempty"`


	// Sort - Multi-value sort order, list of multiple sort values
	Sort *[]Searchsort `json:"sort,omitempty"`


	// Expand - Provides more details about a specified resource
	Expand *[]string `json:"expand,omitempty"`


	// Query
	Query *[]Usersearchcriteria `json:"query,omitempty"`


	// IntegrationPresenceSource - Gets an integration presence for users instead of their defaults. This parameter will only be used when presence is provided as an \"expand\". When using this parameter the maximum number of users that can be returned is 10.
	IntegrationPresenceSource *string `json:"integrationPresenceSource,omitempty"`


	// EnforcePermissions - This property only applies to api/v2/user/search; when set to true add additional search criteria to filter users by: directory:user:view
	EnforcePermissions *bool `json:"enforcePermissions,omitempty"`

}

func (u *Usersearchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Usersearchrequest

	

	return json.Marshal(&struct { 
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Sort *[]Searchsort `json:"sort,omitempty"`
		
		Expand *[]string `json:"expand,omitempty"`
		
		Query *[]Usersearchcriteria `json:"query,omitempty"`
		
		IntegrationPresenceSource *string `json:"integrationPresenceSource,omitempty"`
		
		EnforcePermissions *bool `json:"enforcePermissions,omitempty"`
		*Alias
	}{ 
		SortOrder: u.SortOrder,
		
		SortBy: u.SortBy,
		
		PageSize: u.PageSize,
		
		PageNumber: u.PageNumber,
		
		Sort: u.Sort,
		
		Expand: u.Expand,
		
		Query: u.Query,
		
		IntegrationPresenceSource: u.IntegrationPresenceSource,
		
		EnforcePermissions: u.EnforcePermissions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Usersearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
