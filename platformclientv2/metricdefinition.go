package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Metricdefinition
type Metricdefinition struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// UnitType - The type of associated metric unit
	UnitType *string `json:"unitType,omitempty"`


	// ShortName - An alternate name for this metric definition, often abbreviation
	ShortName *string `json:"shortName,omitempty"`


	// DividendMetrics - Metric names used as dividend
	DividendMetrics *[]string `json:"dividendMetrics,omitempty"`


	// DivisorMetrics - Metric names used as divisor
	DivisorMetrics *[]string `json:"divisorMetrics,omitempty"`


	// DefaultObjective - A predefined default objective for this metric
	DefaultObjective *Defaultobjective `json:"defaultObjective,omitempty"`


	// LockTemplateId - An optional field to specify if this metric definition is locked to certain template. e.g. punctuality
	LockTemplateId *string `json:"lockTemplateId,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Metricdefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Metricdefinition
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		UnitType *string `json:"unitType,omitempty"`
		
		ShortName *string `json:"shortName,omitempty"`
		
		DividendMetrics *[]string `json:"dividendMetrics,omitempty"`
		
		DivisorMetrics *[]string `json:"divisorMetrics,omitempty"`
		
		DefaultObjective *Defaultobjective `json:"defaultObjective,omitempty"`
		
		LockTemplateId *string `json:"lockTemplateId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		UnitType: o.UnitType,
		
		ShortName: o.ShortName,
		
		DividendMetrics: o.DividendMetrics,
		
		DivisorMetrics: o.DivisorMetrics,
		
		DefaultObjective: o.DefaultObjective,
		
		LockTemplateId: o.LockTemplateId,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Metricdefinition) UnmarshalJSON(b []byte) error {
	var MetricdefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &MetricdefinitionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MetricdefinitionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := MetricdefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if UnitType, ok := MetricdefinitionMap["unitType"].(string); ok {
		o.UnitType = &UnitType
	}
	
	if ShortName, ok := MetricdefinitionMap["shortName"].(string); ok {
		o.ShortName = &ShortName
	}
	
	if DividendMetrics, ok := MetricdefinitionMap["dividendMetrics"].([]interface{}); ok {
		DividendMetricsString, _ := json.Marshal(DividendMetrics)
		json.Unmarshal(DividendMetricsString, &o.DividendMetrics)
	}
	
	if DivisorMetrics, ok := MetricdefinitionMap["divisorMetrics"].([]interface{}); ok {
		DivisorMetricsString, _ := json.Marshal(DivisorMetrics)
		json.Unmarshal(DivisorMetricsString, &o.DivisorMetrics)
	}
	
	if DefaultObjective, ok := MetricdefinitionMap["defaultObjective"].(map[string]interface{}); ok {
		DefaultObjectiveString, _ := json.Marshal(DefaultObjective)
		json.Unmarshal(DefaultObjectiveString, &o.DefaultObjective)
	}
	
	if LockTemplateId, ok := MetricdefinitionMap["lockTemplateId"].(string); ok {
		o.LockTemplateId = &LockTemplateId
	}
	
	if SelfUri, ok := MetricdefinitionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Metricdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
