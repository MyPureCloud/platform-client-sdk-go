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

func (o *Exporturi) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Exporturi
	
	ExportTimestamp := new(string)
	if o.ExportTimestamp != nil {
		
		*ExportTimestamp = timeutil.Strftime(o.ExportTimestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExportTimestamp = nil
	}
	
	return json.Marshal(&struct { 
		Uri *string `json:"uri,omitempty"`
		
		ExportTimestamp *string `json:"exportTimestamp,omitempty"`
		*Alias
	}{ 
		Uri: o.Uri,
		
		ExportTimestamp: ExportTimestamp,
		Alias:    (*Alias)(o),
	})
}

func (o *Exporturi) UnmarshalJSON(b []byte) error {
	var ExporturiMap map[string]interface{}
	err := json.Unmarshal(b, &ExporturiMap)
	if err != nil {
		return err
	}
	
	if Uri, ok := ExporturiMap["uri"].(string); ok {
		o.Uri = &Uri
	}
    
	if exportTimestampString, ok := ExporturiMap["exportTimestamp"].(string); ok {
		ExportTimestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", exportTimestampString)
		o.ExportTimestamp = &ExportTimestamp
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Exporturi) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
