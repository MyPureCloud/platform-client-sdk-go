package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeexportjobresponse
type Knowledgeexportjobresponse struct { 
	// Id - Id of the export job.
	Id *string `json:"id,omitempty"`


	// DownloadURL - The URL of the location at which the caller can download the export file, when available.
	DownloadURL *string `json:"downloadURL,omitempty"`


	// FileType - File type of the document
	FileType *string `json:"fileType,omitempty"`


	// CountDocumentProcessed - The current count of the number of records processed.
	CountDocumentProcessed *int `json:"countDocumentProcessed,omitempty"`


	// ExportFilter - Filters to narrow down what to export.
	ExportFilter *Knowledgeexportjobfilter `json:"exportFilter,omitempty"`


	// Status - The status of the export job.
	Status *string `json:"status,omitempty"`


	// KnowledgeBase - Knowledge base which document export belongs to.
	KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`


	// DateCreated - The timestamp of when the export began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The timestamp of when the export stopped. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ErrorInformation - Any error information, or null of the processing is not in failed state.
	ErrorInformation *Errorbody `json:"errorInformation,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Knowledgeexportjobresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeexportjobresponse
	
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
		
		DownloadURL *string `json:"downloadURL,omitempty"`
		
		FileType *string `json:"fileType,omitempty"`
		
		CountDocumentProcessed *int `json:"countDocumentProcessed,omitempty"`
		
		ExportFilter *Knowledgeexportjobfilter `json:"exportFilter,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ErrorInformation *Errorbody `json:"errorInformation,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		DownloadURL: o.DownloadURL,
		
		FileType: o.FileType,
		
		CountDocumentProcessed: o.CountDocumentProcessed,
		
		ExportFilter: o.ExportFilter,
		
		Status: o.Status,
		
		KnowledgeBase: o.KnowledgeBase,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ErrorInformation: o.ErrorInformation,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeexportjobresponse) UnmarshalJSON(b []byte) error {
	var KnowledgeexportjobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeexportjobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgeexportjobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DownloadURL, ok := KnowledgeexportjobresponseMap["downloadURL"].(string); ok {
		o.DownloadURL = &DownloadURL
	}
    
	if FileType, ok := KnowledgeexportjobresponseMap["fileType"].(string); ok {
		o.FileType = &FileType
	}
    
	if CountDocumentProcessed, ok := KnowledgeexportjobresponseMap["countDocumentProcessed"].(float64); ok {
		CountDocumentProcessedInt := int(CountDocumentProcessed)
		o.CountDocumentProcessed = &CountDocumentProcessedInt
	}
	
	if ExportFilter, ok := KnowledgeexportjobresponseMap["exportFilter"].(map[string]interface{}); ok {
		ExportFilterString, _ := json.Marshal(ExportFilter)
		json.Unmarshal(ExportFilterString, &o.ExportFilter)
	}
	
	if Status, ok := KnowledgeexportjobresponseMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if KnowledgeBase, ok := KnowledgeexportjobresponseMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if dateCreatedString, ok := KnowledgeexportjobresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgeexportjobresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ErrorInformation, ok := KnowledgeexportjobresponseMap["errorInformation"].(map[string]interface{}); ok {
		ErrorInformationString, _ := json.Marshal(ErrorInformation)
		json.Unmarshal(ErrorInformationString, &o.ErrorInformation)
	}
	
	if SelfUri, ok := KnowledgeexportjobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeexportjobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
