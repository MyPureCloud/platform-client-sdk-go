package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadata
type Wfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadata struct { 
	// Version
	Version *int `json:"version,omitempty"`


	// ModifiedBy
	ModifiedBy *Wfmbushorttermforecastupdatecompletetopicuserreference `json:"modifiedBy,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`

}

func (o *Wfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadata
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Version *int `json:"version,omitempty"`
		
		ModifiedBy *Wfmbushorttermforecastupdatecompletetopicuserreference `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		*Alias
	}{ 
		Version: o.Version,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadata) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadataMap)
	if err != nil {
		return err
	}
	
	if Version, ok := WfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadataMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if ModifiedBy, ok := WfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadataMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := WfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadataMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastupdatecompletetopicwfmversionedentitymetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
