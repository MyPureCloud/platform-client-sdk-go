package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemailsearchcriteria
type Voicemailsearchcriteria struct { 
	// EndValue - The end value of the range. This field is used for range search types.
	EndValue *string `json:"endValue,omitempty"`


	// Values - A list of values for the search to match against
	Values *[]string `json:"values,omitempty"`


	// StartValue - The start value of the range. This field is used for range search types.
	StartValue *string `json:"startValue,omitempty"`


	// Fields - Field names to search against
	Fields *[]string `json:"fields,omitempty"`


	// Value - A value for the search to match against
	Value *string `json:"value,omitempty"`


	// Operator - How to apply this search criteria against other criteria
	Operator *string `json:"operator,omitempty"`


	// Group - Groups multiple conditions
	Group *[]Voicemailsearchcriteria `json:"group,omitempty"`


	// DateFormat - Set date format for criteria values when using date range search type.  Supports Java date format syntax, example yyyy-MM-dd'T'HH:mm:ss.SSSX.
	DateFormat *string `json:"dateFormat,omitempty"`


	// VarType - Search Type
	VarType *string `json:"type,omitempty"`

}

func (u *Voicemailsearchcriteria) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemailsearchcriteria

	

	return json.Marshal(&struct { 
		EndValue *string `json:"endValue,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		StartValue *string `json:"startValue,omitempty"`
		
		Fields *[]string `json:"fields,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Group *[]Voicemailsearchcriteria `json:"group,omitempty"`
		
		DateFormat *string `json:"dateFormat,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		EndValue: u.EndValue,
		
		Values: u.Values,
		
		StartValue: u.StartValue,
		
		Fields: u.Fields,
		
		Value: u.Value,
		
		Operator: u.Operator,
		
		Group: u.Group,
		
		DateFormat: u.DateFormat,
		
		VarType: u.VarType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Voicemailsearchcriteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
