package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Facetentry
type Facetentry struct { 
	// Attribute
	Attribute *Termattribute `json:"attribute,omitempty"`


	// Statistics
	Statistics *Facetstatistics `json:"statistics,omitempty"`


	// Other
	Other *int `json:"other,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// Missing
	Missing *int `json:"missing,omitempty"`


	// TermCount
	TermCount *int `json:"termCount,omitempty"`


	// TermType
	TermType *string `json:"termType,omitempty"`


	// Terms
	Terms *[]Facetterm `json:"terms,omitempty"`

}

// String returns a JSON representation of the model
func (o *Facetentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
