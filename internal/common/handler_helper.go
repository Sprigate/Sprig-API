package common

import (
	"encoding/json"
	"errors"
	// "log"
	"net/http"
	"sprig/internal/model"
	"strconv"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func ParseIntParam(raw string, def int) (int, error) {
	if raw == "" {
		return def, nil
	}

	return strconv.Atoi(raw)
}

func ParsePagination(r *http.Request) (model.PaginationParams, error) {
	q := r.URL.Query()

	page, err := ParseIntParam(q.Get("page"), 1)
	if err != nil {
		return model.PaginationParams{}, errors.New("Parameter page harus berupa angka")
	}

	limit, err := ParseIntParam(q.Get("limit"), 10)
	if err != nil {
		return model.PaginationParams{}, errors.New("Parameter limit harus berupa angka")
	}

	return model.PaginationParams{Page: page, Limit: limit}, nil
}

func ParseBudgetFilter(r *http.Request) (model.BudgetFilter, error) {
	q := r.URL.Query()

	// NOTES : Karna ini parameternya seperti contoh dibawah
	// api/v1/budgets?month=1&year=2026
	// yang dimana ini merupakan query parameter, bukan path parameter
	monthStr, yearStr := q.Get("month"), q.Get("year")

	// ---- Debug State ----
	// log.Printf("month\t: %s", monthStr)
	// log.Printf("year\t: %s", yearStr)

	// tanpa filter (parameter)
	if monthStr == "" || yearStr == "" {
		return model.BudgetFilter{}, nil
	}

	// jika filter (parameter) diisi
	month, err := strconv.ParseUint(monthStr, 10, 8)
	if err != nil {
		return model.BudgetFilter{}, errors.New("Parameter month harus berupa angka")
	}
	year, err := strconv.ParseUint(yearStr, 10, 16)
	if err != nil {
		return model.BudgetFilter{}, errors.New("Parameter year harus berupa angka")
	}

	// DEBUG
	// log.Printf("%v, (month) data type\t: %T", uint8(month), uint8(month))
	// log.Printf("%v, (year) data type\t: %T", uint16(year), uint16(year))

	// NOTE : Karna hasil parse dari strconv.ParseUint masih
	// bertipe uint64, jadi konversi lagi aja.
	// example : uint8(month). uint64 -> uint8
	return model.BudgetFilter{Month: uint8(month), Year: uint16(year)}, nil
}
