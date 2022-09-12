package services

import (
	"context"
	"testing"
)

func TestSeaports_Run(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		s := Seaports{
			Store:    nil,
			Api:      nil,
			Importer: nil,
		}
		s.Run(context.TODO())
	})
}
