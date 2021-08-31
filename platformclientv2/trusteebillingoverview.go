package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trusteebillingoverview
type Trusteebillingoverview struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Organization - Organization
	Organization *Namedentity `json:"organization,omitempty"`


	// Currency - The currency type.
	Currency *string `json:"currency,omitempty"`


	// EnabledProducts - The charge short names for products enabled during the specified period.
	EnabledProducts *[]string `json:"enabledProducts,omitempty"`


	// SubscriptionType - The subscription type.
	SubscriptionType *string `json:"subscriptionType,omitempty"`


	// RampPeriodStartDate - Date-time the ramp period starts. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RampPeriodStartDate *time.Time `json:"rampPeriodStartDate,omitempty"`


	// RampPeriodEndDate - Date-time the ramp period ends. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RampPeriodEndDate *time.Time `json:"rampPeriodEndDate,omitempty"`


	// BillingPeriodStartDate - Date-time the billing period started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	BillingPeriodStartDate *time.Time `json:"billingPeriodStartDate,omitempty"`


	// BillingPeriodEndDate - Date-time the billing period ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	BillingPeriodEndDate *time.Time `json:"billingPeriodEndDate,omitempty"`


	// Usages - Usages for the specified period.
	Usages *[]Subscriptionoverviewusage `json:"usages,omitempty"`


	// ContractAmendmentDate - Date-time the contract was last amended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ContractAmendmentDate *time.Time `json:"contractAmendmentDate,omitempty"`


	// ContractEffectiveDate - Date-time the contract became effective. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ContractEffectiveDate *time.Time `json:"contractEffectiveDate,omitempty"`


	// ContractEndDate - Date-time the contract ends. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ContractEndDate *time.Time `json:"contractEndDate,omitempty"`


	// MinimumMonthlyAmount - Minimum amount that will be charged for the month
	MinimumMonthlyAmount *string `json:"minimumMonthlyAmount,omitempty"`


	// InRampPeriod
	InRampPeriod *bool `json:"inRampPeriod,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Trusteebillingoverview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trusteebillingoverview
	
	RampPeriodStartDate := new(string)
	if o.RampPeriodStartDate != nil {
		
		*RampPeriodStartDate = timeutil.Strftime(o.RampPeriodStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RampPeriodStartDate = nil
	}
	
	RampPeriodEndDate := new(string)
	if o.RampPeriodEndDate != nil {
		
		*RampPeriodEndDate = timeutil.Strftime(o.RampPeriodEndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RampPeriodEndDate = nil
	}
	
	BillingPeriodStartDate := new(string)
	if o.BillingPeriodStartDate != nil {
		
		*BillingPeriodStartDate = timeutil.Strftime(o.BillingPeriodStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		BillingPeriodStartDate = nil
	}
	
	BillingPeriodEndDate := new(string)
	if o.BillingPeriodEndDate != nil {
		
		*BillingPeriodEndDate = timeutil.Strftime(o.BillingPeriodEndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		BillingPeriodEndDate = nil
	}
	
	ContractAmendmentDate := new(string)
	if o.ContractAmendmentDate != nil {
		
		*ContractAmendmentDate = timeutil.Strftime(o.ContractAmendmentDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ContractAmendmentDate = nil
	}
	
	ContractEffectiveDate := new(string)
	if o.ContractEffectiveDate != nil {
		
		*ContractEffectiveDate = timeutil.Strftime(o.ContractEffectiveDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ContractEffectiveDate = nil
	}
	
	ContractEndDate := new(string)
	if o.ContractEndDate != nil {
		
		*ContractEndDate = timeutil.Strftime(o.ContractEndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ContractEndDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Organization *Namedentity `json:"organization,omitempty"`
		
		Currency *string `json:"currency,omitempty"`
		
		EnabledProducts *[]string `json:"enabledProducts,omitempty"`
		
		SubscriptionType *string `json:"subscriptionType,omitempty"`
		
		RampPeriodStartDate *string `json:"rampPeriodStartDate,omitempty"`
		
		RampPeriodEndDate *string `json:"rampPeriodEndDate,omitempty"`
		
		BillingPeriodStartDate *string `json:"billingPeriodStartDate,omitempty"`
		
		BillingPeriodEndDate *string `json:"billingPeriodEndDate,omitempty"`
		
		Usages *[]Subscriptionoverviewusage `json:"usages,omitempty"`
		
		ContractAmendmentDate *string `json:"contractAmendmentDate,omitempty"`
		
		ContractEffectiveDate *string `json:"contractEffectiveDate,omitempty"`
		
		ContractEndDate *string `json:"contractEndDate,omitempty"`
		
		MinimumMonthlyAmount *string `json:"minimumMonthlyAmount,omitempty"`
		
		InRampPeriod *bool `json:"inRampPeriod,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Organization: o.Organization,
		
		Currency: o.Currency,
		
		EnabledProducts: o.EnabledProducts,
		
		SubscriptionType: o.SubscriptionType,
		
		RampPeriodStartDate: RampPeriodStartDate,
		
		RampPeriodEndDate: RampPeriodEndDate,
		
		BillingPeriodStartDate: BillingPeriodStartDate,
		
		BillingPeriodEndDate: BillingPeriodEndDate,
		
		Usages: o.Usages,
		
		ContractAmendmentDate: ContractAmendmentDate,
		
		ContractEffectiveDate: ContractEffectiveDate,
		
		ContractEndDate: ContractEndDate,
		
		MinimumMonthlyAmount: o.MinimumMonthlyAmount,
		
		InRampPeriod: o.InRampPeriod,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Trusteebillingoverview) UnmarshalJSON(b []byte) error {
	var TrusteebillingoverviewMap map[string]interface{}
	err := json.Unmarshal(b, &TrusteebillingoverviewMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TrusteebillingoverviewMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := TrusteebillingoverviewMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Organization, ok := TrusteebillingoverviewMap["organization"].(map[string]interface{}); ok {
		OrganizationString, _ := json.Marshal(Organization)
		json.Unmarshal(OrganizationString, &o.Organization)
	}
	
	if Currency, ok := TrusteebillingoverviewMap["currency"].(string); ok {
		o.Currency = &Currency
	}
	
	if EnabledProducts, ok := TrusteebillingoverviewMap["enabledProducts"].([]interface{}); ok {
		EnabledProductsString, _ := json.Marshal(EnabledProducts)
		json.Unmarshal(EnabledProductsString, &o.EnabledProducts)
	}
	
	if SubscriptionType, ok := TrusteebillingoverviewMap["subscriptionType"].(string); ok {
		o.SubscriptionType = &SubscriptionType
	}
	
	if rampPeriodStartDateString, ok := TrusteebillingoverviewMap["rampPeriodStartDate"].(string); ok {
		RampPeriodStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", rampPeriodStartDateString)
		o.RampPeriodStartDate = &RampPeriodStartDate
	}
	
	if rampPeriodEndDateString, ok := TrusteebillingoverviewMap["rampPeriodEndDate"].(string); ok {
		RampPeriodEndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", rampPeriodEndDateString)
		o.RampPeriodEndDate = &RampPeriodEndDate
	}
	
	if billingPeriodStartDateString, ok := TrusteebillingoverviewMap["billingPeriodStartDate"].(string); ok {
		BillingPeriodStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", billingPeriodStartDateString)
		o.BillingPeriodStartDate = &BillingPeriodStartDate
	}
	
	if billingPeriodEndDateString, ok := TrusteebillingoverviewMap["billingPeriodEndDate"].(string); ok {
		BillingPeriodEndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", billingPeriodEndDateString)
		o.BillingPeriodEndDate = &BillingPeriodEndDate
	}
	
	if Usages, ok := TrusteebillingoverviewMap["usages"].([]interface{}); ok {
		UsagesString, _ := json.Marshal(Usages)
		json.Unmarshal(UsagesString, &o.Usages)
	}
	
	if contractAmendmentDateString, ok := TrusteebillingoverviewMap["contractAmendmentDate"].(string); ok {
		ContractAmendmentDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", contractAmendmentDateString)
		o.ContractAmendmentDate = &ContractAmendmentDate
	}
	
	if contractEffectiveDateString, ok := TrusteebillingoverviewMap["contractEffectiveDate"].(string); ok {
		ContractEffectiveDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", contractEffectiveDateString)
		o.ContractEffectiveDate = &ContractEffectiveDate
	}
	
	if contractEndDateString, ok := TrusteebillingoverviewMap["contractEndDate"].(string); ok {
		ContractEndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", contractEndDateString)
		o.ContractEndDate = &ContractEndDate
	}
	
	if MinimumMonthlyAmount, ok := TrusteebillingoverviewMap["minimumMonthlyAmount"].(string); ok {
		o.MinimumMonthlyAmount = &MinimumMonthlyAmount
	}
	
	if InRampPeriod, ok := TrusteebillingoverviewMap["inRampPeriod"].(bool); ok {
		o.InRampPeriod = &InRampPeriod
	}
	
	if SelfUri, ok := TrusteebillingoverviewMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trusteebillingoverview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
