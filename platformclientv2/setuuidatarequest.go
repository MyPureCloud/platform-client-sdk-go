package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Setuuidatarequest
type Setuuidatarequest struct { 
	// UuiData - The value of the uuiData to set.
	UuiData *string `json:"uuiData,omitempty"`

}

func (o *Setuuidatarequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Setuuidatarequest
	
	return json.Marshal(&struct { 
		UuiData *string `json:"uuiData,omitempty"`
		*Alias
	}{ 
		UuiData: o.UuiData,
		Alias:    (*Alias)(o),
	})
}

func (o *Setuuidatarequest) UnmarshalJSON(b []byte) error {
	var SetuuidatarequestMap map[string]interface{}
	err := json.Unmarshal(b, &SetuuidatarequestMap)
	if err != nil {
		return err
	}
	
	if UuiData, ok := SetuuidatarequestMap["uuiData"].(string); ok {
		o.UuiData = &UuiData
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Setuuidatarequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
