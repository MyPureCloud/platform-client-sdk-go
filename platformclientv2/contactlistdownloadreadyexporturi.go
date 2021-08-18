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


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Contactlistdownloadreadyexporturi) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistdownloadreadyexporturi

	

	return json.Marshal(&struct { 
		Uri *string `json:"uri,omitempty"`
		
		ExportTimestamp *string `json:"exportTimestamp,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Uri: u.Uri,
		
		ExportTimestamp: u.ExportTimestamp,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contactlistdownloadreadyexporturi) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
