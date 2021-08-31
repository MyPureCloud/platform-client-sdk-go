package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Initiatescreenrecording
type Initiatescreenrecording struct { 
	// RecordACW
	RecordACW *bool `json:"recordACW,omitempty"`


	// ArchiveRetention
	ArchiveRetention *Archiveretention `json:"archiveRetention,omitempty"`


	// DeleteRetention
	DeleteRetention *Deleteretention `json:"deleteRetention,omitempty"`

}

func (o *Initiatescreenrecording) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Initiatescreenrecording
	
	return json.Marshal(&struct { 
		RecordACW *bool `json:"recordACW,omitempty"`
		
		ArchiveRetention *Archiveretention `json:"archiveRetention,omitempty"`
		
		DeleteRetention *Deleteretention `json:"deleteRetention,omitempty"`
		*Alias
	}{ 
		RecordACW: o.RecordACW,
		
		ArchiveRetention: o.ArchiveRetention,
		
		DeleteRetention: o.DeleteRetention,
		Alias:    (*Alias)(o),
	})
}

func (o *Initiatescreenrecording) UnmarshalJSON(b []byte) error {
	var InitiatescreenrecordingMap map[string]interface{}
	err := json.Unmarshal(b, &InitiatescreenrecordingMap)
	if err != nil {
		return err
	}
	
	if RecordACW, ok := InitiatescreenrecordingMap["recordACW"].(bool); ok {
		o.RecordACW = &RecordACW
	}
	
	if ArchiveRetention, ok := InitiatescreenrecordingMap["archiveRetention"].(map[string]interface{}); ok {
		ArchiveRetentionString, _ := json.Marshal(ArchiveRetention)
		json.Unmarshal(ArchiveRetentionString, &o.ArchiveRetention)
	}
	
	if DeleteRetention, ok := InitiatescreenrecordingMap["deleteRetention"].(map[string]interface{}); ok {
		DeleteRetentionString, _ := json.Marshal(DeleteRetention)
		json.Unmarshal(DeleteRetentionString, &o.DeleteRetention)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Initiatescreenrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
