package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Orphanrecording) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
