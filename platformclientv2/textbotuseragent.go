package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotuseragent - Information about the caller executing a bot flow.
type Textbotuseragent struct { 
	// Name - The name of the user agent.
	Name *string `json:"name,omitempty"`

}

func (o *Textbotuseragent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotuseragent
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotuseragent) UnmarshalJSON(b []byte) error {
	var TextbotuseragentMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotuseragentMap)
	if err != nil {
		return err
	}
	
	if Name, ok := TextbotuseragentMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotuseragent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
