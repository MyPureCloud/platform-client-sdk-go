package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adhocrecordingtopicconversationdata
type Adhocrecordingtopicconversationdata struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Adhocrecordingtopicconversationdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adhocrecordingtopicconversationdata
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Adhocrecordingtopicconversationdata) UnmarshalJSON(b []byte) error {
	var AdhocrecordingtopicconversationdataMap map[string]interface{}
	err := json.Unmarshal(b, &AdhocrecordingtopicconversationdataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AdhocrecordingtopicconversationdataMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Adhocrecordingtopicconversationdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
