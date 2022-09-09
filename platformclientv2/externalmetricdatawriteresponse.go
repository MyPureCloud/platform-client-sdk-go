package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalmetricdatawriteresponse - External metric data write response
type Externalmetricdatawriteresponse struct { 
	// ProcessedEntities - The list of processed entities
	ProcessedEntities *[]Externalmetricdataprocesseditem `json:"processedEntities,omitempty"`


	// UnprocessedEntities - The list of unprocessed entities
	UnprocessedEntities *[]Externalmetricdataunprocesseditem `json:"unprocessedEntities,omitempty"`

}

func (o *Externalmetricdatawriteresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalmetricdatawriteresponse
	
	return json.Marshal(&struct { 
		ProcessedEntities *[]Externalmetricdataprocesseditem `json:"processedEntities,omitempty"`
		
		UnprocessedEntities *[]Externalmetricdataunprocesseditem `json:"unprocessedEntities,omitempty"`
		*Alias
	}{ 
		ProcessedEntities: o.ProcessedEntities,
		
		UnprocessedEntities: o.UnprocessedEntities,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalmetricdatawriteresponse) UnmarshalJSON(b []byte) error {
	var ExternalmetricdatawriteresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalmetricdatawriteresponseMap)
	if err != nil {
		return err
	}
	
	if ProcessedEntities, ok := ExternalmetricdatawriteresponseMap["processedEntities"].([]interface{}); ok {
		ProcessedEntitiesString, _ := json.Marshal(ProcessedEntities)
		json.Unmarshal(ProcessedEntitiesString, &o.ProcessedEntities)
	}
	
	if UnprocessedEntities, ok := ExternalmetricdatawriteresponseMap["unprocessedEntities"].([]interface{}); ok {
		UnprocessedEntitiesString, _ := json.Marshal(UnprocessedEntities)
		json.Unmarshal(UnprocessedEntitiesString, &o.UnprocessedEntities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalmetricdatawriteresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
