package entry

import "diary-app/internal/domain/entities"

func ToEntryModel(entry entities.Entry) *EntryModel {
	return &EntryModel{
		ID:    uint(entry.ID),
		Title: entry.Title,
		Text:  entry.Text,
		Date:  entry.Date,
	}
}

func ToEntry(em *EntryModel) *entities.Entry {
	return &entities.Entry{
		ID:    int(em.ID),
		Title: em.Title,
		Text:  em.Text,
		Date:  em.Date,
	}
}
