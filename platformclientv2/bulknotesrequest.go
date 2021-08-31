package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulknotesrequest
type Bulknotesrequest struct { 
	// Entities
	Entities *[]Note `json:"entities,omitempty"`

}

func (o *Bulknotesrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulknotesrequest
	
	return json.Marshal(&struct { 
		Entities *[]Note `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulknotesrequest) UnmarshalJSON(b []byte) error {
	var BulknotesrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BulknotesrequestMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BulknotesrequestMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulknotesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
