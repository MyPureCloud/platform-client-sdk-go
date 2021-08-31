package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptsearchrequest
type Transcriptsearchrequest struct { 
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


	// ReturnFields
	ReturnFields *[]string `json:"returnFields,omitempty"`


	// Types - Resource domain type to search
	Types *[]string `json:"types,omitempty"`


	// Query - The search criteria
	Query *[]Transcriptsearchcriteria `json:"query,omitempty"`

}

func (o *Transcriptsearchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptsearchrequest
	
	return json.Marshal(&struct { 
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Sort *[]Searchsort `json:"sort,omitempty"`
		
		ReturnFields *[]string `json:"returnFields,omitempty"`
		
		Types *[]string `json:"types,omitempty"`
		
		Query *[]Transcriptsearchcriteria `json:"query,omitempty"`
		*Alias
	}{ 
		SortOrder: o.SortOrder,
		
		SortBy: o.SortBy,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Sort: o.Sort,
		
		ReturnFields: o.ReturnFields,
		
		Types: o.Types,
		
		Query: o.Query,
		Alias:    (*Alias)(o),
	})
}

func (o *Transcriptsearchrequest) UnmarshalJSON(b []byte) error {
	var TranscriptsearchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptsearchrequestMap)
	if err != nil {
		return err
	}
	
	if SortOrder, ok := TranscriptsearchrequestMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
	
	if SortBy, ok := TranscriptsearchrequestMap["sortBy"].(string); ok {
		o.SortBy = &SortBy
	}
	
	if PageSize, ok := TranscriptsearchrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := TranscriptsearchrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Sort, ok := TranscriptsearchrequestMap["sort"].([]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	
	if ReturnFields, ok := TranscriptsearchrequestMap["returnFields"].([]interface{}); ok {
		ReturnFieldsString, _ := json.Marshal(ReturnFields)
		json.Unmarshal(ReturnFieldsString, &o.ReturnFields)
	}
	
	if Types, ok := TranscriptsearchrequestMap["types"].([]interface{}); ok {
		TypesString, _ := json.Marshal(Types)
		json.Unmarshal(TypesString, &o.Types)
	}
	
	if Query, ok := TranscriptsearchrequestMap["query"].([]interface{}); ok {
		QueryString, _ := json.Marshal(Query)
		json.Unmarshal(QueryString, &o.Query)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptsearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
