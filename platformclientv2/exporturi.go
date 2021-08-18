package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Exporturi
type Exporturi struct { 
	// Uri
	Uri *string `json:"uri,omitempty"`


	// ExportTimestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExportTimestamp *time.Time `json:"exportTimestamp,omitempty"`

}

func (u *Exporturi) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Exporturi

	
	ExportTimestamp := new(string)
	if u.ExportTimestamp != nil {
		
		*ExportTimestamp = timeutil.Strftime(u.ExportTimestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExportTimestamp = nil
	}
	

	return json.Marshal(&struct { 
		Uri *string `json:"uri,omitempty"`
		
		ExportTimestamp *string `json:"exportTimestamp,omitempty"`
		*Alias
	}{ 
		Uri: u.Uri,
		
		ExportTimestamp: ExportTimestamp,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Exporturi) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
