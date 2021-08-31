package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Generaltopic
type Generaltopic struct { 
	// Name
	Name *string `json:"name,omitempty"`

}

func (o *Generaltopic) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Generaltopic
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Generaltopic) UnmarshalJSON(b []byte) error {
	var GeneraltopicMap map[string]interface{}
	err := json.Unmarshal(b, &GeneraltopicMap)
	if err != nil {
		return err
	}
	
	if Name, ok := GeneraltopicMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Generaltopic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
