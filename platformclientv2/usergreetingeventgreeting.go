package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Usergreetingeventgreeting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Usergreetingeventgreeting

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		OwnerType *string `json:"ownerType,omitempty"`
		
		Owner *Usergreetingeventgreetingowner `json:"owner,omitempty"`
		
		GreetingAudioFile *Usergreetingeventgreetingaudiofile `json:"greetingAudioFile,omitempty"`
		
		AudioTTS *string `json:"audioTTS,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		VarType: u.VarType,
		
		OwnerType: u.OwnerType,
		
		Owner: u.Owner,
		
		GreetingAudioFile: u.GreetingAudioFile,
		
		AudioTTS: u.AudioTTS,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Usergreetingeventgreeting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
