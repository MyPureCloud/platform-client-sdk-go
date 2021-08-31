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

func (o *Importreport) MarshalJSON() ([]byte, error) {
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
		Errors: o.Errors,
		
		Validated: o.Validated,
		
		Imported: o.Imported,
		
		TotalDocuments: o.TotalDocuments,
		Alias:    (*Alias)(o),
	})
}

func (o *Importreport) UnmarshalJSON(b []byte) error {
	var ImportreportMap map[string]interface{}
	err := json.Unmarshal(b, &ImportreportMap)
	if err != nil {
		return err
	}
	
	if Errors, ok := ImportreportMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	
	if Validated, ok := ImportreportMap["validated"].(map[string]interface{}); ok {
		ValidatedString, _ := json.Marshal(Validated)
		json.Unmarshal(ValidatedString, &o.Validated)
	}
	
	if Imported, ok := ImportreportMap["imported"].(map[string]interface{}); ok {
		ImportedString, _ := json.Marshal(Imported)
		json.Unmarshal(ImportedString, &o.Imported)
	}
	
	if TotalDocuments, ok := ImportreportMap["totalDocuments"].(float64); ok {
		TotalDocumentsInt := int(TotalDocuments)
		o.TotalDocuments = &TotalDocumentsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Importreport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
