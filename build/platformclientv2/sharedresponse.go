package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Sharedresponse
type Sharedresponse struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// DownloadUri
	DownloadUri *string `json:"downloadUri,omitempty"`


	// ViewUri
	ViewUri *string `json:"viewUri,omitempty"`


	// Document
	Document *Document `json:"document,omitempty"`


	// Share
	Share *Share `json:"share,omitempty"`

}

// String returns a JSON representation of the model
func (o *Sharedresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
