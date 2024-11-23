package main

import (
	"ImpatientOrderSystem/internal/database"
	"encoding/json"
	"net/http"
	"time"
)

type MedicationOrder struct {
	OrderNumber     int       `json:"order_number"`
	FileNumber      int       `json:"file_number"`
	NurseName       string    `json:"nurse_name"`
	Ward            string    `json:"ward"`
	Bed             string    `json:"bed"`
	Medication      string    `json:"medication"`
	Uom             string    `json:"uom"`
	RequestTime     time.Time `json:"request_time"`
	NurseRemarks    string    `json:"nurse_remarks"`
	Status          string    `json:"status"`
	PharmacyRemarks string    `json:"pharmacy_remarks"`
}

func (cfg *config) handlerMedicationOrderCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		FileNumber      int       `json:"file_number"`
		NurseName       string    `json:"nurse_name"`
		Ward            string    `json:"ward"`
		Bed             string    `json:"bed"`
		Medication      string    `json:"medication"`
		Uom             string    `json:"uom"`
		RequestTime     time.Time `json:"request_time"`
		NurseRemarks    string    `json:"nurse_remarks"`
		Status          string    `json:"status"`
		PharmacyRemarks string    `json:"pharmacy_remarks"`
	}
	var err error

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Cannot decode parameters", err)
		return
	}

}

func (cfg *config) handlerMedicationOrderList(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Data []MedicationOrder `json:"data"`
	}

	medicationOrderListDB, err := cfg.db.GetMedicationOrderList(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Cannot get medication order list from database", err)
		return
	}

	// transform from []database.MedicationOrder to []MedicationOrder
	medicationOrderList := make([]MedicationOrder, len(medicationOrderListDB))
	for i, medicationOrder := range medicationOrderListDB {
		medicationOrderList[i] = MedicationOrder{
			OrderNumber:     int(medicationOrder.OrderNumber),
			FileNumber:      int(medicationOrder.FileNumber),
			NurseName:       medicationOrder.NurseName.String,
			Ward:            medicationOrder.Ward.String,
			Bed:             medicationOrder.Bed.String,
			Medication:      medicationOrder.Medication.String,
			Uom:             medicationOrder.Uom.String,
			RequestTime:     medicationOrder.RequestTime,
			NurseRemarks:    medicationOrder.NurseRemarks.String,
			Status:          medicationOrder.Status,
			PharmacyRemarks: medicationOrder.PharmacyRemarks.String,
		}
	}

	respondWithJSON(w, http.StatusOK, response{
		Data: medicationOrderList,
	})
}

func (cfg *config) handlerDispense(w http.ResponseWriter, r *http.Request) {

	UpdateMedicationOrder := database.MedicationOrder{}

	err := cfg.db.UpdateMedicationOrder(r.Context(), UpdateMedicationOrder.OrderNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}
