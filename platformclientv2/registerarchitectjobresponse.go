package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Registerarchitectjobresponse
type Registerarchitectjobresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// PresignedUrl - Presigned URL to upload the file in S3
	PresignedUrl *string `json:"presignedUrl,omitempty"`


	// Headers - Required headers when uploading a file through PUT request to the URL
	Headers *map[string]string `json:"headers,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Registerarchitectjobresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Registerarchitectjobresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		PresignedUrl *string `json:"presignedUrl,omitempty"`
		
		Headers *map[string]string `json:"headers,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		PresignedUrl: o.PresignedUrl,
		
		Headers: o.Headers,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Registerarchitectjobresponse) UnmarshalJSON(b []byte) error {
	var RegisterarchitectjobresponseMap map[string]interface{}
	err := json.Unmarshal(b, &RegisterarchitectjobresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RegisterarchitectjobresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if PresignedUrl, ok := RegisterarchitectjobresponseMap["presignedUrl"].(string); ok {
		o.PresignedUrl = &PresignedUrl
	}
    
	if Headers, ok := RegisterarchitectjobresponseMap["headers"].(map[string]interface{}); ok {
		HeadersString, _ := json.Marshal(Headers)
		json.Unmarshal(HeadersString, &o.Headers)
	}
	
	if SelfUri, ok := RegisterarchitectjobresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Registerarchitectjobresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
