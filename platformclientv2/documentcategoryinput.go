package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentcategoryinput
type Documentcategoryinput struct { 
	// Id - KnowledgeBase Category ID
	Id *string `json:"id,omitempty"`

}

func (o *Documentcategoryinput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentcategoryinput
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentcategoryinput) UnmarshalJSON(b []byte) error {
	var DocumentcategoryinputMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentcategoryinputMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DocumentcategoryinputMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentcategoryinput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
