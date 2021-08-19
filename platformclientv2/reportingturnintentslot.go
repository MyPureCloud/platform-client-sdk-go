package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturnintentslot
type Reportingturnintentslot struct { 
	// Name - The name of the slot.
	Name *string `json:"name,omitempty"`


	// Value - The value of the slot.
	Value *string `json:"value,omitempty"`


	// VarType - The NLU entity type of the slot (either builtin or user defined)
	VarType *string `json:"type,omitempty"`


	// Confidence - The confidence score this slot received during detection.
	Confidence *float64 `json:"confidence,omitempty"`

}

func (u *Reportingturnintentslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingturnintentslot

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Confidence *float64 `json:"confidence,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Value: u.Value,
		
		VarType: u.VarType,
		
		Confidence: u.Confidence,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportingturnintentslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
