package usecase

import (
	"fmt"

	"github.com/dolginaa/url-shortener/internal/domain"
)

type StorageMock struct {
	getResultMap  map[string]StorageGetMockResult
	saveResultMap map[string]error
}

type StorageGetMockResult struct {
	originalURL domain.OriginalURL
	err         error
}

func NewStorageMock(
	resultMap map[string]StorageGetMockResult,
	saveResultMap map[string]error,
) StorageMock {
	return StorageMock{
		getResultMap:  resultMap,
		saveResultMap: saveResultMap,
	}
}

func (sm StorageMock) GetByShort(sURL domain.ShortenedURL) (domain.OriginalURL, error) {
	res, found := sm.getResultMap[sURL.ShortenedURL]
	if !found {
		return domain.OriginalURL{}, fmt.Errorf("error getting mock result")
	}

	return res.originalURL, res.err
}

func (sm StorageMock) Save(oURL domain.OriginalURL, sURL domain.ShortenedURL) error {
	resErr, found := sm.saveResultMap[oURL.OriginalURL]
	if !found {
		return fmt.Errorf("couldn't find save result in map")
	}

	return resErr
}
