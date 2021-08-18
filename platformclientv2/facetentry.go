package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Facetentry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facetentry

	

	return json.Marshal(&struct { 
		Attribute *Termattribute `json:"attribute,omitempty"`
		
		Statistics *Facetstatistics `json:"statistics,omitempty"`
		
		Other *int `json:"other,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		Missing *int `json:"missing,omitempty"`
		
		TermCount *int `json:"termCount,omitempty"`
		
		TermType *string `json:"termType,omitempty"`
		
		Terms *[]Facetterm `json:"terms,omitempty"`
		*Alias
	}{ 
		Attribute: u.Attribute,
		
		Statistics: u.Statistics,
		
		Other: u.Other,
		
		Total: u.Total,
		
		Missing: u.Missing,
		
		TermCount: u.TermCount,
		
		TermType: u.TermType,
		
		Terms: u.Terms,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Facetentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
