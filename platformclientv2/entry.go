package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Entry
type Entry struct { 
	// Value - A value included in this facet.
	Value *string `json:"value,omitempty"`


	// Count - The number of results with this value.
	Count *int `json:"count,omitempty"`

}

func (o *Entry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Entry
	
	return json.Marshal(&struct { 
		Value *string `json:"value,omitempty"`
		
		Count *int `json:"count,omitempty"`
		*Alias
	}{ 
		Value: o.Value,
		
		Count: o.Count,
		Alias:    (*Alias)(o),
	})
}

func (o *Entry) UnmarshalJSON(b []byte) error {
	var EntryMap map[string]interface{}
	err := json.Unmarshal(b, &EntryMap)
	if err != nil {
		return err
	}
	
	if Value, ok := EntryMap["value"].(string); ok {
		o.Value = &Value
	}
	
	if Count, ok := EntryMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Entry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
