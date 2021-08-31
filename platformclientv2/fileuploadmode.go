package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Fileuploadmode
type Fileuploadmode struct { 
	// FileTypes - A list of supported content types for uploading files
	FileTypes *[]string `json:"fileTypes,omitempty"`


	// MaxFileSizeKB - The maximum file size for file uploads in kilobytes. Default is 10240 (10 MB)
	MaxFileSizeKB *int `json:"maxFileSizeKB,omitempty"`

}

func (o *Fileuploadmode) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Fileuploadmode
	
	return json.Marshal(&struct { 
		FileTypes *[]string `json:"fileTypes,omitempty"`
		
		MaxFileSizeKB *int `json:"maxFileSizeKB,omitempty"`
		*Alias
	}{ 
		FileTypes: o.FileTypes,
		
		MaxFileSizeKB: o.MaxFileSizeKB,
		Alias:    (*Alias)(o),
	})
}

func (o *Fileuploadmode) UnmarshalJSON(b []byte) error {
	var FileuploadmodeMap map[string]interface{}
	err := json.Unmarshal(b, &FileuploadmodeMap)
	if err != nil {
		return err
	}
	
	if FileTypes, ok := FileuploadmodeMap["fileTypes"].([]interface{}); ok {
		FileTypesString, _ := json.Marshal(FileTypes)
		json.Unmarshal(FileTypesString, &o.FileTypes)
	}
	
	if MaxFileSizeKB, ok := FileuploadmodeMap["maxFileSizeKB"].(float64); ok {
		MaxFileSizeKBInt := int(MaxFileSizeKB)
		o.MaxFileSizeKB = &MaxFileSizeKBInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Fileuploadmode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
