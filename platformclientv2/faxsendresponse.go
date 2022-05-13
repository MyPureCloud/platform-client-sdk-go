package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Faxsendresponse
type Faxsendresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// UploadDestinationUri
	UploadDestinationUri *string `json:"uploadDestinationUri,omitempty"`


	// UploadMethodType
	UploadMethodType *string `json:"uploadMethodType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Faxsendresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Faxsendresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		UploadDestinationUri *string `json:"uploadDestinationUri,omitempty"`
		
		UploadMethodType *string `json:"uploadMethodType,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		UploadDestinationUri: o.UploadDestinationUri,
		
		UploadMethodType: o.UploadMethodType,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Faxsendresponse) UnmarshalJSON(b []byte) error {
	var FaxsendresponseMap map[string]interface{}
	err := json.Unmarshal(b, &FaxsendresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FaxsendresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FaxsendresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if UploadDestinationUri, ok := FaxsendresponseMap["uploadDestinationUri"].(string); ok {
		o.UploadDestinationUri = &UploadDestinationUri
	}
    
	if UploadMethodType, ok := FaxsendresponseMap["uploadMethodType"].(string); ok {
		o.UploadMethodType = &UploadMethodType
	}
    
	if SelfUri, ok := FaxsendresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Faxsendresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
