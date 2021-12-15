package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistdownloadreadyexporturi
type Contactlistdownloadreadyexporturi struct { 
	// Uri
	Uri *string `json:"uri,omitempty"`


	// ExportTimestamp
	ExportTimestamp *string `json:"exportTimestamp,omitempty"`

}

func (o *Contactlistdownloadreadyexporturi) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistdownloadreadyexporturi
	
	return json.Marshal(&struct { 
		Uri *string `json:"uri,omitempty"`
		
		ExportTimestamp *string `json:"exportTimestamp,omitempty"`
		*Alias
	}{ 
		Uri: o.Uri,
		
		ExportTimestamp: o.ExportTimestamp,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactlistdownloadreadyexporturi) UnmarshalJSON(b []byte) error {
	var ContactlistdownloadreadyexporturiMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistdownloadreadyexporturiMap)
	if err != nil {
		return err
	}
	
	if Uri, ok := ContactlistdownloadreadyexporturiMap["uri"].(string); ok {
		o.Uri = &Uri
	}
	
	if ExportTimestamp, ok := ContactlistdownloadreadyexporturiMap["exportTimestamp"].(string); ok {
		o.ExportTimestamp = &ExportTimestamp
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistdownloadreadyexporturi) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
