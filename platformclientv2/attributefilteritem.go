package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Attributefilteritem
type Attributefilteritem struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Values
	Values *[]string `json:"values,omitempty"`

}

func (u *Attributefilteritem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Attributefilteritem

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Operator: u.Operator,
		
		Values: u.Values,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Attributefilteritem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
