package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagemediadata
type Messagemediadata struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Url - The location of the media, useful for retrieving it
	Url *string `json:"url,omitempty"`


	// MediaType - The detected internet media type of the the media object.  If null then the media type should be dictated by the url.
	MediaType *string `json:"mediaType,omitempty"`


	// ContentLengthBytes - The optional content length of the the media object, in bytes.
	ContentLengthBytes *int `json:"contentLengthBytes,omitempty"`


	// UploadUrl - The URL returned to upload an attachment
	UploadUrl *string `json:"uploadUrl,omitempty"`


	// Status - The status of the media, indicates if the media is in the process of uploading. If the upload fails, the media becomes invalid
	Status *string `json:"status,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Messagemediadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagemediadata
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		ContentLengthBytes *int `json:"contentLengthBytes,omitempty"`
		
		UploadUrl *string `json:"uploadUrl,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Url: o.Url,
		
		MediaType: o.MediaType,
		
		ContentLengthBytes: o.ContentLengthBytes,
		
		UploadUrl: o.UploadUrl,
		
		Status: o.Status,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagemediadata) UnmarshalJSON(b []byte) error {
	var MessagemediadataMap map[string]interface{}
	err := json.Unmarshal(b, &MessagemediadataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MessagemediadataMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := MessagemediadataMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Url, ok := MessagemediadataMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if MediaType, ok := MessagemediadataMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ContentLengthBytes, ok := MessagemediadataMap["contentLengthBytes"].(float64); ok {
		ContentLengthBytesInt := int(ContentLengthBytes)
		o.ContentLengthBytes = &ContentLengthBytesInt
	}
	
	if UploadUrl, ok := MessagemediadataMap["uploadUrl"].(string); ok {
		o.UploadUrl = &UploadUrl
	}
    
	if Status, ok := MessagemediadataMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if SelfUri, ok := MessagemediadataMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Messagemediadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
