package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Guestcategoryresponselisting
type Guestcategoryresponselisting struct { 
	// Entities
	Entities *[]Guestcategoryresponse `json:"entities,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// SessionId
	SessionId *string `json:"sessionId,omitempty"`

}

func (o *Guestcategoryresponselisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Guestcategoryresponselisting
	
	return json.Marshal(&struct { 
		Entities *[]Guestcategoryresponse `json:"entities,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		NextUri: o.NextUri,
		
		SelfUri: o.SelfUri,
		
		PreviousUri: o.PreviousUri,
		
		SessionId: o.SessionId,
		Alias:    (*Alias)(o),
	})
}

func (o *Guestcategoryresponselisting) UnmarshalJSON(b []byte) error {
	var GuestcategoryresponselistingMap map[string]interface{}
	err := json.Unmarshal(b, &GuestcategoryresponselistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := GuestcategoryresponselistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if NextUri, ok := GuestcategoryresponselistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
    
	if SelfUri, ok := GuestcategoryresponselistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if PreviousUri, ok := GuestcategoryresponselistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
    
	if SessionId, ok := GuestcategoryresponselistingMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Guestcategoryresponselisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
