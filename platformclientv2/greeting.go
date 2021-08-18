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

func (u *Greeting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Greeting

	
	CreatedDate := new(string)
	if u.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(u.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		Name: u.Name,
		
		VarType: u.VarType,
		
		OwnerType: u.OwnerType,
		
		Owner: u.Owner,
		
		AudioFile: u.AudioFile,
		
		AudioTTS: u.AudioTTS,
		
		CreatedDate: CreatedDate,
		
		CreatedBy: u.CreatedBy,
		
		ModifiedDate: ModifiedDate,
		
		ModifiedBy: u.ModifiedBy,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Greeting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
