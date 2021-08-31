package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Deleteretention
type Deleteretention struct { 
	// Days
	Days *int `json:"days,omitempty"`

}

func (o *Deleteretention) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Deleteretention
	
	return json.Marshal(&struct { 
		Days *int `json:"days,omitempty"`
		*Alias
	}{ 
		Days: o.Days,
		Alias:    (*Alias)(o),
	})
}

func (o *Deleteretention) UnmarshalJSON(b []byte) error {
	var DeleteretentionMap map[string]interface{}
	err := json.Unmarshal(b, &DeleteretentionMap)
	if err != nil {
		return err
	}
	
	if Days, ok := DeleteretentionMap["days"].(float64); ok {
		DaysInt := int(Days)
		o.Days = &DaysInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Deleteretention) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
