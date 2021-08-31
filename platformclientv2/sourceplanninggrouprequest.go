package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sourceplanninggrouprequest - Source planning group
type Sourceplanninggrouprequest struct { 
	// Id - The ID of the planning group
	Id *string `json:"id,omitempty"`


	// Metadata - Version metadata for the planning group
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Sourceplanninggrouprequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sourceplanninggrouprequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Sourceplanninggrouprequest) UnmarshalJSON(b []byte) error {
	var SourceplanninggrouprequestMap map[string]interface{}
	err := json.Unmarshal(b, &SourceplanninggrouprequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SourceplanninggrouprequestMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Metadata, ok := SourceplanninggrouprequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sourceplanninggrouprequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
