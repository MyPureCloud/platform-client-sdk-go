package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responseassetsearchrequest
type Responseassetsearchrequest struct { 
	// PageSize - The number of results per page. Default: 25, Maximum: 100.
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - The page of resources you want to retrieve
	PageNumber *int `json:"pageNumber,omitempty"`


	// SortOrder - The sort order for results
	SortOrder *string `json:"sortOrder,omitempty"`


	// SortBy - The field in the resource that you want to sort the results by
	SortBy *string `json:"sortBy,omitempty"`


	// Query - Filter the query results.
	Query *[]Responseassetfilter `json:"query,omitempty"`

}

func (o *Responseassetsearchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responseassetsearchrequest
	
	return json.Marshal(&struct { 
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		
		Query *[]Responseassetfilter `json:"query,omitempty"`
		*Alias
	}{ 
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		SortOrder: o.SortOrder,
		
		SortBy: o.SortBy,
		
		Query: o.Query,
		Alias:    (*Alias)(o),
	})
}

func (o *Responseassetsearchrequest) UnmarshalJSON(b []byte) error {
	var ResponseassetsearchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ResponseassetsearchrequestMap)
	if err != nil {
		return err
	}
	
	if PageSize, ok := ResponseassetsearchrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := ResponseassetsearchrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if SortOrder, ok := ResponseassetsearchrequestMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
    
	if SortBy, ok := ResponseassetsearchrequestMap["sortBy"].(string); ok {
		o.SortBy = &SortBy
	}
    
	if Query, ok := ResponseassetsearchrequestMap["query"].([]interface{}); ok {
		QueryString, _ := json.Marshal(Query)
		json.Unmarshal(QueryString, &o.Query)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Responseassetsearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
