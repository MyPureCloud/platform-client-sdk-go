package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Responsequeryrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
