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

func (o *Knowledgeimport) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeimport
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		UploadKey: o.UploadKey,
		
		FileType: o.FileType,
		
		IgnoreHeaders: o.IgnoreHeaders,
		
		Status: o.Status,
		
		Report: o.Report,
		
		KnowledgeBase: o.KnowledgeBase,
		
		LanguageCode: o.LanguageCode,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeimport) UnmarshalJSON(b []byte) error {
	var KnowledgeimportMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeimportMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgeimportMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := KnowledgeimportMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if UploadKey, ok := KnowledgeimportMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    
	if FileType, ok := KnowledgeimportMap["fileType"].(string); ok {
		o.FileType = &FileType
	}
    
	if IgnoreHeaders, ok := KnowledgeimportMap["ignoreHeaders"].(bool); ok {
		o.IgnoreHeaders = &IgnoreHeaders
	}
    
	if Status, ok := KnowledgeimportMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Report, ok := KnowledgeimportMap["report"].(map[string]interface{}); ok {
		ReportString, _ := json.Marshal(Report)
		json.Unmarshal(ReportString, &o.Report)
	}
	
	if KnowledgeBase, ok := KnowledgeimportMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if LanguageCode, ok := KnowledgeimportMap["languageCode"].(string); ok {
		o.LanguageCode = &LanguageCode
	}
    
	if dateCreatedString, ok := KnowledgeimportMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgeimportMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SelfUri, ok := KnowledgeimportMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeimport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
