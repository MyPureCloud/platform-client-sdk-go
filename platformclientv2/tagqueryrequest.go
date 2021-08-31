package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Tagqueryrequest
type Tagqueryrequest struct { 
	// Query
	Query *string `json:"query,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`

}

func (o *Tagqueryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Tagqueryrequest
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		*Alias
	}{ 
		Query: o.Query,
		
		PageNumber: o.PageNumber,
		
		PageSize: o.PageSize,
		Alias:    (*Alias)(o),
	})
}

func (o *Tagqueryrequest) UnmarshalJSON(b []byte) error {
	var TagqueryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TagqueryrequestMap)
	if err != nil {
		return err
	}
	
	if Query, ok := TagqueryrequestMap["query"].(string); ok {
		o.Query = &Query
	}
	
	if PageNumber, ok := TagqueryrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := TagqueryrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Tagqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
