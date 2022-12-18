package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Datatablesdomainentitylisting
type Datatablesdomainentitylisting struct { 
	// Entities
	Entities *[]Datatable `json:"entities,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// FirstUri
	FirstUri *string `json:"firstUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// LastUri
	LastUri *string `json:"lastUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

func (o *Datatablesdomainentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Datatablesdomainentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Datatable `json:"entities,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		FirstUri *string `json:"firstUri,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		LastUri *string `json:"lastUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Total: o.Total,
		
		FirstUri: o.FirstUri,
		
		NextUri: o.NextUri,
		
		LastUri: o.LastUri,
		
		PreviousUri: o.PreviousUri,
		
		SelfUri: o.SelfUri,
		
		PageCount: o.PageCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Datatablesdomainentitylisting) UnmarshalJSON(b []byte) error {
	var DatatablesdomainentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &DatatablesdomainentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := DatatablesdomainentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := DatatablesdomainentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := DatatablesdomainentitylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := DatatablesdomainentitylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if FirstUri, ok := DatatablesdomainentitylistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
    
	if NextUri, ok := DatatablesdomainentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
    
	if LastUri, ok := DatatablesdomainentitylistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
    
	if PreviousUri, ok := DatatablesdomainentitylistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
    
	if SelfUri, ok := DatatablesdomainentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if PageCount, ok := DatatablesdomainentitylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Datatablesdomainentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
