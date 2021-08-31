package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Faxtopicworkspacedata
type Faxtopicworkspacedata struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Faxtopicworkspacedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Faxtopicworkspacedata
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Faxtopicworkspacedata) UnmarshalJSON(b []byte) error {
	var FaxtopicworkspacedataMap map[string]interface{}
	err := json.Unmarshal(b, &FaxtopicworkspacedataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FaxtopicworkspacedataMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Faxtopicworkspacedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
