package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencebulkquery
type Wfmhistoricaladherencebulkquery struct { 
	// Items - The historical adherence items to query
	Items *[]Wfmhistoricaladherencebulkitem `json:"items,omitempty"`


	// TimeZone - The time zone, in olson format, to use in defining days when computing adherence. The results will be returned as UTC timestamps regardless of the time zone input.
	TimeZone *string `json:"timeZone,omitempty"`

}

func (o *Wfmhistoricaladherencebulkquery) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherencebulkquery
	
	return json.Marshal(&struct { 
		Items *[]Wfmhistoricaladherencebulkitem `json:"items,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		*Alias
	}{ 
		Items: o.Items,
		
		TimeZone: o.TimeZone,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherencebulkquery) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherencebulkqueryMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherencebulkqueryMap)
	if err != nil {
		return err
	}
	
	if Items, ok := WfmhistoricaladherencebulkqueryMap["items"].([]interface{}); ok {
		ItemsString, _ := json.Marshal(Items)
		json.Unmarshal(ItemsString, &o.Items)
	}
	
	if TimeZone, ok := WfmhistoricaladherencebulkqueryMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
