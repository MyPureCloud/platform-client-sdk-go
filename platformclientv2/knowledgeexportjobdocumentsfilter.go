package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeexportjobdocumentsfilter
type Knowledgeexportjobdocumentsfilter struct { 
	// Interval - Retrieves the documents modified in specified date and time range. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`

}

func (o *Knowledgeexportjobdocumentsfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeexportjobdocumentsfilter
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeexportjobdocumentsfilter) UnmarshalJSON(b []byte) error {
	var KnowledgeexportjobdocumentsfilterMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeexportjobdocumentsfilterMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := KnowledgeexportjobdocumentsfilterMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeexportjobdocumentsfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
