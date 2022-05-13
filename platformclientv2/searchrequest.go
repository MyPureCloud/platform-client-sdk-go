package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Searchrequest
type Searchrequest struct { 
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


	// ReturnFields - A List of strings.  Possible values are any field in the resource you are searching on.  The other option is to use ALL_FIELDS, when this is provided all fields in the resource will be returned in the search results.
	ReturnFields *[]string `json:"returnFields,omitempty"`


	// Expand - Provides more details about a specified resource
	Expand *[]string `json:"expand,omitempty"`


	// Types - Resource domain type to search
	Types *[]string `json:"types,omitempty"`


	// Query - The search criteria
	Query *[]Searchcriteria `json:"query,omitempty"`


	// Aggregations - Aggregation criteria
	Aggregations *[]Searchaggregation `json:"aggregations,omitempty"`

}

func (o *Searchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Searchrequest
	
	return json.Marshal(&struct { 
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Sort *[]Searchsort `json:"sort,omitempty"`
		
		ReturnFields *[]string `json:"returnFields,omitempty"`
		
		Expand *[]string `json:"expand,omitempty"`
		
		Types *[]string `json:"types,omitempty"`
		
		Query *[]Searchcriteria `json:"query,omitempty"`
		
		Aggregations *[]Searchaggregation `json:"aggregations,omitempty"`
		*Alias
	}{ 
		SortOrder: o.SortOrder,
		
		SortBy: o.SortBy,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Sort: o.Sort,
		
		ReturnFields: o.ReturnFields,
		
		Expand: o.Expand,
		
		Types: o.Types,
		
		Query: o.Query,
		
		Aggregations: o.Aggregations,
		Alias:    (*Alias)(o),
	})
}

func (o *Searchrequest) UnmarshalJSON(b []byte) error {
	var SearchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SearchrequestMap)
	if err != nil {
		return err
	}
	
	if SortOrder, ok := SearchrequestMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
    
	if SortBy, ok := SearchrequestMap["sortBy"].(string); ok {
		o.SortBy = &SortBy
	}
    
	if PageSize, ok := SearchrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := SearchrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Sort, ok := SearchrequestMap["sort"].([]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	
	if ReturnFields, ok := SearchrequestMap["returnFields"].([]interface{}); ok {
		ReturnFieldsString, _ := json.Marshal(ReturnFields)
		json.Unmarshal(ReturnFieldsString, &o.ReturnFields)
	}
	
	if Expand, ok := SearchrequestMap["expand"].([]interface{}); ok {
		ExpandString, _ := json.Marshal(Expand)
		json.Unmarshal(ExpandString, &o.Expand)
	}
	
	if Types, ok := SearchrequestMap["types"].([]interface{}); ok {
		TypesString, _ := json.Marshal(Types)
		json.Unmarshal(TypesString, &o.Types)
	}
	
	if Query, ok := SearchrequestMap["query"].([]interface{}); ok {
		QueryString, _ := json.Marshal(Query)
		json.Unmarshal(QueryString, &o.Query)
	}
	
	if Aggregations, ok := SearchrequestMap["aggregations"].([]interface{}); ok {
		AggregationsString, _ := json.Marshal(Aggregations)
		json.Unmarshal(AggregationsString, &o.Aggregations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Searchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
