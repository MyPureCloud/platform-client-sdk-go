package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistentitylisting
type Contactlistentitylisting struct { 
	// Entities
	Entities *[]Contactlist `json:"entities,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// FirstUri
	FirstUri *string `json:"firstUri,omitempty"`


	// LastUri
	LastUri *string `json:"lastUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

func (o *Contactlistentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Contactlist `json:"entities,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		FirstUri *string `json:"firstUri,omitempty"`
		
		LastUri *string `json:"lastUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Total: o.Total,
		
		FirstUri: o.FirstUri,
		
		LastUri: o.LastUri,
		
		SelfUri: o.SelfUri,
		
		NextUri: o.NextUri,
		
		PreviousUri: o.PreviousUri,
		
		PageCount: o.PageCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactlistentitylisting) UnmarshalJSON(b []byte) error {
	var ContactlistentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ContactlistentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := ContactlistentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := ContactlistentitylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := ContactlistentitylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if FirstUri, ok := ContactlistentitylistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
    
	if LastUri, ok := ContactlistentitylistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
    
	if SelfUri, ok := ContactlistentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if NextUri, ok := ContactlistentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
    
	if PreviousUri, ok := ContactlistentitylistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
    
	if PageCount, ok := ContactlistentitylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
