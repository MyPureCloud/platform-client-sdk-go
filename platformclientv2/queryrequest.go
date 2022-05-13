package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queryrequest
type Queryrequest struct { 
	// QueryPhrase
	QueryPhrase *string `json:"queryPhrase,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// FacetNameRequests
	FacetNameRequests *[]string `json:"facetNameRequests,omitempty"`


	// Sort
	Sort *[]Sortitem `json:"sort,omitempty"`


	// Filters
	Filters *[]Contentfilteritem `json:"filters,omitempty"`


	// AttributeFilters
	AttributeFilters *[]Attributefilteritem `json:"attributeFilters,omitempty"`


	// IncludeShares
	IncludeShares *bool `json:"includeShares,omitempty"`

}

func (o *Queryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queryrequest
	
	return json.Marshal(&struct { 
		QueryPhrase *string `json:"queryPhrase,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		FacetNameRequests *[]string `json:"facetNameRequests,omitempty"`
		
		Sort *[]Sortitem `json:"sort,omitempty"`
		
		Filters *[]Contentfilteritem `json:"filters,omitempty"`
		
		AttributeFilters *[]Attributefilteritem `json:"attributeFilters,omitempty"`
		
		IncludeShares *bool `json:"includeShares,omitempty"`
		*Alias
	}{ 
		QueryPhrase: o.QueryPhrase,
		
		PageNumber: o.PageNumber,
		
		PageSize: o.PageSize,
		
		FacetNameRequests: o.FacetNameRequests,
		
		Sort: o.Sort,
		
		Filters: o.Filters,
		
		AttributeFilters: o.AttributeFilters,
		
		IncludeShares: o.IncludeShares,
		Alias:    (*Alias)(o),
	})
}

func (o *Queryrequest) UnmarshalJSON(b []byte) error {
	var QueryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &QueryrequestMap)
	if err != nil {
		return err
	}
	
	if QueryPhrase, ok := QueryrequestMap["queryPhrase"].(string); ok {
		o.QueryPhrase = &QueryPhrase
	}
    
	if PageNumber, ok := QueryrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := QueryrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if FacetNameRequests, ok := QueryrequestMap["facetNameRequests"].([]interface{}); ok {
		FacetNameRequestsString, _ := json.Marshal(FacetNameRequests)
		json.Unmarshal(FacetNameRequestsString, &o.FacetNameRequests)
	}
	
	if Sort, ok := QueryrequestMap["sort"].([]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	
	if Filters, ok := QueryrequestMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	
	if AttributeFilters, ok := QueryrequestMap["attributeFilters"].([]interface{}); ok {
		AttributeFiltersString, _ := json.Marshal(AttributeFilters)
		json.Unmarshal(AttributeFiltersString, &o.AttributeFilters)
	}
	
	if IncludeShares, ok := QueryrequestMap["includeShares"].(bool); ok {
		o.IncludeShares = &IncludeShares
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
