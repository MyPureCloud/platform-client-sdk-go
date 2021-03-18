package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Searchaggregation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
