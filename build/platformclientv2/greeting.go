package platformclientv2
import (
	"time"
	"encoding/json"
)

// Greeting
type Greeting struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType - Greeting type
	VarType *string `json:"type,omitempty"`


	// OwnerType - Greeting owner type
	OwnerType *string `json:"ownerType,omitempty"`


	// Owner - Greeting owner
	Owner *Domainentity `json:"owner,omitempty"`


	// AudioFile
	AudioFile *Greetingaudiofile `json:"audioFile,omitempty"`


	// AudioTTS
	AudioTTS *string `json:"audioTTS,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// CreatedBy
	CreatedBy *string `json:"createdBy,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// ModifiedBy
	ModifiedBy *string `json:"modifiedBy,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Greeting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
