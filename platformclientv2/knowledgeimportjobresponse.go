package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeimportjobresponse
type Knowledgeimportjobresponse struct { 
	// Id - Id of the import job
	Id *string `json:"id,omitempty"`


	// UploadKey - Upload key
	UploadKey *string `json:"uploadKey,omitempty"`


	// FileType - File type of the document
	FileType *string `json:"fileType,omitempty"`


	// Settings - Additional optional settings
	Settings *Knowledgeimportjobsettings `json:"settings,omitempty"`


	// Status - Status of the import job
	Status *string `json:"status,omitempty"`


	// Report - Report of the import job
	Report *Knowledgeimportjobreport `json:"report,omitempty"`


	// KnowledgeBase - Knowledge base which document import does belong to
	KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`


	// DateCreated - Created date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Knowledgeimportjobresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeimportjobresponse
	
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
		
		UploadKey *string `json:"uploadKey,omitempty"`
		
		FileType *string `json:"fileType,omitempty"`
		
		Settings *Knowledgeimportjobsettings `json:"settings,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Report *Knowledgeimportjobreport `json:"report,omitempty"`
		
		KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		UploadKey: o.UploadKey,
		
		FileType: o.FileType,
		
		Settings: o.Settings,
		
		Status: o.Status,
		
		Report: o.Report,
		
		KnowledgeBase: o.KnowledgeBase,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeimportjobresponse) UnmarshalJSON(b []byte) error {
	var KnowledgeimportjobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeimportjobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgeimportjobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if UploadKey, ok := KnowledgeimportjobresponseMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    
	if FileType, ok := KnowledgeimportjobresponseMap["fileType"].(string); ok {
		o.FileType = &FileType
	}
    
	if Settings, ok := KnowledgeimportjobresponseMap["settings"].(map[string]interface{}); ok {
		SettingsString, _ := json.Marshal(Settings)
		json.Unmarshal(SettingsString, &o.Settings)
	}
	
	if Status, ok := KnowledgeimportjobresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Report, ok := KnowledgeimportjobresponseMap["report"].(map[string]interface{}); ok {
		ReportString, _ := json.Marshal(Report)
		json.Unmarshal(ReportString, &o.Report)
	}
	
	if KnowledgeBase, ok := KnowledgeimportjobresponseMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if dateCreatedString, ok := KnowledgeimportjobresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgeimportjobresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SelfUri, ok := KnowledgeimportjobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
