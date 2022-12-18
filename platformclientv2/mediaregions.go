package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediaregions
type Mediaregions struct { 
	// AwsHomeRegion - The AWS region your organization is in.
	AwsHomeRegion *string `json:"awsHomeRegion,omitempty"`


	// AwsCoreRegions - The list of AWS regions to which Genesys Cloud is deployed with full functionality including media streaming.
	AwsCoreRegions *[]string `json:"awsCoreRegions,omitempty"`


	// AwsSatelliteRegions - The list of AWS regions that Genesys Cloud uses only for media streaming.
	AwsSatelliteRegions *[]string `json:"awsSatelliteRegions,omitempty"`

}

func (o *Mediaregions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediaregions
	
	return json.Marshal(&struct { 
		AwsHomeRegion *string `json:"awsHomeRegion,omitempty"`
		
		AwsCoreRegions *[]string `json:"awsCoreRegions,omitempty"`
		
		AwsSatelliteRegions *[]string `json:"awsSatelliteRegions,omitempty"`
		*Alias
	}{ 
		AwsHomeRegion: o.AwsHomeRegion,
		
		AwsCoreRegions: o.AwsCoreRegions,
		
		AwsSatelliteRegions: o.AwsSatelliteRegions,
		Alias:    (*Alias)(o),
	})
}

func (o *Mediaregions) UnmarshalJSON(b []byte) error {
	var MediaregionsMap map[string]interface{}
	err := json.Unmarshal(b, &MediaregionsMap)
	if err != nil {
		return err
	}
	
	if AwsHomeRegion, ok := MediaregionsMap["awsHomeRegion"].(string); ok {
		o.AwsHomeRegion = &AwsHomeRegion
	}
    
	if AwsCoreRegions, ok := MediaregionsMap["awsCoreRegions"].([]interface{}); ok {
		AwsCoreRegionsString, _ := json.Marshal(AwsCoreRegions)
		json.Unmarshal(AwsCoreRegionsString, &o.AwsCoreRegions)
	}
	
	if AwsSatelliteRegions, ok := MediaregionsMap["awsSatelliteRegions"].([]interface{}); ok {
		AwsSatelliteRegionsString, _ := json.Marshal(AwsSatelliteRegions)
		json.Unmarshal(AwsSatelliteRegionsString, &o.AwsSatelliteRegions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mediaregions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
