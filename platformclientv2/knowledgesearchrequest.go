package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesearchrequest
type Knowledgesearchrequest struct { 
	// Query - Input query to search content in the knowledge base
	Query *string `json:"query,omitempty"`


	// PageSize - Page size of the returned results
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - Page number of the returned results
	PageNumber *int `json:"pageNumber,omitempty"`


	// DocumentType - Document type to be used while searching
	DocumentType *string `json:"documentType,omitempty"`


	// LanguageCode - query search for specific languageCode
	LanguageCode *string `json:"languageCode,omitempty"`


	// SearchOnDraftDocuments - If true the search query will be executed on draft documents, else it will be on active documents
	SearchOnDraftDocuments *bool `json:"searchOnDraftDocuments,omitempty"`

}

func (u *Knowledgesearchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgesearchrequest

	

	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		DocumentType *string `json:"documentType,omitempty"`
		
		LanguageCode *string `json:"languageCode,omitempty"`
		
		SearchOnDraftDocuments *bool `json:"searchOnDraftDocuments,omitempty"`
		*Alias
	}{ 
		Query: u.Query,
		
		PageSize: u.PageSize,
		
		PageNumber: u.PageNumber,
		
		DocumentType: u.DocumentType,
		
		LanguageCode: u.LanguageCode,
		
		SearchOnDraftDocuments: u.SearchOnDraftDocuments,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Knowledgesearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
