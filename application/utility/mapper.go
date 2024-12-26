package utility

import (
	"gin/application/usecase/poll/results"
	"gin/domain/entities"
)

// map utility.PaginatedResponse[entities.Poll] -> utility.PaginatedResponse[results.GetPollResult]
func MapPoll(src PaginatedResponse[entities.Poll]) PaginatedResponse[results.GetPollResult] {

	var mappedResults PaginatedResponse[results.GetPollResult]
	mappedResults.Page = src.Page
	mappedResults.PageSize = src.PageSize
	mappedResults.TotalCount = src.TotalCount
	mappedResults.TotalPages = src.TotalPages

	for _, poll := range src.Data {
		mappedResults.Data = append(mappedResults.Data, results.GetPollResult{
			ID:          poll.ID,
			Title:       poll.Title,
			Description: poll.Description,
			ExpiresAt:   poll.ExpiresAt,
			Categories: func() []struct {
				CategoryID   uint
				CategoryName string
				Votes        int
			} {
				catList := make([]struct {
					CategoryID   uint
					CategoryName string
					Votes        int
				}, 0, len(poll.Categories))

				for _, c := range poll.Categories {
					catList = append(catList, struct {
						CategoryID   uint
						CategoryName string
						Votes        int
					}{
						CategoryID:   c.ID,
						CategoryName: c.Name,
						Votes:        len(c.Votes),
					})
				}
				return catList
			}(),
		})
	}

	return mappedResults
}

func MapSinglePoll(src *entities.Poll) *results.GetPollResult {

	mappedResult := results.GetPollResult{
		ID:          src.ID,
		Title:       src.Title,
		Description: src.Description,
		ExpiresAt:   src.ExpiresAt,
		Categories: func() []struct {
			CategoryID   uint
			CategoryName string
			Votes        int
		} {
			catList := make([]struct {
				CategoryID   uint
				CategoryName string
				Votes        int
			}, 0, len(src.Categories))

			for _, c := range src.Categories {
				catList = append(catList, struct {
					CategoryID   uint
					CategoryName string
					Votes        int
				}{
					CategoryID:   c.ID,
					CategoryName: c.Name,
					Votes:        len(c.Votes),
				})
			}
			return catList
		}(),
	}

	return &mappedResult
}
