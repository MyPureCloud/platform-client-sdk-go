package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Orphanrecording
type Orphanrecording struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// CreatedTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedTime *time.Time `json:"createdTime,omitempty"`


	// RecoveredTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RecoveredTime *time.Time `json:"recoveredTime,omitempty"`


	// ProviderType
	ProviderType *string `json:"providerType,omitempty"`


	// MediaSizeBytes
	MediaSizeBytes *int `json:"mediaSizeBytes,omitempty"`


	// MediaType
	MediaType *string `json:"mediaType,omitempty"`


	// FileState
	FileState *string `json:"fileState,omitempty"`


	// ProviderEndpoint
	ProviderEndpoint *Endpoint `json:"providerEndpoint,omitempty"`


	// Recording
	Recording *Recording `json:"recording,omitempty"`


	// OrphanStatus - The status of the orphaned recording's conversation.
	OrphanStatus *string `json:"orphanStatus,omitempty"`


	// SourceOrphaningId - An identifier used during recovery operations by the supplying hybrid platform to track back and determine which interaction this recording is associated with
	SourceOrphaningId *string `json:"sourceOrphaningId,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Orphanrecording) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Orphanrecording

	
	CreatedTime := new(string)
	if u.CreatedTime != nil {
		
		*CreatedTime = timeutil.Strftime(u.CreatedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedTime = nil
	}
	
	RecoveredTime := new(string)
	if u.RecoveredTime != nil {
		
		*RecoveredTime = timeutil.Strftime(u.RecoveredTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RecoveredTime = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		CreatedTime *string `json:"createdTime,omitempty"`
		
		RecoveredTime *string `json:"recoveredTime,omitempty"`
		
		ProviderType *string `json:"providerType,omitempty"`
		
		MediaSizeBytes *int `json:"mediaSizeBytes,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		FileState *string `json:"fileState,omitempty"`
		
		ProviderEndpoint *Endpoint `json:"providerEndpoint,omitempty"`
		
		Recording *Recording `json:"recording,omitempty"`
		
		OrphanStatus *string `json:"orphanStatus,omitempty"`
		
		SourceOrphaningId *string `json:"sourceOrphaningId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		CreatedTime: CreatedTime,
		
		RecoveredTime: RecoveredTime,
		
		ProviderType: u.ProviderType,
		
		MediaSizeBytes: u.MediaSizeBytes,
		
		MediaType: u.MediaType,
		
		FileState: u.FileState,
		
		ProviderEndpoint: u.ProviderEndpoint,
		
		Recording: u.Recording,
		
		OrphanStatus: u.OrphanStatus,
		
		SourceOrphaningId: u.SourceOrphaningId,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Orphanrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
