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

func (o *Groupgreetingeventgreeting) MarshalJSON() ([]byte, error) {
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

func (o *Groupgreetingeventgreeting) UnmarshalJSON(b []byte) error {
	var GroupgreetingeventgreetingMap map[string]interface{}
	err := json.Unmarshal(b, &GroupgreetingeventgreetingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GroupgreetingeventgreetingMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := GroupgreetingeventgreetingMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if VarType, ok := GroupgreetingeventgreetingMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if OwnerType, ok := GroupgreetingeventgreetingMap["ownerType"].(string); ok {
		o.OwnerType = &OwnerType
	}
	
	if Owner, ok := GroupgreetingeventgreetingMap["owner"].(map[string]interface{}); ok {
		OwnerString, _ := json.Marshal(Owner)
		json.Unmarshal(OwnerString, &o.Owner)
	}
	
	if GreetingAudioFile, ok := GroupgreetingeventgreetingMap["greetingAudioFile"].(map[string]interface{}); ok {
		GreetingAudioFileString, _ := json.Marshal(GreetingAudioFile)
		json.Unmarshal(GreetingAudioFileString, &o.GreetingAudioFile)
	}
	
	if AudioTTS, ok := GroupgreetingeventgreetingMap["audioTTS"].(string); ok {
		o.AudioTTS = &AudioTTS
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Groupgreetingeventgreeting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
