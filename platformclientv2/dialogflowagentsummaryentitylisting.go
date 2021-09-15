package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialogflowagentsummaryentitylisting
type Dialogflowagentsummaryentitylisting struct { 
	// Entities
	Entities *[]Dialogflowagentsummary `json:"entities,omitempty"`


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


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// LastUri
	LastUri *string `json:"lastUri,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

func (o *Dialogflowagentsummaryentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialogflowagentsummaryentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Dialogflowagentsummary `json:"entities,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		FirstUri *string `json:"firstUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
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
		
		PreviousUri: o.PreviousUri,
		
		NextUri: o.NextUri,
		
		LastUri: o.LastUri,
		
		PageCount: o.PageCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialogflowagentsummaryentitylisting) UnmarshalJSON(b []byte) error {
	var DialogflowagentsummaryentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &DialogflowagentsummaryentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := DialogflowagentsummaryentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := DialogflowagentsummaryentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := DialogflowagentsummaryentitylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := DialogflowagentsummaryentitylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if FirstUri, ok := DialogflowagentsummaryentitylistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
	
	if SelfUri, ok := DialogflowagentsummaryentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if PreviousUri, ok := DialogflowagentsummaryentitylistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
	
	if NextUri, ok := DialogflowagentsummaryentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
	
	if LastUri, ok := DialogflowagentsummaryentitylistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
	
	if PageCount, ok := DialogflowagentsummaryentitylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialogflowagentsummaryentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
