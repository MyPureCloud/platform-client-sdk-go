package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Cursororganizationlisting
type Cursororganizationlisting struct { 
	// Entities
	Entities *[]Externalorganization `json:"entities,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`

}

func (o *Cursororganizationlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Cursororganizationlisting
	
	return json.Marshal(&struct { 
		Entities *[]Externalorganization `json:"entities,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		NextUri: o.NextUri,
		
		SelfUri: o.SelfUri,
		
		PreviousUri: o.PreviousUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Cursororganizationlisting) UnmarshalJSON(b []byte) error {
	var CursororganizationlistingMap map[string]interface{}
	err := json.Unmarshal(b, &CursororganizationlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := CursororganizationlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if NextUri, ok := CursororganizationlistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
	
	if SelfUri, ok := CursororganizationlistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if PreviousUri, ok := CursororganizationlistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Cursororganizationlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
