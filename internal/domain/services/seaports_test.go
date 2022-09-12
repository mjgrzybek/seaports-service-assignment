package services

import (
	"context"
	"github.com/stretchr/testify/mock"
	"seaports-service-assignment/internal/application/safeExit"
	"seaports-service-assignment/mocks"
	"testing"
)

func TestSeaports_Run(t *testing.T) {
	t.Run("Sample", func(t *testing.T) {
		go readDoneChan()

		store := mocks.NewStore(t)
		api := mocks.NewApi(t)
		api.On("Start", mock.Anything).Return(nil)
		importer := mocks.NewImporter(t)

		s := Seaports{
			Store:    store,
			Api:      api,
			Importer: importer,
		}
		s.Run(context.TODO())
	})
}

func readDoneChan() {
	<-safeExit.Done
}
