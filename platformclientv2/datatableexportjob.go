package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Datatableexportjob - State information for an export job of rows from a datatable
type Datatableexportjob struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Owner - The PureCloud user who started the export job
	Owner *Addressableentityref `json:"owner,omitempty"`


	// Status - The status of the export job
	Status *string `json:"status,omitempty"`


	// DateCreated - The timestamp of when the export began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateCompleted - The timestamp of when the export stopped (either successfully or unsuccessfully). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// DownloadURI - The URL of the location at which the caller can download the export file, when available
	DownloadURI *string `json:"downloadURI,omitempty"`


	// ErrorInformation - Any error information, or null of the processing is not in an error state
	ErrorInformation *Errorbody `json:"errorInformation,omitempty"`


	// CountRecordsProcessed - The current count of the number of records processed
	CountRecordsProcessed *int `json:"countRecordsProcessed,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Datatableexportjob) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Datatableexportjob
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Owner *Addressableentityref `json:"owner,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		DownloadURI *string `json:"downloadURI,omitempty"`
		
		ErrorInformation *Errorbody `json:"errorInformation,omitempty"`
		
		CountRecordsProcessed *int `json:"countRecordsProcessed,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Owner: o.Owner,
		
		Status: o.Status,
		
		DateCreated: DateCreated,
		
		DateCompleted: DateCompleted,
		
		DownloadURI: o.DownloadURI,
		
		ErrorInformation: o.ErrorInformation,
		
		CountRecordsProcessed: o.CountRecordsProcessed,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Datatableexportjob) UnmarshalJSON(b []byte) error {
	var DatatableexportjobMap map[string]interface{}
	err := json.Unmarshal(b, &DatatableexportjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DatatableexportjobMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DatatableexportjobMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Owner, ok := DatatableexportjobMap["owner"].(map[string]interface{}); ok {
		OwnerString, _ := json.Marshal(Owner)
		json.Unmarshal(OwnerString, &o.Owner)
	}
	
	if Status, ok := DatatableexportjobMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if dateCreatedString, ok := DatatableexportjobMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateCompletedString, ok := DatatableexportjobMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if DownloadURI, ok := DatatableexportjobMap["downloadURI"].(string); ok {
		o.DownloadURI = &DownloadURI
	}
	
	if ErrorInformation, ok := DatatableexportjobMap["errorInformation"].(map[string]interface{}); ok {
		ErrorInformationString, _ := json.Marshal(ErrorInformation)
		json.Unmarshal(ErrorInformationString, &o.ErrorInformation)
	}
	
	if CountRecordsProcessed, ok := DatatableexportjobMap["countRecordsProcessed"].(float64); ok {
		CountRecordsProcessedInt := int(CountRecordsProcessed)
		o.CountRecordsProcessed = &CountRecordsProcessedInt
	}
	
	if SelfUri, ok := DatatableexportjobMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Datatableexportjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
