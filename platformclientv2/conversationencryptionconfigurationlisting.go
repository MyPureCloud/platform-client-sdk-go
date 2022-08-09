package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationencryptionconfigurationlisting
type Conversationencryptionconfigurationlisting struct { 
	// Total
	Total *int `json:"total,omitempty"`


	// Entities
	Entities *[]Conversationencryptionconfiguration `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Conversationencryptionconfigurationlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationencryptionconfigurationlisting
	
	return json.Marshal(&struct { 
		Total *int `json:"total,omitempty"`
		
		Entities *[]Conversationencryptionconfiguration `json:"entities,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Total: o.Total,
		
		Entities: o.Entities,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationencryptionconfigurationlisting) UnmarshalJSON(b []byte) error {
	var ConversationencryptionconfigurationlistingMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationencryptionconfigurationlistingMap)
	if err != nil {
		return err
	}
	
	if Total, ok := ConversationencryptionconfigurationlistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if Entities, ok := ConversationencryptionconfigurationlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if SelfUri, ok := ConversationencryptionconfigurationlistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationencryptionconfigurationlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
