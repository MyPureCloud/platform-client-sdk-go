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

func (o *Edgelogsjobfile) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgelogsjobfile
	
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
	
	TimeCreated := new(string)
	if o.TimeCreated != nil {
		
		*TimeCreated = timeutil.Strftime(o.TimeCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		TimeCreated = nil
	}
	
	TimeModified := new(string)
	if o.TimeModified != nil {
		
		*TimeModified = timeutil.Strftime(o.TimeModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		CreatedBy: o.CreatedBy,
		
		State: o.State,
		
		ModifiedByApp: o.ModifiedByApp,
		
		CreatedByApp: o.CreatedByApp,
		
		TimeCreated: TimeCreated,
		
		TimeModified: TimeModified,
		
		SizeBytes: o.SizeBytes,
		
		UploadStatus: o.UploadStatus,
		
		EdgePath: o.EdgePath,
		
		DownloadId: o.DownloadId,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgelogsjobfile) UnmarshalJSON(b []byte) error {
	var EdgelogsjobfileMap map[string]interface{}
	err := json.Unmarshal(b, &EdgelogsjobfileMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EdgelogsjobfileMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := EdgelogsjobfileMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := EdgelogsjobfileMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := EdgelogsjobfileMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := EdgelogsjobfileMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := EdgelogsjobfileMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := EdgelogsjobfileMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := EdgelogsjobfileMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := EdgelogsjobfileMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := EdgelogsjobfileMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := EdgelogsjobfileMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if timeCreatedString, ok := EdgelogsjobfileMap["timeCreated"].(string); ok {
		TimeCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeCreatedString)
		o.TimeCreated = &TimeCreated
	}
	
	if timeModifiedString, ok := EdgelogsjobfileMap["timeModified"].(string); ok {
		TimeModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeModifiedString)
		o.TimeModified = &TimeModified
	}
	
	if SizeBytes, ok := EdgelogsjobfileMap["sizeBytes"].(float64); ok {
		o.SizeBytes = &SizeBytes
	}
	
	if UploadStatus, ok := EdgelogsjobfileMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
	
	if EdgePath, ok := EdgelogsjobfileMap["edgePath"].(string); ok {
		o.EdgePath = &EdgePath
	}
	
	if DownloadId, ok := EdgelogsjobfileMap["downloadId"].(string); ok {
		o.DownloadId = &DownloadId
	}
	
	if SelfUri, ok := EdgelogsjobfileMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgelogsjobfile) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
