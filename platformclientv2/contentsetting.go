package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentsetting
type Contentsetting struct { 
	// Story - Settings relating to facebook and instagram stories feature
	Story *Storysetting `json:"story,omitempty"`

}

func (o *Contentsetting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentsetting
	
	return json.Marshal(&struct { 
		Story *Storysetting `json:"story,omitempty"`
		*Alias
	}{ 
		Story: o.Story,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentsetting) UnmarshalJSON(b []byte) error {
	var ContentsettingMap map[string]interface{}
	err := json.Unmarshal(b, &ContentsettingMap)
	if err != nil {
		return err
	}
	
	if Story, ok := ContentsettingMap["story"].(map[string]interface{}); ok {
		StoryString, _ := json.Marshal(Story)
		json.Unmarshal(StoryString, &o.Story)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentsetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
