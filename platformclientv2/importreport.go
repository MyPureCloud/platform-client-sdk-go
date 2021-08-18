package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Importreport
type Importreport struct { 
	// Errors
	Errors *[]Importerror `json:"errors,omitempty"`


	// Validated
	Validated *Resultcounters `json:"validated,omitempty"`


	// Imported
	Imported *Resultcounters `json:"imported,omitempty"`


	// TotalDocuments
	TotalDocuments *int `json:"totalDocuments,omitempty"`

}

func (u *Importreport) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Importreport

	

	return json.Marshal(&struct { 
		Errors *[]Importerror `json:"errors,omitempty"`
		
		Validated *Resultcounters `json:"validated,omitempty"`
		
		Imported *Resultcounters `json:"imported,omitempty"`
		
		TotalDocuments *int `json:"totalDocuments,omitempty"`
		*Alias
	}{ 
		Errors: u.Errors,
		
		Validated: u.Validated,
		
		Imported: u.Imported,
		
		TotalDocuments: u.TotalDocuments,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Importreport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
