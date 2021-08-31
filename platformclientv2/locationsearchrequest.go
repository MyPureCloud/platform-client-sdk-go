package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationsearchrequest
type Locationsearchrequest struct { 
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
	Query *[]Locationsearchcriteria `json:"query,omitempty"`

}

func (o *Locationsearchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Locationsearchrequest
	
	return json.Marshal(&struct { 
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Sort *[]Searchsort `json:"sort,omitempty"`
		
		Expand *[]string `json:"expand,omitempty"`
		
		Query *[]Locationsearchcriteria `json:"query,omitempty"`
		*Alias
	}{ 
		SortOrder: o.SortOrder,
		
		SortBy: o.SortBy,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Sort: o.Sort,
		
		Expand: o.Expand,
		
		Query: o.Query,
		Alias:    (*Alias)(o),
	})
}

func (o *Locationsearchrequest) UnmarshalJSON(b []byte) error {
	var LocationsearchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &LocationsearchrequestMap)
	if err != nil {
		return err
	}
	
	if SortOrder, ok := LocationsearchrequestMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
	
	if SortBy, ok := LocationsearchrequestMap["sortBy"].(string); ok {
		o.SortBy = &SortBy
	}
	
	if PageSize, ok := LocationsearchrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := LocationsearchrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Sort, ok := LocationsearchrequestMap["sort"].([]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	
	if Expand, ok := LocationsearchrequestMap["expand"].([]interface{}); ok {
		ExpandString, _ := json.Marshal(Expand)
		json.Unmarshal(ExpandString, &o.Expand)
	}
	
	if Query, ok := LocationsearchrequestMap["query"].([]interface{}); ok {
		QueryString, _ := json.Marshal(Query)
		json.Unmarshal(QueryString, &o.Query)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Locationsearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
