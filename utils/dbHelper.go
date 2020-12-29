package utils

import (
	"errors"
	"fmt"

	"github.com/Bendomey/task-assignment/graph/model"
	"github.com/Bendomey/task-assignment/models"
)

// GenerateQuery takes a loook at what is coming from client and then generates a sieve
func GenerateQuery(filter *model.GetUsersInput, pagination *model.Pagination) (*models.FilterQuery, error) {
	filterResult := models.FilterQuery{
		Limit:     "",
		Skip:      "",
		Order:     "ASC",
		OrderBy:   "created_at",
		Search:    "",
		DateRange: "",
	}

	if pagination.Limit != nil {
		filterResult.Limit = fmt.Sprintf("LIMIT %d", *pagination.Limit)
	}

	if pagination.Skip != nil {
		filterResult.Skip = fmt.Sprintf("OFFSET %d", *pagination.Skip)
	}

	if filter.OrderBy != nil {
		filterResult.OrderBy = string(*filter.OrderBy)
	}

	if filter.Order != nil {
		filterResult.Order = string(*filter.Order)
	}

	if filter.DateField != nil {
		if filter.DateRange == nil {
			return nil, errors.New("DateRange is required")
		}
		filterResult.DateRange = fmt.Sprintf(" (USER1.%s BETWEEN '%s' AND '%s') AND ", *filter.DateField, filter.DateRange.StartDate.Format("2006/01/02T15:04:05.00000"), filter.DateRange.EndDate.Format("2006/01/02T15:04:05.00000"))
	}

	if filter.Search != nil {
		if filter.SearchFields == nil {
			return nil, errors.New("At least one search field is required")
		}
		searchFields := ""

		for i := 0; i < len(filter.SearchFields); i++ {
			searchFields += fmt.Sprintf(" USER1.%s LIKE '%%%s%%' %s", filter.SearchFields[i], *filter.Search, getLast(i, len(filter.SearchFields)))
		}

		filterResult.Search = fmt.Sprintf("(%s) AND", searchFields)
	}
	return &filterResult, nil
}

func getLast(i int, total int) string {

	if i == total-1 {
		return ""
	}
	return "OR"

}
