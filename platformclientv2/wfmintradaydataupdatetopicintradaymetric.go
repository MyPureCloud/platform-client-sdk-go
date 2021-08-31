package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradaydataupdatetopicintradaymetric
type Wfmintradaydataupdatetopicintradaymetric struct { 
	// Category
	Category *string `json:"category,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`

}

func (o *Wfmintradaydataupdatetopicintradaymetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradaymetric
	
	return json.Marshal(&struct { 
		Category *string `json:"category,omitempty"`
		
		Version *string `json:"version,omitempty"`
		*Alias
	}{ 
		Category: o.Category,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintradaydataupdatetopicintradaymetric) UnmarshalJSON(b []byte) error {
	var WfmintradaydataupdatetopicintradaymetricMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintradaydataupdatetopicintradaymetricMap)
	if err != nil {
		return err
	}
	
	if Category, ok := WfmintradaydataupdatetopicintradaymetricMap["category"].(string); ok {
		o.Category = &Category
	}
	
	if Version, ok := WfmintradaydataupdatetopicintradaymetricMap["version"].(string); ok {
		o.Version = &Version
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradaymetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
