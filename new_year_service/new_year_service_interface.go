package new_year_service

type NewYearService interface {
	GetRandomImageURL() string
	GetRandomMessage() NewYearMessage
}
