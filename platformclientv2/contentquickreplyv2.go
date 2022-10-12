package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentquickreplyv2 - Quick reply object V2.
type Contentquickreplyv2 struct { 
	// Title - Text to show as the title of the quick reply.
	Title *string `json:"title,omitempty"`


	// Actions - An array of quick reply objects.
	Actions *[]Contentquickreply `json:"actions,omitempty"`

}

func (o *Contentquickreplyv2) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentquickreplyv2
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Actions *[]Contentquickreply `json:"actions,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Actions: o.Actions,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentquickreplyv2) UnmarshalJSON(b []byte) error {
	var Contentquickreplyv2Map map[string]interface{}
	err := json.Unmarshal(b, &Contentquickreplyv2Map)
	if err != nil {
		return err
	}
	
	if Title, ok := Contentquickreplyv2Map["title"].(string); ok {
		o.Title = &Title
	}
    
	if Actions, ok := Contentquickreplyv2Map["actions"].([]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentquickreplyv2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
