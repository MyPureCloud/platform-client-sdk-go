package platformclientv2
import (
	"encoding/json"
)

// Groupgreetingeventgreeting
type Groupgreetingeventgreeting struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// OwnerType
	OwnerType *string `json:"ownerType,omitempty"`


	// Owner
	Owner *Groupgreetingeventgreetingowner `json:"owner,omitempty"`


	// GreetingAudioFile
	GreetingAudioFile *Groupgreetingeventgreetingaudiofile `json:"greetingAudioFile,omitempty"`


	// AudioTTS
	AudioTTS *string `json:"audioTTS,omitempty"`

}

// String returns a JSON representation of the model
func (o *Groupgreetingeventgreeting) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
