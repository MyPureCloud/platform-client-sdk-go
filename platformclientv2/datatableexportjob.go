package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Datatableexportjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
