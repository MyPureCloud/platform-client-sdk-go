package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentqueryrequest
type Contentqueryrequest struct { 
	// QueryPhrase
	QueryPhrase *string `json:"queryPhrase,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// FacetNameRequests
	FacetNameRequests *[]string `json:"facetNameRequests,omitempty"`


	// Sort
	Sort *[]Contentsortitem `json:"sort,omitempty"`


	// Filters
	Filters *[]Contentfacetfilteritem `json:"filters,omitempty"`


	// AttributeFilters
	AttributeFilters *[]Contentattributefilteritem `json:"attributeFilters,omitempty"`


	// IncludeShares
	IncludeShares *bool `json:"includeShares,omitempty"`

}

func (o *Contentqueryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentqueryrequest
	
	return json.Marshal(&struct { 
		QueryPhrase *string `json:"queryPhrase,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		FacetNameRequests *[]string `json:"facetNameRequests,omitempty"`
		
		Sort *[]Contentsortitem `json:"sort,omitempty"`
		
		Filters *[]Contentfacetfilteritem `json:"filters,omitempty"`
		
		AttributeFilters *[]Contentattributefilteritem `json:"attributeFilters,omitempty"`
		
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

func (o *Contentqueryrequest) UnmarshalJSON(b []byte) error {
	var ContentqueryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ContentqueryrequestMap)
	if err != nil {
		return err
	}
	
	if QueryPhrase, ok := ContentqueryrequestMap["queryPhrase"].(string); ok {
		o.QueryPhrase = &QueryPhrase
	}
	
	if PageNumber, ok := ContentqueryrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := ContentqueryrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if FacetNameRequests, ok := ContentqueryrequestMap["facetNameRequests"].([]interface{}); ok {
		FacetNameRequestsString, _ := json.Marshal(FacetNameRequests)
		json.Unmarshal(FacetNameRequestsString, &o.FacetNameRequests)
	}
	
	if Sort, ok := ContentqueryrequestMap["sort"].([]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	
	if Filters, ok := ContentqueryrequestMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	
	if AttributeFilters, ok := ContentqueryrequestMap["attributeFilters"].([]interface{}); ok {
		AttributeFiltersString, _ := json.Marshal(AttributeFilters)
		json.Unmarshal(AttributeFiltersString, &o.AttributeFilters)
	}
	
	if IncludeShares, ok := ContentqueryrequestMap["includeShares"].(bool); ok {
		o.IncludeShares = &IncludeShares
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
