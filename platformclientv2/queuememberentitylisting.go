package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queuememberentitylisting
type Queuememberentitylisting struct { 
	// Entities
	Entities *[]Queuemember `json:"entities,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// FirstUri
	FirstUri *string `json:"firstUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`

}

func (o *Queuememberentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queuememberentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Queuemember `json:"entities,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		FirstUri *string `json:"firstUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		PageNumber: o.PageNumber,
		
		PageSize: o.PageSize,
		
		FirstUri: o.FirstUri,
		
		SelfUri: o.SelfUri,
		
		PreviousUri: o.PreviousUri,
		
		NextUri: o.NextUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Queuememberentitylisting) UnmarshalJSON(b []byte) error {
	var QueuememberentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &QueuememberentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := QueuememberentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageNumber, ok := QueuememberentitylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := QueuememberentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if FirstUri, ok := QueuememberentitylistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
	
	if SelfUri, ok := QueuememberentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if PreviousUri, ok := QueuememberentitylistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
	
	if NextUri, ok := QueuememberentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queuememberentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
