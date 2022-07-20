package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalmetricdatawriterequest
type Externalmetricdatawriterequest struct { 
	// Items - A list of external metric data items. A maximum of 100 items are allowed.
	Items *[]Externalmetricdataitem `json:"items,omitempty"`

}

func (o *Externalmetricdatawriterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalmetricdatawriterequest
	
	return json.Marshal(&struct { 
		Items *[]Externalmetricdataitem `json:"items,omitempty"`
		*Alias
	}{ 
		Items: o.Items,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalmetricdatawriterequest) UnmarshalJSON(b []byte) error {
	var ExternalmetricdatawriterequestMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalmetricdatawriterequestMap)
	if err != nil {
		return err
	}
	
	if Items, ok := ExternalmetricdatawriterequestMap["items"].([]interface{}); ok {
		ItemsString, _ := json.Marshal(Items)
		json.Unmarshal(ItemsString, &o.Items)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalmetricdatawriterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
