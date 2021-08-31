package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditsearchresult
type Auditsearchresult struct { 
	// PageNumber - Which page was returned.
	PageNumber *int `json:"pageNumber,omitempty"`


	// PageSize - The number of results in a page.
	PageSize *int `json:"pageSize,omitempty"`


	// Total - The total number of results.
	Total *int `json:"total,omitempty"`


	// PageCount - The number of pages of results.
	PageCount *int `json:"pageCount,omitempty"`


	// FacetInfo
	FacetInfo *[]Facetinfo `json:"facetInfo,omitempty"`


	// AuditMessages
	AuditMessages *[]Auditmessage `json:"auditMessages,omitempty"`

}

func (o *Auditsearchresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditsearchresult
	
	return json.Marshal(&struct { 
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		FacetInfo *[]Facetinfo `json:"facetInfo,omitempty"`
		
		AuditMessages *[]Auditmessage `json:"auditMessages,omitempty"`
		*Alias
	}{ 
		PageNumber: o.PageNumber,
		
		PageSize: o.PageSize,
		
		Total: o.Total,
		
		PageCount: o.PageCount,
		
		FacetInfo: o.FacetInfo,
		
		AuditMessages: o.AuditMessages,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditsearchresult) UnmarshalJSON(b []byte) error {
	var AuditsearchresultMap map[string]interface{}
	err := json.Unmarshal(b, &AuditsearchresultMap)
	if err != nil {
		return err
	}
	
	if PageNumber, ok := AuditsearchresultMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := AuditsearchresultMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if Total, ok := AuditsearchresultMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if PageCount, ok := AuditsearchresultMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if FacetInfo, ok := AuditsearchresultMap["facetInfo"].([]interface{}); ok {
		FacetInfoString, _ := json.Marshal(FacetInfo)
		json.Unmarshal(FacetInfoString, &o.FacetInfo)
	}
	
	if AuditMessages, ok := AuditsearchresultMap["auditMessages"].([]interface{}); ok {
		AuditMessagesString, _ := json.Marshal(AuditMessages)
		json.Unmarshal(AuditMessagesString, &o.AuditMessages)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditsearchresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
