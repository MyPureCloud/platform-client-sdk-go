package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Managementunitlisting
type Managementunitlisting struct { 
	// Entities
	Entities *[]Managementunit `json:"entities,omitempty"`


	// PageSize - Deprecated, paging is not supported
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - Deprecated, paging is not supported
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total - Deprecated, paging is not supported
	Total *int `json:"total,omitempty"`


	// FirstUri - Deprecated, paging is not supported
	FirstUri *string `json:"firstUri,omitempty"`


	// LastUri - Deprecated, paging is not supported
	LastUri *string `json:"lastUri,omitempty"`


	// NextUri - Deprecated, paging is not supported
	NextUri *string `json:"nextUri,omitempty"`


	// PageCount - Deprecated, paging is not supported
	PageCount *int `json:"pageCount,omitempty"`


	// PreviousUri - Deprecated, paging is not supported
	PreviousUri *string `json:"previousUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Managementunitlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Managementunitlisting
	
	return json.Marshal(&struct { 
		Entities *[]Managementunit `json:"entities,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		FirstUri *string `json:"firstUri,omitempty"`
		
		LastUri *string `json:"lastUri,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Total: o.Total,
		
		FirstUri: o.FirstUri,
		
		LastUri: o.LastUri,
		
		NextUri: o.NextUri,
		
		PageCount: o.PageCount,
		
		PreviousUri: o.PreviousUri,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Managementunitlisting) UnmarshalJSON(b []byte) error {
	var ManagementunitlistingMap map[string]interface{}
	err := json.Unmarshal(b, &ManagementunitlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ManagementunitlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := ManagementunitlistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := ManagementunitlistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := ManagementunitlistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if FirstUri, ok := ManagementunitlistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
    
	if LastUri, ok := ManagementunitlistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
    
	if NextUri, ok := ManagementunitlistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
    
	if PageCount, ok := ManagementunitlistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if PreviousUri, ok := ManagementunitlistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
    
	if SelfUri, ok := ManagementunitlistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Managementunitlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
