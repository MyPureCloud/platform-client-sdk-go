package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturnintent
type Reportingturnintent struct { 
	// Name - The name of the intent detected during this reporting turn.
	Name *string `json:"name,omitempty"`


	// Confidence - The confidence score of the intent detected during this reporting turn.
	Confidence *float64 `json:"confidence,omitempty"`


	// Slots - The slots detected during this reporting turn.
	Slots *[]Reportingturnintentslot `json:"slots,omitempty"`

}

func (u *Reportingturnintent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingturnintent

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Confidence *float64 `json:"confidence,omitempty"`
		
		Slots *[]Reportingturnintentslot `json:"slots,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Confidence: u.Confidence,
		
		Slots: u.Slots,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportingturnintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
