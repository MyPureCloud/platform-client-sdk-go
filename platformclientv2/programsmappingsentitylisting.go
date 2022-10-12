package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Programsmappingsentitylisting
type Programsmappingsentitylisting struct { 
	// Entities
	Entities *[]Programmappings `json:"entities,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Programsmappingsentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Programsmappingsentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Programmappings `json:"entities,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		PageSize: o.PageSize,
		
		NextUri: o.NextUri,
		
		PageCount: o.PageCount,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Programsmappingsentitylisting) UnmarshalJSON(b []byte) error {
	var ProgramsmappingsentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &ProgramsmappingsentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := ProgramsmappingsentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := ProgramsmappingsentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if NextUri, ok := ProgramsmappingsentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
    
	if PageCount, ok := ProgramsmappingsentitylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	
	if SelfUri, ok := ProgramsmappingsentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Programsmappingsentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
