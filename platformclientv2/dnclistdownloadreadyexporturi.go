package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dnclistdownloadreadyexporturi
type Dnclistdownloadreadyexporturi struct { 
	// Uri
	Uri *string `json:"uri,omitempty"`


	// ExportTimestamp
	ExportTimestamp *string `json:"exportTimestamp,omitempty"`

}

func (o *Dnclistdownloadreadyexporturi) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dnclistdownloadreadyexporturi
	
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

func (o *Dnclistdownloadreadyexporturi) UnmarshalJSON(b []byte) error {
	var DnclistdownloadreadyexporturiMap map[string]interface{}
	err := json.Unmarshal(b, &DnclistdownloadreadyexporturiMap)
	if err != nil {
		return err
	}
	
	if Uri, ok := DnclistdownloadreadyexporturiMap["uri"].(string); ok {
		o.Uri = &Uri
	}
    
	if ExportTimestamp, ok := DnclistdownloadreadyexporturiMap["exportTimestamp"].(string); ok {
		o.ExportTimestamp = &ExportTimestamp
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dnclistdownloadreadyexporturi) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
