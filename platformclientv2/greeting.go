package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Greeting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Greeting
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		OwnerType *string `json:"ownerType,omitempty"`
		
		Owner *Domainentity `json:"owner,omitempty"`
		
		AudioFile *Greetingaudiofile `json:"audioFile,omitempty"`
		
		AudioTTS *string `json:"audioTTS,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		OwnerType: o.OwnerType,
		
		Owner: o.Owner,
		
		AudioFile: o.AudioFile,
		
		AudioTTS: o.AudioTTS,
		
		CreatedDate: CreatedDate,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedDate: ModifiedDate,
		
		ModifiedBy: o.ModifiedBy,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Greeting) UnmarshalJSON(b []byte) error {
	var GreetingMap map[string]interface{}
	err := json.Unmarshal(b, &GreetingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GreetingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := GreetingMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := GreetingMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if OwnerType, ok := GreetingMap["ownerType"].(string); ok {
		o.OwnerType = &OwnerType
	}
    
	if Owner, ok := GreetingMap["owner"].(map[string]interface{}); ok {
		OwnerString, _ := json.Marshal(Owner)
		json.Unmarshal(OwnerString, &o.Owner)
	}
	
	if AudioFile, ok := GreetingMap["audioFile"].(map[string]interface{}); ok {
		AudioFileString, _ := json.Marshal(AudioFile)
		json.Unmarshal(AudioFileString, &o.AudioFile)
	}
	
	if AudioTTS, ok := GreetingMap["audioTTS"].(string); ok {
		o.AudioTTS = &AudioTTS
	}
    
	if createdDateString, ok := GreetingMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if CreatedBy, ok := GreetingMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
    
	if modifiedDateString, ok := GreetingMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if ModifiedBy, ok := GreetingMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
    
	if SelfUri, ok := GreetingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Greeting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
