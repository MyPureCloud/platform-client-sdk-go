package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkshifttradestateupdaterequest
type Bulkshifttradestateupdaterequest struct { 
	// Entities - The shift trades to update
	Entities *[]Bulkupdateshifttradestaterequestitem `json:"entities,omitempty"`

}

func (o *Bulkshifttradestateupdaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkshifttradestateupdaterequest
	
	return json.Marshal(&struct { 
		Entities *[]Bulkupdateshifttradestaterequestitem `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkshifttradestateupdaterequest) UnmarshalJSON(b []byte) error {
	var BulkshifttradestateupdaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &BulkshifttradestateupdaterequestMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := BulkshifttradestateupdaterequestMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkshifttradestateupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
