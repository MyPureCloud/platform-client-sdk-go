package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Searchsort
type Searchsort struct { 
	// SortOrder - The sort order for results
	SortOrder *string `json:"sortOrder,omitempty"`


	// SortBy - The field in the resource that you want to sort the results by
	SortBy *string `json:"sortBy,omitempty"`

}

func (o *Searchsort) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Searchsort
	
	return json.Marshal(&struct { 
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		*Alias
	}{ 
		SortOrder: o.SortOrder,
		
		SortBy: o.SortBy,
		Alias:    (*Alias)(o),
	})
}

func (o *Searchsort) UnmarshalJSON(b []byte) error {
	var SearchsortMap map[string]interface{}
	err := json.Unmarshal(b, &SearchsortMap)
	if err != nil {
		return err
	}
	
	if SortOrder, ok := SearchsortMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
	
	if SortBy, ok := SearchsortMap["sortBy"].(string); ok {
		o.SortBy = &SortBy
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Searchsort) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
