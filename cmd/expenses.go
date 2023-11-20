package main

import (
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/harshadmanglani/splitwise/jwt"
	"github.com/harshadmanglani/splitwise/models"
	"github.com/labstack/echo/v4"
)

func calculateAmount(splitMode string, value int, fullAmount int, size int) int {
	switch splitMode {
	case "EQUAL":
		return fullAmount / size
	case "PERCENTAGE":
		return value * fullAmount / 100
	case "AMOUNT":
		return value
	}
	return 0
}

func createExpense(ctx echo.Context) error {
	userId := ctx.Get("claims").(jwt.Claims).Subject
	var request models.CreateExpenseRequest

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	if err := validateRequest(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, okResp{err.Error()})
	}

	uu, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	var expense *models.Expense = &models.Expense{
		ExpenseId:   uu.String(),
		Title:       request.Title,
		Amount:      request.Amount,
		OwnerUserId: userId,
		SplitMode:   request.SplitMode,
	}

	// TODO: figure out how to run multiple queries together
	queries.InsertExpense.Get(&expense.Id, expense.ExpenseId, expense.Amount, expense.Title, expense.SplitMode, expense.OwnerUserId, nil)

	var splits []*models.Split = make([]*models.Split, len(request.Splits))

	for i, split := range request.Splits {
		settled := false
		if split.UserUuid == expense.OwnerUserId {
			settled = true
		}
		uu, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
		}
		amount := calculateAmount(expense.SplitMode, split.Value, expense.Amount, len(request.Splits))
		splits[i] = &models.Split{
			SplitId:      uu.String(),
			OwedByUserId: split.UserUuid,
			OwedToUserId: expense.OwnerUserId,
			Amount:       amount,
			ExpenseId:    expense.ExpenseId,
			Settled:      settled,
		}
		queries.InsertSplit.Get(&splits[i].Id, splits[i].ExpenseId, splits[i].Amount, splits[i].OwedByUserId, splits[i].OwedToUserId, splits[i].Settled, nil)
	}

	return ctx.JSON(http.StatusOK, okResp{expense})
}

func validateRequest(request *models.CreateExpenseRequest) error {
	totalAmount := request.Amount
	totalCalculatedAmount := 0
	for _, split := range request.Splits {
		totalCalculatedAmount += calculateAmount(request.SplitMode, split.Value, request.Amount, len(request.Splits))
	}
	if totalAmount != totalCalculatedAmount {
		return fmt.Errorf("INVALID_SPLIT_VALUES")
	}
	return nil
}
