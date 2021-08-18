package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Groupgreetingeventgreeting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Groupgreetingeventgreeting

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		OwnerType *string `json:"ownerType,omitempty"`
		
		Owner *Groupgreetingeventgreetingowner `json:"owner,omitempty"`
		
		GreetingAudioFile *Groupgreetingeventgreetingaudiofile `json:"greetingAudioFile,omitempty"`
		
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
func (o *Groupgreetingeventgreeting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
