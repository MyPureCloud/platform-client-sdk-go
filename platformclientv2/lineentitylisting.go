package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lineentitylisting
type Lineentitylisting struct { 
	// Entities
	Entities *[]Line `json:"entities,omitempty"`


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

func (o *Lineentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lineentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Line `json:"entities,omitempty"`
		
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

func (o *Lineentitylisting) UnmarshalJSON(b []byte) error {
	var LineentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &LineentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := LineentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := LineentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := LineentitylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := LineentitylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if FirstUri, ok := LineentitylistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
	
	if SelfUri, ok := LineentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if NextUri, ok := LineentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
	
	if PreviousUri, ok := LineentitylistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
	
	if LastUri, ok := LineentitylistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
	
	if PageCount, ok := LineentitylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Lineentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
