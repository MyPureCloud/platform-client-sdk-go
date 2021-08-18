package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Record
type Record struct { 
	// Name - The name of the record.
	Name *string `json:"name,omitempty"`


	// VarType - The type of the record. (Example values:  MX, TXT, CNAME)
	VarType *string `json:"type,omitempty"`


	// Value - The value of the record.
	Value *string `json:"value,omitempty"`

}

func (u *Record) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Record

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		VarType: u.VarType,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Record) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
