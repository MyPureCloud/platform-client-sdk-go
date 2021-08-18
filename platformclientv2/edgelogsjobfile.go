package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgelogsjobfile
type Edgelogsjobfile struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`


	// Description - The resource's description.
	Description *string `json:"description,omitempty"`


	// Version - The current version of the resource.
	Version *int `json:"version,omitempty"`


	// DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - The ID of the user that last modified the resource.
	ModifiedBy *string `json:"modifiedBy,omitempty"`


	// CreatedBy - The ID of the user that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`


	// State - Indicates if the resource is active, inactive, or deleted.
	State *string `json:"state,omitempty"`


	// ModifiedByApp - The application that last modified the resource.
	ModifiedByApp *string `json:"modifiedByApp,omitempty"`


	// CreatedByApp - The application that created the resource.
	CreatedByApp *string `json:"createdByApp,omitempty"`


	// TimeCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	TimeCreated *time.Time `json:"timeCreated,omitempty"`


	// TimeModified - The time this log file was last modified on the Edge. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	TimeModified *time.Time `json:"timeModified,omitempty"`


	// SizeBytes - The size of this file in bytes.
	SizeBytes *float64 `json:"sizeBytes,omitempty"`


	// UploadStatus - The status of the upload of this file from the Edge to the cloud.  Use /upload to start an upload.
	UploadStatus *string `json:"uploadStatus,omitempty"`


	// EdgePath - The path of this file on the Edge.
	EdgePath *string `json:"edgePath,omitempty"`


	// DownloadId - The download ID to use with the downloads API.
	DownloadId *string `json:"downloadId,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Edgelogsjobfile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgelogsjobfile

	
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
	
	TimeCreated := new(string)
	if u.TimeCreated != nil {
		
		*TimeCreated = timeutil.Strftime(u.TimeCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		TimeCreated = nil
	}
	
	TimeModified := new(string)
	if u.TimeModified != nil {
		
		*TimeModified = timeutil.Strftime(u.TimeModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		TimeModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ModifiedByApp *string `json:"modifiedByApp,omitempty"`
		
		CreatedByApp *string `json:"createdByApp,omitempty"`
		
		TimeCreated *string `json:"timeCreated,omitempty"`
		
		TimeModified *string `json:"timeModified,omitempty"`
		
		SizeBytes *float64 `json:"sizeBytes,omitempty"`
		
		UploadStatus *string `json:"uploadStatus,omitempty"`
		
		EdgePath *string `json:"edgePath,omitempty"`
		
		DownloadId *string `json:"downloadId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Version: u.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: u.ModifiedBy,
		
		CreatedBy: u.CreatedBy,
		
		State: u.State,
		
		ModifiedByApp: u.ModifiedByApp,
		
		CreatedByApp: u.CreatedByApp,
		
		TimeCreated: TimeCreated,
		
		TimeModified: TimeModified,
		
		SizeBytes: u.SizeBytes,
		
		UploadStatus: u.UploadStatus,
		
		EdgePath: u.EdgePath,
		
		DownloadId: u.DownloadId,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgelogsjobfile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
