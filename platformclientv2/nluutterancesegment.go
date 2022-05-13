package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluutterancesegment
type Nluutterancesegment struct { 
	// Text - The text of the segment.
	Text *string `json:"text,omitempty"`


	// Entity - The entity annotation of the segment.
	Entity *Namedentityannotation `json:"entity,omitempty"`

}

func (o *Nluutterancesegment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluutterancesegment
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		Entity *Namedentityannotation `json:"entity,omitempty"`
		*Alias
	}{ 
		Text: o.Text,
		
		Entity: o.Entity,
		Alias:    (*Alias)(o),
	})
}

func (o *Nluutterancesegment) UnmarshalJSON(b []byte) error {
	var NluutterancesegmentMap map[string]interface{}
	err := json.Unmarshal(b, &NluutterancesegmentMap)
	if err != nil {
		return err
	}
	
	if Text, ok := NluutterancesegmentMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Entity, ok := NluutterancesegmentMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nluutterancesegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
