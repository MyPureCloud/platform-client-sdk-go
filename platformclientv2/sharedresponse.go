package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Sharedresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sharedresponse

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DownloadUri *string `json:"downloadUri,omitempty"`
		
		ViewUri *string `json:"viewUri,omitempty"`
		
		Document *Document `json:"document,omitempty"`
		
		Share *Share `json:"share,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		DownloadUri: u.DownloadUri,
		
		ViewUri: u.ViewUri,
		
		Document: u.Document,
		
		Share: u.Share,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Sharedresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
