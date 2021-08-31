package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting
type Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting struct { 
	// Entities
	Entities *[]Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult `json:"entities,omitempty"`

}

func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting
	
	return json.Marshal(&struct { 
		Entities *[]Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting) UnmarshalJSON(b []byte) error {
	var WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlistingMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
