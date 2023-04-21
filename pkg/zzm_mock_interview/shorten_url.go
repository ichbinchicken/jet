package zzm_mock_interview

import "fmt"

type URLShortener interface {
	RegisterNewShortURL(shortURL string, actualURL string, userID string) error
	MakeRequestWith(shortURL string) (string, error)
	GetURLHitsCountFor(userID string) (int, error)
	GetMostPopularActualURL() string
}

type ActualURLMetaData struct {
	URLValue string
	UserId   string
}

type URLShortenerImpl struct {
	ShortURLToActualURL map[string]*ActualURLMetaData
	UserIDToURLCount    map[string]int
	ActualURLToCount    map[string]int
}

func NewURLShortener() URLShortener {
	return &URLShortenerImpl{
		ShortURLToActualURL: make(map[string]*ActualURLMetaData),
		UserIDToURLCount:    make(map[string]int),
		ActualURLToCount:    make(map[string]int),
	}
}

func (u *URLShortenerImpl) RegisterNewShortURL(shortURL string, actualURL string, userID string) error {
	_, ok := u.ShortURLToActualURL[shortURL]

	// If the shortURL key exists, we don't want to continue.
	// To make the question simplified, we define the shortURL key is unique globally across all user IDs.
	// Potentially, for a scale-up question, we could define the shortURL key is unique for each user ID.
	if ok {
		return fmt.Errorf("shortURL key: %s has already existed", shortURL)
	} else {
		u.ShortURLToActualURL[shortURL] = &ActualURLMetaData{
			URLValue: actualURL,
			UserId:   userID,
		}
	}

	return nil
}

func (u *URLShortenerImpl) MakeRequestWith(shortURL string) (string, error) {
	metaData, ok := u.ShortURLToActualURL[shortURL]
	if !ok {
		return "", fmt.Errorf("shortURL key: %s does not exist", shortURL)
	}

	userID := metaData.UserId
	actualURL := metaData.URLValue
	u.UserIDToURLCount[userID] += 1
	u.ActualURLToCount[actualURL] += 1

	return metaData.URLValue, nil

}

func (u *URLShortenerImpl) GetURLHitsCountFor(userID string) (int, error) {
	if val, ok := u.UserIDToURLCount[userID]; ok {
		return val, nil
	} else {
		return -1, fmt.Errorf("userId: %s does not exist", userID)
	}
}

func (u *URLShortenerImpl) GetMostPopularActualURL() string {
	// what if there is more one max value?
	// what if all counts are zero = default value of max
	return "not implemented"
}
