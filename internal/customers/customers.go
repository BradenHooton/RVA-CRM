package customers

import (
	"time"
	"rva_crm/internal/core"
	"github.com/google/uuid"
)

type Customer struct {
    core.BaseModel
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	CompanyName    string    `json:"company_name"`
	JobTitle       string    `json:"job_title"`

	// Status and Classification
    Status       CustomerStatus `json:"status"`
    CustomerType CustomerType   `json:"customer_type"`
    Source       string         `json:"source"` // How they found us
    
    // Relationships
    Addresses []Address `json:"addresses"`
    
    // Metadata
    Tags        []string `json:"tags"`
    CustomFields map[string]interface{} `json:"custom_fields"`
}

type CustomerStatus string
const (
    CustomerStatusActive    CustomerStatus = "active"
    CustomerStatusInactive  CustomerStatus = "inactive"
    CustomerStatusBlocked   CustomerStatus = "blocked"
)

type CustomerType string
const (
    CustomerTypeProspect CustomerType = "prospect"
    CustomerTypeLead     CustomerType = "lead"
    CustomerTypeActive   CustomerType = "active"
    CustomerTypeChurned  CustomerType = "churned"
)

type CustomerSegment struct {
    core.BaseModel
    Name string `json:"name"`
    Description string `json:"description"`
    Criteria map[string]any `json:"criteria"`
}


// Address represents customer addresses
type Address struct {
    core.BaseModel
    CustomerID   uuid.UUID   `json:"customer_id"`
    Type         AddressType `json:"type"`
    Street1      string      `json:"street1"`
    Street2      string      `json:"street2"`
    City         string      `json:"city"`
    State        string      `json:"state"`
    PostalCode   string      `json:"postal_code"`
    Country      string      `json:"country"`
    IsDefault    bool        `json:"is_default"`
}

type AddressType string
const (
    AddressTypeBilling  AddressType = "billing"
    AddressTypeShipping AddressType = "shipping"
)

// Opporunity represents a potential customer/deal/engagement/project/etc.
type Opportunity struct {
    core.BaseModel
    CustomerID uuid.UUID `json:"customer_id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Value float64 `json:"value"`
    Stage OpportunityStage `json:"stage"`
    Probability float64 `json:"probability"`
    ExpectedCloseDate time.Time `json:"expected_close_date"`
    ActualCloseDate time.Time `json:"actual_close_date"`
    Source string `json:"source"`
    Products []OpportunityProduct `json:"products"`
}

type OpportunityStage string 
    const (
        StageProspecting OpportunityStage = "prospecting"
        StageQualified OpportunityStage = "qualified"
        StageProposal OpportunityStage = "proposal"
        StageNegotiation OpportunityStage = "negotiation"
        StageClosed OpportunityStage = "closed"
        StageLost OpportunityStage = "lost"
    )

type OpportunityProduct string
const (
    OpportunityProductService OpportunityProduct = "service - recurring"
    OpportunityProductOneTime OpportunityProduct = "service - one-time"
    OpportunityProductTaxStrategy OpportunityProduct = "tax strategy"
    OpportunityProductDueDiligence OpportunityProduct = "due diligence"
    OpportunityProductEntityFormation OpportunityProduct = "entity formation"
    OpportunityProductOther OpportunityProduct = "other"
)

type Lead struct {
    core.BaseModel
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"`
    Phone string `json:"phone"`
    Company string `json:"company"`
    Source string `json:"source"`
    Status LeadStatus `json:"status"`
    Score int `json:"score"`
    AssignedTo uuid.UUID `json:"assigned_to"`
    CustomerID uuid.UUID `json:"customer_id"`
}

type LeadStatus string
const (
    LeadStatusNew LeadStatus = "new"
    LeadStatusContacted LeadStatus = "contacted"
    LeadStatusQualified LeadStatus = "qualified"
    LeadStatusLost LeadStatus = "lost"
    LeadStatusWon LeadStatus = "won"
)