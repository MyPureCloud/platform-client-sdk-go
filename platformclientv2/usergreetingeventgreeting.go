package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Usergreetingeventgreeting
type Usergreetingeventgreeting struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// OwnerType
	OwnerType *string `json:"ownerType,omitempty"`


	// Owner
	Owner *Usergreetingeventgreetingowner `json:"owner,omitempty"`


	// GreetingAudioFile
	GreetingAudioFile *Usergreetingeventgreetingaudiofile `json:"greetingAudioFile,omitempty"`


	// AudioTTS
	AudioTTS *string `json:"audioTTS,omitempty"`

}

// String returns a JSON representation of the model
func (o *Usergreetingeventgreeting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
