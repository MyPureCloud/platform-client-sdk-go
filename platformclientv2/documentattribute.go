package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentattribute
type Documentattribute struct { 
	// Attribute
	Attribute *Attribute `json:"attribute,omitempty"`


	// Values
	Values *[]string `json:"values,omitempty"`

}

func (u *Documentattribute) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentattribute

	

	return json.Marshal(&struct { 
		Attribute *Attribute `json:"attribute,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Attribute: u.Attribute,
		
		Values: u.Values,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Documentattribute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
