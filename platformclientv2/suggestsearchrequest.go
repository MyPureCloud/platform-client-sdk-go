package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Suggestsearchrequest
type Suggestsearchrequest struct { 
	// Expand - Provides more details about a specified resource
	Expand *[]string `json:"expand,omitempty"`


	// Types - Resource domain type to search
	Types *[]string `json:"types,omitempty"`


	// Query - Suggest query
	Query *[]Suggestsearchcriteria `json:"query,omitempty"`

}

func (o *Suggestsearchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Suggestsearchrequest
	
	return json.Marshal(&struct { 
		Expand *[]string `json:"expand,omitempty"`
		
		Types *[]string `json:"types,omitempty"`
		
		Query *[]Suggestsearchcriteria `json:"query,omitempty"`
		*Alias
	}{ 
		Expand: o.Expand,
		
		Types: o.Types,
		
		Query: o.Query,
		Alias:    (*Alias)(o),
	})
}

func (o *Suggestsearchrequest) UnmarshalJSON(b []byte) error {
	var SuggestsearchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SuggestsearchrequestMap)
	if err != nil {
		return err
	}
	
	if Expand, ok := SuggestsearchrequestMap["expand"].([]interface{}); ok {
		ExpandString, _ := json.Marshal(Expand)
		json.Unmarshal(ExpandString, &o.Expand)
	}
	
	if Types, ok := SuggestsearchrequestMap["types"].([]interface{}); ok {
		TypesString, _ := json.Marshal(Types)
		json.Unmarshal(TypesString, &o.Types)
	}
	
	if Query, ok := SuggestsearchrequestMap["query"].([]interface{}); ok {
		QueryString, _ := json.Marshal(Query)
		json.Unmarshal(QueryString, &o.Query)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Suggestsearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
