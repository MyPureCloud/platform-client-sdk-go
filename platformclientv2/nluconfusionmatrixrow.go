package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluconfusionmatrixrow
type Nluconfusionmatrixrow struct { 
	// Name - The name of the intent for the row.
	Name *string `json:"name,omitempty"`


	// Columns - The columns of confusion matrix for the intent
	Columns *[]Nluconfusionmatrixcolumn `json:"columns,omitempty"`

}

func (o *Nluconfusionmatrixrow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluconfusionmatrixrow
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Columns *[]Nluconfusionmatrixcolumn `json:"columns,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Columns: o.Columns,
		Alias:    (*Alias)(o),
	})
}

func (o *Nluconfusionmatrixrow) UnmarshalJSON(b []byte) error {
	var NluconfusionmatrixrowMap map[string]interface{}
	err := json.Unmarshal(b, &NluconfusionmatrixrowMap)
	if err != nil {
		return err
	}
	
	if Name, ok := NluconfusionmatrixrowMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Columns, ok := NluconfusionmatrixrowMap["columns"].([]interface{}); ok {
		ColumnsString, _ := json.Marshal(Columns)
		json.Unmarshal(ColumnsString, &o.Columns)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nluconfusionmatrixrow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
