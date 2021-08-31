package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkorganizationsrequest
type Bulkorganizationsrequest struct { 
	// Entities
	Entities *[]Externalorganization `json:"entities,omitempty"`

}

func (o *Bulkorganizationsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkorganizationsrequest
	
	return json.Marshal(&struct { 
		Entities *[]Externalorganization `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkorganizationsrequest) UnmarshalJSON(b []byte) error {
	var BulkorganizationsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BulkorganizationsrequestMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BulkorganizationsrequestMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkorganizationsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
