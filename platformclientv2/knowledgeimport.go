package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Knowledgeimport) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeimport

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		UploadKey *string `json:"uploadKey,omitempty"`
		
		FileType *string `json:"fileType,omitempty"`
		
		IgnoreHeaders *bool `json:"ignoreHeaders,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Report *Importreport `json:"report,omitempty"`
		
		KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`
		
		LanguageCode *string `json:"languageCode,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		UploadKey: u.UploadKey,
		
		FileType: u.FileType,
		
		IgnoreHeaders: u.IgnoreHeaders,
		
		Status: u.Status,
		
		Report: u.Report,
		
		KnowledgeBase: u.KnowledgeBase,
		
		LanguageCode: u.LanguageCode,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Knowledgeimport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
