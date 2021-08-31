package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkidsrequest
type Bulkidsrequest struct { 
	// Entities
	Entities *[]Entity `json:"entities,omitempty"`

}

func (o *Bulkidsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkidsrequest
	
	return json.Marshal(&struct { 
		Entities *[]Entity `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkidsrequest) UnmarshalJSON(b []byte) error {
	var BulkidsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BulkidsrequestMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BulkidsrequestMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkidsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
