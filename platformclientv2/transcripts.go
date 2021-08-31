package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcripts
type Transcripts struct { 
	// ExactMatch - List of transcript contents which needs to satisfy exact match criteria
	ExactMatch *[]string `json:"exactMatch,omitempty"`


	// Contains - List of transcript contents which needs to satisfy contains criteria
	Contains *[]string `json:"contains,omitempty"`


	// DoesNotContain - List of transcript contents which needs to satisfy does not contain criteria
	DoesNotContain *[]string `json:"doesNotContain,omitempty"`

}

func (o *Transcripts) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcripts
	
	return json.Marshal(&struct { 
		ExactMatch *[]string `json:"exactMatch,omitempty"`
		
		Contains *[]string `json:"contains,omitempty"`
		
		DoesNotContain *[]string `json:"doesNotContain,omitempty"`
		*Alias
	}{ 
		ExactMatch: o.ExactMatch,
		
		Contains: o.Contains,
		
		DoesNotContain: o.DoesNotContain,
		Alias:    (*Alias)(o),
	})
}

func (o *Transcripts) UnmarshalJSON(b []byte) error {
	var TranscriptsMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptsMap)
	if err != nil {
		return err
	}
	
	if ExactMatch, ok := TranscriptsMap["exactMatch"].([]interface{}); ok {
		ExactMatchString, _ := json.Marshal(ExactMatch)
		json.Unmarshal(ExactMatchString, &o.ExactMatch)
	}
	
	if Contains, ok := TranscriptsMap["contains"].([]interface{}); ok {
		ContainsString, _ := json.Marshal(Contains)
		json.Unmarshal(ContainsString, &o.Contains)
	}
	
	if DoesNotContain, ok := TranscriptsMap["doesNotContain"].([]interface{}); ok {
		DoesNotContainString, _ := json.Marshal(DoesNotContain)
		json.Unmarshal(DoesNotContainString, &o.DoesNotContain)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcripts) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
