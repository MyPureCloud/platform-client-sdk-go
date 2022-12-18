package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyapp
type Journeyapp struct { 
	// Name - Name of the application (e.g. mybankingapp).
	Name *string `json:"name,omitempty"`


	// Namespace - Namespace of the application (e.g. com.genesys.bancodinero).
	Namespace *string `json:"namespace,omitempty"`


	// Version - Version of the application (e.g. 5.9.27).
	Version *string `json:"version,omitempty"`


	// BuildNumber - Build number of the application (e.g. 701).
	BuildNumber *string `json:"buildNumber,omitempty"`

}

func (o *Journeyapp) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyapp
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Namespace *string `json:"namespace,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		BuildNumber *string `json:"buildNumber,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Namespace: o.Namespace,
		
		Version: o.Version,
		
		BuildNumber: o.BuildNumber,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeyapp) UnmarshalJSON(b []byte) error {
	var JourneyappMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyappMap)
	if err != nil {
		return err
	}
	
	if Name, ok := JourneyappMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Namespace, ok := JourneyappMap["namespace"].(string); ok {
		o.Namespace = &Namespace
	}
    
	if Version, ok := JourneyappMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if BuildNumber, ok := JourneyappMap["buildNumber"].(string); ok {
		o.BuildNumber = &BuildNumber
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyapp) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
