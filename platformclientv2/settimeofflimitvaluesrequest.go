package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Settimeofflimitvaluesrequest
type Settimeofflimitvaluesrequest struct { 
	// Values
	Values *[]Timeofflimitrange `json:"values,omitempty"`


	// Metadata - Version metadata for the time off limit
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Settimeofflimitvaluesrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Settimeofflimitvaluesrequest
	
	return json.Marshal(&struct { 
		Values *[]Timeofflimitrange `json:"values,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Values: o.Values,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Settimeofflimitvaluesrequest) UnmarshalJSON(b []byte) error {
	var SettimeofflimitvaluesrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SettimeofflimitvaluesrequestMap)
	if err != nil {
		return err
	}
	
	if Values, ok := SettimeofflimitvaluesrequestMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if Metadata, ok := SettimeofflimitvaluesrequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Settimeofflimitvaluesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
