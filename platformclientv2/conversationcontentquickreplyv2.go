package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentquickreplyv2 - Quick reply object V2.
type Conversationcontentquickreplyv2 struct { 
	// Title - Text to show as the title of the quick reply.
	Title *string `json:"title,omitempty"`


	// Actions - An array of quick reply objects.
	Actions *[]Conversationcontentquickreply `json:"actions,omitempty"`

}

func (o *Conversationcontentquickreplyv2) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcontentquickreplyv2
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Actions *[]Conversationcontentquickreply `json:"actions,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcontentquickreplyv2) UnmarshalJSON(b []byte) error {
	var Conversationcontentquickreplyv2Map map[string]interface{}
	err := json.Unmarshal(b, &Conversationcontentquickreplyv2Map)
	if err != nil {
		return err
	}
	
	if Title, ok := Conversationcontentquickreplyv2Map["title"].(string); ok {
		o.Title = &Title
	}
    
	if Actions, ok := Conversationcontentquickreplyv2Map["actions"].([]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentquickreplyv2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
