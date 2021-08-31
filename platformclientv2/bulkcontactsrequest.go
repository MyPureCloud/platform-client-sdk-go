package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkcontactsrequest
type Bulkcontactsrequest struct { 
	// Entities
	Entities *[]Externalcontact `json:"entities,omitempty"`

}

func (o *Bulkcontactsrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkcontactsrequest
	
	return json.Marshal(&struct { 
		Entities *[]Externalcontact `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkcontactsrequest) UnmarshalJSON(b []byte) error {
	var BulkcontactsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &BulkcontactsrequestMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BulkcontactsrequestMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkcontactsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
