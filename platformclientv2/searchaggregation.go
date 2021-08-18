package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Searchaggregation
type Searchaggregation struct { 
	// Field - The field used for aggregation
	Field *string `json:"field,omitempty"`


	// Name - The name of the aggregation. The response aggregation uses this name.
	Name *string `json:"name,omitempty"`


	// VarType - The type of aggregation to perform
	VarType *string `json:"type,omitempty"`


	// Value - A value to use for aggregation
	Value *string `json:"value,omitempty"`


	// Size - The number aggregations results to return out of the entire result set
	Size *int `json:"size,omitempty"`


	// Order - The order in which aggregation results are sorted
	Order *[]string `json:"order,omitempty"`

}

func (u *Searchaggregation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Searchaggregation

	

	return json.Marshal(&struct { 
		Field *string `json:"field,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		Order *[]string `json:"order,omitempty"`
		*Alias
	}{ 
		Field: u.Field,
		
		Name: u.Name,
		
		VarType: u.VarType,
		
		Value: u.Value,
		
		Size: u.Size,
		
		Order: u.Order,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Searchaggregation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
