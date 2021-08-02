package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeimport
type Knowledgeimport struct { 
	// Id - Id of the import operation
	Id *string `json:"id,omitempty"`


	// Name - Name of the import operation
	Name *string `json:"name,omitempty"`


	// UploadKey - Upload key
	UploadKey *string `json:"uploadKey,omitempty"`


	// FileType - file type of the document
	FileType *string `json:"fileType,omitempty"`


	// IgnoreHeaders - Ignore headers for the specified file
	IgnoreHeaders *bool `json:"ignoreHeaders,omitempty"`


	// Status - Status of the operation
	Status *string `json:"status,omitempty"`


	// Report - Report of the import operation
	Report *Importreport `json:"report,omitempty"`


	// KnowledgeBase - Knowledge base which document import does belong to
	KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`


	// LanguageCode - Language code
	LanguageCode *string `json:"languageCode,omitempty"`


	// DateCreated - Created date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Knowledgeimport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
