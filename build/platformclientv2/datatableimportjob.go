package platformclientv2
import (
	"time"
	"encoding/json"
)

// Datatableimportjob - State information for an import job of rows to a datatable
type Datatableimportjob struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Owner - The PureCloud user who started the import job
	Owner *Addressableentityref `json:"owner,omitempty"`


	// Status - The status of the import job
	Status *string `json:"status,omitempty"`


	// DateCreated - The timestamp of when the import began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateCompleted - The timestamp of when the import stopped (either successfully or unsuccessfully). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// UploadURI - The URL of the location at which the caller can upload the file to be imported
	UploadURI *string `json:"uploadURI,omitempty"`


	// ImportMode - The indication of whether the processing should remove rows that don't appear in the import file
	ImportMode *string `json:"importMode,omitempty"`


	// ErrorInformation - Any error information, or null of the processing is not in an error state
	ErrorInformation *Errorbody `json:"errorInformation,omitempty"`


	// CountRecordsUpdated - The current count of the number of records processed
	CountRecordsUpdated *int `json:"countRecordsUpdated,omitempty"`


	// CountRecordsDeleted - The current count of the number of records deleted
	CountRecordsDeleted *int `json:"countRecordsDeleted,omitempty"`


	// CountRecordsFailed - The current count of the number of records that failed to import
	CountRecordsFailed *int `json:"countRecordsFailed,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Datatableimportjob) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
