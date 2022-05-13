package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adhocrecordingtopicworkspacedata
type Adhocrecordingtopicworkspacedata struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Adhocrecordingtopicworkspacedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adhocrecordingtopicworkspacedata
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Adhocrecordingtopicworkspacedata) UnmarshalJSON(b []byte) error {
	var AdhocrecordingtopicworkspacedataMap map[string]interface{}
	err := json.Unmarshal(b, &AdhocrecordingtopicworkspacedataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AdhocrecordingtopicworkspacedataMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Adhocrecordingtopicworkspacedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
