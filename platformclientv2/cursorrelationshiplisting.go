package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Cursorrelationshiplisting
type Cursorrelationshiplisting struct { 
	// Entities
	Entities *[]Relationship `json:"entities,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// Cursors - The cursor that points to the next set of entities being returned.
	Cursors *Cursors `json:"cursors,omitempty"`

}

func (o *Cursorrelationshiplisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Cursorrelationshiplisting
	
	return json.Marshal(&struct { 
		Entities *[]Relationship `json:"entities,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		
		Cursors *Cursors `json:"cursors,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		NextUri: o.NextUri,
		
		SelfUri: o.SelfUri,
		
		PreviousUri: o.PreviousUri,
		
		Cursors: o.Cursors,
		Alias:    (*Alias)(o),
	})
}

func (o *Cursorrelationshiplisting) UnmarshalJSON(b []byte) error {
	var CursorrelationshiplistingMap map[string]interface{}
	err := json.Unmarshal(b, &CursorrelationshiplistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := CursorrelationshiplistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if NextUri, ok := CursorrelationshiplistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
    
	if SelfUri, ok := CursorrelationshiplistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if PreviousUri, ok := CursorrelationshiplistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
    
	if Cursors, ok := CursorrelationshiplistingMap["cursors"].(map[string]interface{}); ok {
		CursorsString, _ := json.Marshal(Cursors)
		json.Unmarshal(CursorsString, &o.Cursors)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Cursorrelationshiplisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
