package platformclientv2
import (
	"encoding/json"
)

// Knowledgesearchrequest
type Knowledgesearchrequest struct { 
	// Query - Input query to search content in the knowledge base
	Query *string `json:"query,omitempty"`


	// PageSize - Page size of the returned results
	PageSize *int32 `json:"pageSize,omitempty"`


	// PageNumber - Page number of the returned results
	PageNumber *int32 `json:"pageNumber,omitempty"`


	// DocumentType - Document type to be used while searching
	DocumentType *string `json:"documentType,omitempty"`


	// LanguageCode - query search for specific languageCode
	LanguageCode *string `json:"languageCode,omitempty"`


	// SearchOnDraftDocuments - If true the search query will be executed on draft documents, else it will be on active documents
	SearchOnDraftDocuments *bool `json:"searchOnDraftDocuments,omitempty"`

}

// String returns a JSON representation of the model
func (o *Knowledgesearchrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
