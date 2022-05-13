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

func (o *Usergreetingeventgreeting) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		OwnerType: o.OwnerType,
		
		Owner: o.Owner,
		
		GreetingAudioFile: o.GreetingAudioFile,
		
		AudioTTS: o.AudioTTS,
		Alias:    (*Alias)(o),
	})
}

func (o *Usergreetingeventgreeting) UnmarshalJSON(b []byte) error {
	var UsergreetingeventgreetingMap map[string]interface{}
	err := json.Unmarshal(b, &UsergreetingeventgreetingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UsergreetingeventgreetingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UsergreetingeventgreetingMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := UsergreetingeventgreetingMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if OwnerType, ok := UsergreetingeventgreetingMap["ownerType"].(string); ok {
		o.OwnerType = &OwnerType
	}
    
	if Owner, ok := UsergreetingeventgreetingMap["owner"].(map[string]interface{}); ok {
		OwnerString, _ := json.Marshal(Owner)
		json.Unmarshal(OwnerString, &o.Owner)
	}
	
	if GreetingAudioFile, ok := UsergreetingeventgreetingMap["greetingAudioFile"].(map[string]interface{}); ok {
		GreetingAudioFileString, _ := json.Marshal(GreetingAudioFile)
		json.Unmarshal(GreetingAudioFileString, &o.GreetingAudioFile)
	}
	
	if AudioTTS, ok := UsergreetingeventgreetingMap["audioTTS"].(string); ok {
		o.AudioTTS = &AudioTTS
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Usergreetingeventgreeting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
