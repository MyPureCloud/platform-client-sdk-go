package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistfilterentitylisting
type Contactlistfilterentitylisting struct { 
	// Entities
	Entities *[]Contactlistfilter `json:"entities,omitempty"`


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

func (o *Contactlistfilterentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistfilterentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Contactlistfilter `json:"entities,omitempty"`
		
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

func (o *Contactlistfilterentitylisting) UnmarshalJSON(b []byte) error {
	var ContactlistfilterentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistfilterentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ContactlistfilterentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := ContactlistfilterentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := ContactlistfilterentitylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := ContactlistfilterentitylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if FirstUri, ok := ContactlistfilterentitylistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
	
	if SelfUri, ok := ContactlistfilterentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if NextUri, ok := ContactlistfilterentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
	
	if PreviousUri, ok := ContactlistfilterentitylistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
	
	if LastUri, ok := ContactlistfilterentitylistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
	
	if PageCount, ok := ContactlistfilterentitylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistfilterentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
