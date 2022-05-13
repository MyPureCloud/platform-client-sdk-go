package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatmessageentitylist
type Webchatmessageentitylist struct { 
	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// Entities
	Entities *[]Webchatmessage `json:"entities,omitempty"`


	// PreviousPage
	PreviousPage *string `json:"previousPage,omitempty"`


	// Next
	Next *string `json:"next,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Webchatmessageentitylist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatmessageentitylist
	
	return json.Marshal(&struct { 
		PageSize *int `json:"pageSize,omitempty"`
		
		Entities *[]Webchatmessage `json:"entities,omitempty"`
		
		PreviousPage *string `json:"previousPage,omitempty"`
		
		Next *string `json:"next,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		PageSize: o.PageSize,
		
		Entities: o.Entities,
		
		PreviousPage: o.PreviousPage,
		
		Next: o.Next,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Webchatmessageentitylist) UnmarshalJSON(b []byte) error {
	var WebchatmessageentitylistMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatmessageentitylistMap)
	if err != nil {
		return err
	}
	
	if PageSize, ok := WebchatmessageentitylistMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if Entities, ok := WebchatmessageentitylistMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PreviousPage, ok := WebchatmessageentitylistMap["previousPage"].(string); ok {
		o.PreviousPage = &PreviousPage
	}
    
	if Next, ok := WebchatmessageentitylistMap["next"].(string); ok {
		o.Next = &Next
	}
    
	if SelfUri, ok := WebchatmessageentitylistMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatmessageentitylist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
