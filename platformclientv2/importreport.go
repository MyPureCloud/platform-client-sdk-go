package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Importreport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
