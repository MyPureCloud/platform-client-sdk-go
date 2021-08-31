package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Retentionduration
type Retentionduration struct { 
	// ArchiveRetention
	ArchiveRetention *Archiveretention `json:"archiveRetention,omitempty"`


	// DeleteRetention
	DeleteRetention *Deleteretention `json:"deleteRetention,omitempty"`

}

func (o *Retentionduration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Retentionduration
	
	return json.Marshal(&struct { 
		ArchiveRetention *Archiveretention `json:"archiveRetention,omitempty"`
		
		DeleteRetention *Deleteretention `json:"deleteRetention,omitempty"`
		*Alias
	}{ 
		ArchiveRetention: o.ArchiveRetention,
		
		DeleteRetention: o.DeleteRetention,
		Alias:    (*Alias)(o),
	})
}

func (o *Retentionduration) UnmarshalJSON(b []byte) error {
	var RetentiondurationMap map[string]interface{}
	err := json.Unmarshal(b, &RetentiondurationMap)
	if err != nil {
		return err
	}
	
	if ArchiveRetention, ok := RetentiondurationMap["archiveRetention"].(map[string]interface{}); ok {
		ArchiveRetentionString, _ := json.Marshal(ArchiveRetention)
		json.Unmarshal(ArchiveRetentionString, &o.ArchiveRetention)
	}
	
	if DeleteRetention, ok := RetentiondurationMap["deleteRetention"].(map[string]interface{}); ok {
		DeleteRetentionString, _ := json.Marshal(DeleteRetention)
		json.Unmarshal(DeleteRetentionString, &o.DeleteRetention)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Retentionduration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
