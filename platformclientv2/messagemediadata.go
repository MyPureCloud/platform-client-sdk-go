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

func (u *Messagemediadata) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Name: u.Name,
		
		Url: u.Url,
		
		MediaType: u.MediaType,
		
		ContentLengthBytes: u.ContentLengthBytes,
		
		UploadUrl: u.UploadUrl,
		
		Status: u.Status,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Messagemediadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
