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

func (o *Sharedresponse) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		DownloadUri: o.DownloadUri,
		
		ViewUri: o.ViewUri,
		
		Document: o.Document,
		
		Share: o.Share,
		Alias:    (*Alias)(o),
	})
}

func (o *Sharedresponse) UnmarshalJSON(b []byte) error {
	var SharedresponseMap map[string]interface{}
	err := json.Unmarshal(b, &SharedresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SharedresponseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if DownloadUri, ok := SharedresponseMap["downloadUri"].(string); ok {
		o.DownloadUri = &DownloadUri
	}
	
	if ViewUri, ok := SharedresponseMap["viewUri"].(string); ok {
		o.ViewUri = &ViewUri
	}
	
	if Document, ok := SharedresponseMap["document"].(map[string]interface{}); ok {
		DocumentString, _ := json.Marshal(Document)
		json.Unmarshal(DocumentString, &o.Document)
	}
	
	if Share, ok := SharedresponseMap["share"].(map[string]interface{}); ok {
		ShareString, _ := json.Marshal(Share)
		json.Unmarshal(ShareString, &o.Share)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sharedresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
