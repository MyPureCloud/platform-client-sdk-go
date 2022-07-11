package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeofflimit
type Timeofflimit struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Granularity - Granularity choice for the time off limit
	Granularity *string `json:"granularity,omitempty"`


	// DefaultLimitMinutes - The default time off limit value in minutes per granularity interval
	DefaultLimitMinutes *int `json:"defaultLimitMinutes,omitempty"`


	// Metadata - Version metadata for the time off limit
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Timeofflimit) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeofflimit
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		DefaultLimitMinutes *int `json:"defaultLimitMinutes,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Granularity: o.Granularity,
		
		DefaultLimitMinutes: o.DefaultLimitMinutes,
		
		Metadata: o.Metadata,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeofflimit) UnmarshalJSON(b []byte) error {
	var TimeofflimitMap map[string]interface{}
	err := json.Unmarshal(b, &TimeofflimitMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TimeofflimitMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Granularity, ok := TimeofflimitMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if DefaultLimitMinutes, ok := TimeofflimitMap["defaultLimitMinutes"].(float64); ok {
		DefaultLimitMinutesInt := int(DefaultLimitMinutes)
		o.DefaultLimitMinutes = &DefaultLimitMinutesInt
	}
	
	if Metadata, ok := TimeofflimitMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if SelfUri, ok := TimeofflimitMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Timeofflimit) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
