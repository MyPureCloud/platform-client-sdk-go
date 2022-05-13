package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerauditrequest
type Dialerauditrequest struct { 
	// QueryPhrase - The word or words to search for.
	QueryPhrase *string `json:"queryPhrase,omitempty"`


	// QueryFields - The fields in which to search for the queryPhrase.
	QueryFields *[]string `json:"queryFields,omitempty"`


	// Facets - The fields to facet on.
	Facets *[]Auditfacet `json:"facets,omitempty"`


	// Filters - The fields to filter on.
	Filters *[]Auditfilter `json:"filters,omitempty"`

}

func (o *Dialerauditrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerauditrequest
	
	return json.Marshal(&struct { 
		QueryPhrase *string `json:"queryPhrase,omitempty"`
		
		QueryFields *[]string `json:"queryFields,omitempty"`
		
		Facets *[]Auditfacet `json:"facets,omitempty"`
		
		Filters *[]Auditfilter `json:"filters,omitempty"`
		*Alias
	}{ 
		QueryPhrase: o.QueryPhrase,
		
		QueryFields: o.QueryFields,
		
		Facets: o.Facets,
		
		Filters: o.Filters,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerauditrequest) UnmarshalJSON(b []byte) error {
	var DialerauditrequestMap map[string]interface{}
	err := json.Unmarshal(b, &DialerauditrequestMap)
	if err != nil {
		return err
	}
	
	if QueryPhrase, ok := DialerauditrequestMap["queryPhrase"].(string); ok {
		o.QueryPhrase = &QueryPhrase
	}
    
	if QueryFields, ok := DialerauditrequestMap["queryFields"].([]interface{}); ok {
		QueryFieldsString, _ := json.Marshal(QueryFields)
		json.Unmarshal(QueryFieldsString, &o.QueryFields)
	}
	
	if Facets, ok := DialerauditrequestMap["facets"].([]interface{}); ok {
		FacetsString, _ := json.Marshal(Facets)
		json.Unmarshal(FacetsString, &o.Facets)
	}
	
	if Filters, ok := DialerauditrequestMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerauditrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
