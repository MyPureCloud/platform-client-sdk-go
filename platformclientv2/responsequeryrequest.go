package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responsequeryrequest - Used to query for responses
type Responsequeryrequest struct { 
	// QueryPhrase - Query phrase to search response text and name. If not set will match all.
	QueryPhrase *string `json:"queryPhrase,omitempty"`


	// PageSize - The maximum number of hits to return. Default: 25, Maximum: 500.
	PageSize *int `json:"pageSize,omitempty"`


	// Filters - Filter the query results.
	Filters *[]Responsefilter `json:"filters,omitempty"`

}

func (o *Responsequeryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responsequeryrequest
	
	return json.Marshal(&struct { 
		QueryPhrase *string `json:"queryPhrase,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		Filters *[]Responsefilter `json:"filters,omitempty"`
		*Alias
	}{ 
		QueryPhrase: o.QueryPhrase,
		
		PageSize: o.PageSize,
		
		Filters: o.Filters,
		Alias:    (*Alias)(o),
	})
}

func (o *Responsequeryrequest) UnmarshalJSON(b []byte) error {
	var ResponsequeryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ResponsequeryrequestMap)
	if err != nil {
		return err
	}
	
	if QueryPhrase, ok := ResponsequeryrequestMap["queryPhrase"].(string); ok {
		o.QueryPhrase = &QueryPhrase
	}
    
	if PageSize, ok := ResponsequeryrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if Filters, ok := ResponsequeryrequestMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Responsequeryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
