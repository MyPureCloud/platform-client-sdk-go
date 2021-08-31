package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Datatablerowentitylisting
type Datatablerowentitylisting struct { 
	// Entities
	Entities *[]map[string]interface{} `json:"entities,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// FirstUri
	FirstUri *string `json:"firstUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// LastUri
	LastUri *string `json:"lastUri,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

func (o *Datatablerowentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Datatablerowentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]map[string]interface{} `json:"entities,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		FirstUri *string `json:"firstUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		
		LastUri *string `json:"lastUri,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Total: o.Total,
		
		FirstUri: o.FirstUri,
		
		SelfUri: o.SelfUri,
		
		NextUri: o.NextUri,
		
		PreviousUri: o.PreviousUri,
		
		LastUri: o.LastUri,
		
		PageCount: o.PageCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Datatablerowentitylisting) UnmarshalJSON(b []byte) error {
	var DatatablerowentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &DatatablerowentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := DatatablerowentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := DatatablerowentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := DatatablerowentitylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := DatatablerowentitylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if FirstUri, ok := DatatablerowentitylistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
	
	if SelfUri, ok := DatatablerowentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if NextUri, ok := DatatablerowentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
	
	if PreviousUri, ok := DatatablerowentitylistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
	
	if LastUri, ok := DatatablerowentitylistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
	
	if PageCount, ok := DatatablerowentitylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Datatablerowentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
