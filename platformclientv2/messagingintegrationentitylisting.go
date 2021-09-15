package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingintegrationentitylisting
type Messagingintegrationentitylisting struct { 
	// Entities
	Entities *[]Messagingintegration `json:"entities,omitempty"`


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

func (o *Messagingintegrationentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagingintegrationentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Messagingintegration `json:"entities,omitempty"`
		
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

func (o *Messagingintegrationentitylisting) UnmarshalJSON(b []byte) error {
	var MessagingintegrationentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingintegrationentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := MessagingintegrationentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := MessagingintegrationentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := MessagingintegrationentitylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := MessagingintegrationentitylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if FirstUri, ok := MessagingintegrationentitylistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
	
	if SelfUri, ok := MessagingintegrationentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if PreviousUri, ok := MessagingintegrationentitylistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
	
	if NextUri, ok := MessagingintegrationentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
	
	if LastUri, ok := MessagingintegrationentitylistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
	
	if PageCount, ok := MessagingintegrationentitylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingintegrationentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
