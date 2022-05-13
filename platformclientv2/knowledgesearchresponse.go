package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesearchresponse
type Knowledgesearchresponse struct { 
	// SearchId - Search Id
	SearchId *string `json:"searchId,omitempty"`


	// Total - Total number of records returned
	Total *int `json:"total,omitempty"`


	// PageCount - Number of pages returned in the result calculated according to the pageSize and the total
	PageCount *int `json:"pageCount,omitempty"`


	// PageSize - Number of records according to the page size
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - Current page number for the returned records
	PageNumber *int `json:"pageNumber,omitempty"`


	// Results - Results associated to the search response
	Results *[]Knowledgesearchdocument `json:"results,omitempty"`

}

func (o *Knowledgesearchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgesearchresponse
	
	return json.Marshal(&struct { 
		SearchId *string `json:"searchId,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Results *[]Knowledgesearchdocument `json:"results,omitempty"`
		*Alias
	}{ 
		SearchId: o.SearchId,
		
		Total: o.Total,
		
		PageCount: o.PageCount,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgesearchresponse) UnmarshalJSON(b []byte) error {
	var KnowledgesearchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgesearchresponseMap)
	if err != nil {
		return err
	}
	
	if SearchId, ok := KnowledgesearchresponseMap["searchId"].(string); ok {
		o.SearchId = &SearchId
	}
    
	if Total, ok := KnowledgesearchresponseMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if PageCount, ok := KnowledgesearchresponseMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if PageSize, ok := KnowledgesearchresponseMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := KnowledgesearchresponseMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Results, ok := KnowledgesearchresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgesearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
