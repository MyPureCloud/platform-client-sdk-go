package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Gkndocumentationsearchrequest
type Gkndocumentationsearchrequest struct { 
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


	// Query
	Query *[]Gkndocumentationsearchcriteria `json:"query,omitempty"`

}

func (o *Gkndocumentationsearchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Gkndocumentationsearchrequest
	
	return json.Marshal(&struct { 
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Sort *[]Searchsort `json:"sort,omitempty"`
		
		Query *[]Gkndocumentationsearchcriteria `json:"query,omitempty"`
		*Alias
	}{ 
		SortOrder: o.SortOrder,
		
		SortBy: o.SortBy,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Sort: o.Sort,
		
		Query: o.Query,
		Alias:    (*Alias)(o),
	})
}

func (o *Gkndocumentationsearchrequest) UnmarshalJSON(b []byte) error {
	var GkndocumentationsearchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &GkndocumentationsearchrequestMap)
	if err != nil {
		return err
	}
	
	if SortOrder, ok := GkndocumentationsearchrequestMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
	
	if SortBy, ok := GkndocumentationsearchrequestMap["sortBy"].(string); ok {
		o.SortBy = &SortBy
	}
	
	if PageSize, ok := GkndocumentationsearchrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := GkndocumentationsearchrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Sort, ok := GkndocumentationsearchrequestMap["sort"].([]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	
	if Query, ok := GkndocumentationsearchrequestMap["query"].([]interface{}); ok {
		QueryString, _ := json.Marshal(Query)
		json.Unmarshal(QueryString, &o.Query)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Gkndocumentationsearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
