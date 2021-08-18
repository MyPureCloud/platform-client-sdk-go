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

func (u *Datatableexportjob) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Datatableexportjob

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateCompleted := new(string)
	if u.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(u.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		Name: u.Name,
		
		Owner: u.Owner,
		
		Status: u.Status,
		
		DateCreated: DateCreated,
		
		DateCompleted: DateCompleted,
		
		DownloadURI: u.DownloadURI,
		
		ErrorInformation: u.ErrorInformation,
		
		CountRecordsProcessed: u.CountRecordsProcessed,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Datatableexportjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
