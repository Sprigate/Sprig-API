package budgets

import (
	"errors"
	"sprig/internal/model"
)

type BudgetService struct {
	repository BudgetsRepository
}

func NewBudgetService(repository BudgetsRepository) *BudgetService {
	return &BudgetService{
		repository: repository,
	}
}

func (s *BudgetService) CreateBudget(req CreateBudgetRequest) (*model.Budget, error) {
	if req.TotalIncome <= 10000 {
		return nil, errors.New("Nilai pemasukan tidak boleh kurang dari Rp 10.000")
	}

	if req.NeedsPercentage <= 0 {
		return nil, errors.New("Persentase kebutuhan tidak boleh kosong")
	}

	if req.WantsPercentage <= 0 {
		return nil, errors.New("Persentase keinginan tidak boleh kosong")
	}

	// NOTES : Pertimbangkan kembali, apakah savings bisa
	// kosong atau tidak?

	budget := &model.Budget{
		UserID:            req.UserID,
		TotalIncome:       req.TotalIncome,
		NeedsPercentage:   req.NeedsPercentage,
		WantsPercentage:   req.WantsPercentage,
		SavingsPercentage: req.SavingsPercentage, // ?
		Month:             req.Month,
		Year:              req.Year,
	}

	if err := s.repository.AddBudget(budget); err != nil {
		return nil, err
	}

	return budget, nil
}

func (s *BudgetService) UpdateBudget(id uint64, req CreateBudgetUpdateRequest) (*model.Budget, error) {
	existingBudget, err := s.repository.FindByID(id)

	// log.Printf("budget raw value(req)\t: %d\n data type\t: %T", *req.TotalIncome, *req.TotalIncome)
	// log.Printf("budget value(db)\t: %d", existingBudget.TotalIncome)

	if err != nil {
		return nil, err
	}

	// anjir 6 turunan loh ya, next time cari cara
	// biar bisa dpersingkat lagi (btw struct kalau di loop gabisa)
	if req.TotalIncome != nil && *req.TotalIncome != 0 {

		// validasi
		if *req.TotalIncome <= 10000 {
			return nil, errors.New("Nilai pemasukkan tidak boleh kurang dari Rp 10.000")
		}

		existingBudget.TotalIncome = *req.TotalIncome
	}

	if req.NeedsPercentage != nil && *req.NeedsPercentage != 0 {
		existingBudget.NeedsPercentage = *req.NeedsPercentage
	}

	if req.WantsPercentage != nil && *req.WantsPercentage != 0 {
		existingBudget.WantsPercentage = *req.WantsPercentage
	}

	if req.SavingsPercentage != nil && *req.SavingsPercentage != 0 {
		existingBudget.SavingsPercentage = *req.SavingsPercentage
	}

	// karna malas ngetik ulang, untuk bagian ini jadi
	// digabung saja
	if req.Month != nil && req.Year != nil {
		if *req.Month != 0 && *req.Year != 0 {
			existingBudget.Month = *req.Month
			existingBudget.Year = *req.Year
		}
	}

	total := existingBudget.NeedsPercentage + existingBudget.WantsPercentage + existingBudget.SavingsPercentage
	if total != 100 {
		return nil, errors.New("Total persentase (needs+wants+savings) harus bernilai 100")
	}

	if err = s.repository.Save(existingBudget); err != nil {
		return nil, err
	}

	return existingBudget, nil
}

func (s *BudgetService) DeleteBudget(id uint64) error {
	if id <= 0 {
		return errors.New("ID tidak boleh kosong")
	}

	if err := s.repository.DeleteByID(id); err != nil {
		return err
	}

	return nil
}

func (s *BudgetService) FindAllBudget() ([]model.Budget, error) {
	budgets, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return budgets, nil
}

func (s *BudgetService) FilterByMonthAndYear(month uint8, year uint16) ([]model.Budget, error) {
	budgets, err := s.repository.FilterByMonthAndYear(month, year)
	if err != nil {
		return nil, err
	}

	return budgets, nil
}
