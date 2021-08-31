package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Pagingspec
type Pagingspec struct { 
	// PageSize - How many results per page
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - How many pages in
	PageNumber *int `json:"pageNumber,omitempty"`

}

func (o *Pagingspec) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Pagingspec
	
	return json.Marshal(&struct { 
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		*Alias
	}{ 
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		Alias:    (*Alias)(o),
	})
}

func (o *Pagingspec) UnmarshalJSON(b []byte) error {
	var PagingspecMap map[string]interface{}
	err := json.Unmarshal(b, &PagingspecMap)
	if err != nil {
		return err
	}
	
	if PageSize, ok := PagingspecMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := PagingspecMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Pagingspec) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
