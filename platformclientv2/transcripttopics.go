package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcripttopics
type Transcripttopics struct { 
	// Includes - List of topics which need to be included in exact match criteria. This field is not mutually exclusive with excludes topic list.
	Includes *[]string `json:"includes,omitempty"`


	// Excludes - List of topics which need to be excluded in exact match criteria. This field is not mutually exclusive with includes topic list.
	Excludes *[]string `json:"excludes,omitempty"`

}

func (o *Transcripttopics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcripttopics
	
	return json.Marshal(&struct { 
		Includes *[]string `json:"includes,omitempty"`
		
		Excludes *[]string `json:"excludes,omitempty"`
		*Alias
	}{ 
		Includes: o.Includes,
		
		Excludes: o.Excludes,
		Alias:    (*Alias)(o),
	})
}

func (o *Transcripttopics) UnmarshalJSON(b []byte) error {
	var TranscripttopicsMap map[string]interface{}
	err := json.Unmarshal(b, &TranscripttopicsMap)
	if err != nil {
		return err
	}
	
	if Includes, ok := TranscripttopicsMap["includes"].([]interface{}); ok {
		IncludesString, _ := json.Marshal(Includes)
		json.Unmarshal(IncludesString, &o.Includes)
	}
	
	if Excludes, ok := TranscripttopicsMap["excludes"].([]interface{}); ok {
		ExcludesString, _ := json.Marshal(Excludes)
		json.Unmarshal(ExcludesString, &o.Excludes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcripttopics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
