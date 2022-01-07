package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Ucicon
type Ucicon struct { 
	// Vector - vector
	Vector *string `json:"vector,omitempty"`

}

func (o *Ucicon) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Ucicon
	
	return json.Marshal(&struct { 
		Vector *string `json:"vector,omitempty"`
		*Alias
	}{ 
		Vector: o.Vector,
		Alias:    (*Alias)(o),
	})
}

func (o *Ucicon) UnmarshalJSON(b []byte) error {
	var UciconMap map[string]interface{}
	err := json.Unmarshal(b, &UciconMap)
	if err != nil {
		return err
	}
	
	if Vector, ok := UciconMap["vector"].(string); ok {
		o.Vector = &Vector
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ucicon) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
