package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationssearchresponse
type Locationssearchresponse struct { 
	// Total - The total number of results found
	Total *int `json:"total,omitempty"`


	// PageCount - The total number of pages
	PageCount *int `json:"pageCount,omitempty"`


	// PageSize - The current page size
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - The current page number
	PageNumber *int `json:"pageNumber,omitempty"`


	// PreviousPage - Q64 value for the previous page of results
	PreviousPage *string `json:"previousPage,omitempty"`


	// CurrentPage - Q64 value for the current page of results
	CurrentPage *string `json:"currentPage,omitempty"`


	// NextPage - Q64 value for the next page of results
	NextPage *string `json:"nextPage,omitempty"`


	// Types - Resource types the search was performed against
	Types *[]string `json:"types,omitempty"`


	// Results - Search results
	Results *[]Locationdefinition `json:"results,omitempty"`

}

func (o *Locationssearchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Locationssearchresponse
	
	return json.Marshal(&struct { 
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PreviousPage *string `json:"previousPage,omitempty"`
		
		CurrentPage *string `json:"currentPage,omitempty"`
		
		NextPage *string `json:"nextPage,omitempty"`
		
		Types *[]string `json:"types,omitempty"`
		
		Results *[]Locationdefinition `json:"results,omitempty"`
		*Alias
	}{ 
		Total: o.Total,
		
		PageCount: o.PageCount,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		PreviousPage: o.PreviousPage,
		
		CurrentPage: o.CurrentPage,
		
		NextPage: o.NextPage,
		
		Types: o.Types,
		
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Locationssearchresponse) UnmarshalJSON(b []byte) error {
	var LocationssearchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &LocationssearchresponseMap)
	if err != nil {
		return err
	}
	
	if Total, ok := LocationssearchresponseMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if PageCount, ok := LocationssearchresponseMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if PageSize, ok := LocationssearchresponseMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := LocationssearchresponseMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PreviousPage, ok := LocationssearchresponseMap["previousPage"].(string); ok {
		o.PreviousPage = &PreviousPage
	}
	
	if CurrentPage, ok := LocationssearchresponseMap["currentPage"].(string); ok {
		o.CurrentPage = &CurrentPage
	}
	
	if NextPage, ok := LocationssearchresponseMap["nextPage"].(string); ok {
		o.NextPage = &NextPage
	}
	
	if Types, ok := LocationssearchresponseMap["types"].([]interface{}); ok {
		TypesString, _ := json.Marshal(Types)
		json.Unmarshal(TypesString, &o.Types)
	}
	
	if Results, ok := LocationssearchresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Locationssearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
