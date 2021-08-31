package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Exportscriptresponse
type Exportscriptresponse struct { 
	// Url
	Url *string `json:"url,omitempty"`

}

func (o *Exportscriptresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Exportscriptresponse
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		Alias:    (*Alias)(o),
	})
}

func (o *Exportscriptresponse) UnmarshalJSON(b []byte) error {
	var ExportscriptresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ExportscriptresponseMap)
	if err != nil {
		return err
	}
	
	if Url, ok := ExportscriptresponseMap["url"].(string); ok {
		o.Url = &Url
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Exportscriptresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
