package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actiontemplatelisting
type Actiontemplatelisting struct { 
	// Entities
	Entities *[]Actiontemplate `json:"entities,omitempty"`


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

func (o *Actiontemplatelisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actiontemplatelisting
	
	return json.Marshal(&struct { 
		Entities *[]Actiontemplate `json:"entities,omitempty"`
		
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

func (o *Actiontemplatelisting) UnmarshalJSON(b []byte) error {
	var ActiontemplatelistingMap map[string]interface{}
	err := json.Unmarshal(b, &ActiontemplatelistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ActiontemplatelistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := ActiontemplatelistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := ActiontemplatelistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := ActiontemplatelistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if FirstUri, ok := ActiontemplatelistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
    
	if LastUri, ok := ActiontemplatelistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
    
	if SelfUri, ok := ActiontemplatelistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if NextUri, ok := ActiontemplatelistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
    
	if PreviousUri, ok := ActiontemplatelistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
    
	if PageCount, ok := ActiontemplatelistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actiontemplatelisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
