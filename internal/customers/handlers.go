package customers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type customerHandler struct {
	service CustomerService
}

type addressHandler struct {
	service AddressService
}

type opportunityHandler struct {
	service OpportunityService
}

func (h *customerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getCustomer(w, r)
	case http.MethodPost:
		h.createCustomer(w, r)
	case http.MethodPut:
		h.updateCustomer(w, r)
	case http.MethodDelete:
		h.deleteCustomer(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *addressHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAddress(w, r)
	case http.MethodPost:
		h.createAddress(w, r)
	case http.MethodPut:
		h.updateAddress(w, r)
	case http.MethodDelete:
		h.deleteAddress(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *opportunityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getOpportunity(w, r)
	case http.MethodPost:
		h.createOpportunity(w, r)
	case http.MethodPut:
		h.updateOpportunity(w, r)
	case http.MethodDelete:
		h.deleteOpportunity(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func NewCustomerHandler(service CustomerService) http.Handler {
	return &customerHandler{service: service}
}

func NewAddressHandler(service AddressService) http.Handler {
	return &addressHandler{service: service}
}

func NewOpportunityHandler(service OpportunityService) http.Handler {
	return &opportunityHandler{service: service}
}

func (h *customerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := r.URL.Query().Get("id")
	customer, err := h.service.GetCustomerByID(r.Context(), uuid.MustParse(customerID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func (h *customerHandler) createCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdCustomer, err := h.service.CreateCustomer(r.Context(), customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdCustomer)
}

func (h *customerHandler) updateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedCustomer, err := h.service.UpdateCustomer(r.Context(), customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedCustomer)
}

func (h *customerHandler) deleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := r.URL.Query().Get("id")
	err := h.service.DeleteCustomer(r.Context(), uuid.MustParse(customerID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer deleted successfully"})
}

func (h *addressHandler) getAddress(w http.ResponseWriter, r *http.Request) {
	addressID := r.URL.Query().Get("id")
	address, err := h.service.GetAddressByID(r.Context(), uuid.MustParse(addressID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(address)
}

func (h *addressHandler) createAddress(w http.ResponseWriter, r *http.Request) {
	var address Address
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdAddress, err := h.service.CreateAddress(r.Context(), address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdAddress)
}

func (h *addressHandler) updateAddress(w http.ResponseWriter, r *http.Request) {
	var address Address
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedAddress, err := h.service.UpdateAddress(r.Context(), address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedAddress)
}

func (h *addressHandler) deleteAddress(w http.ResponseWriter, r *http.Request) {
	addressID := r.URL.Query().Get("id")
	err := h.service.DeleteAddress(r.Context(), uuid.MustParse(addressID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Address deleted successfully"})
}

func (h *opportunityHandler) getOpportunity(w http.ResponseWriter, r *http.Request) {
	opportunityID := r.URL.Query().Get("id")
	opportunity, err := h.service.GetOpportunityByID(r.Context(), uuid.MustParse(opportunityID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(opportunity)
}

func (h *opportunityHandler) createOpportunity(w http.ResponseWriter, r *http.Request) {
	var opportunity Opportunity
	err := json.NewDecoder(r.Body).Decode(&opportunity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdOpportunity, err := h.service.CreateOpportunity(r.Context(), opportunity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdOpportunity)
}

func (h *opportunityHandler) updateOpportunity(w http.ResponseWriter, r *http.Request) {
	var opportunity Opportunity
	err := json.NewDecoder(r.Body).Decode(&opportunity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}	
	updatedOpportunity, err := h.service.UpdateOpportunity(r.Context(), opportunity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedOpportunity)
}

func (h *opportunityHandler) deleteOpportunity(w http.ResponseWriter, r *http.Request) {
	opportunityID := r.URL.Query().Get("id")
	err := h.service.DeleteOpportunity(r.Context(), uuid.MustParse(opportunityID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Opportunity deleted successfully"})
}
