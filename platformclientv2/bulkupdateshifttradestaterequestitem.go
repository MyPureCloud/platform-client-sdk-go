package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkupdateshifttradestaterequestitem
type Bulkupdateshifttradestaterequestitem struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// State - The new state to set on the shift trade
	State *string `json:"state,omitempty"`


	// Metadata - Version metadata for the shift trade
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Bulkupdateshifttradestaterequestitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkupdateshifttradestaterequestitem
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkupdateshifttradestaterequestitem) UnmarshalJSON(b []byte) error {
	var BulkupdateshifttradestaterequestitemMap map[string]interface{}
	err := json.Unmarshal(b, &BulkupdateshifttradestaterequestitemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BulkupdateshifttradestaterequestitemMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := BulkupdateshifttradestaterequestitemMap["state"].(string); ok {
		o.State = &State
	}
    
	if Metadata, ok := BulkupdateshifttradestaterequestitemMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkupdateshifttradestaterequestitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
