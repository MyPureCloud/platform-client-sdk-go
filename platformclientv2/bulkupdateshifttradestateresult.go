package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkupdateshifttradestateresult
type Bulkupdateshifttradestateresult struct { 
	// Entities
	Entities *[]Bulkupdateshifttradestateresultitem `json:"entities,omitempty"`

}

func (o *Bulkupdateshifttradestateresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkupdateshifttradestateresult
	
	return json.Marshal(&struct { 
		Entities *[]Bulkupdateshifttradestateresultitem `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkupdateshifttradestateresult) UnmarshalJSON(b []byte) error {
	var BulkupdateshifttradestateresultMap map[string]interface{}
	err := json.Unmarshal(b, &BulkupdateshifttradestateresultMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BulkupdateshifttradestateresultMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkupdateshifttradestateresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
