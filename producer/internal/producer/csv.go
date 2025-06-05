package producer

import (
	"os"

	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
)

func readCsv(filePath string) ([]Apartment, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Errorf("failed to close file %s: %v", filePath, err)
		}
	}()

	var apartments []Apartment
	if err = gocsv.Unmarshal(file, &apartments); err != nil {
		return nil, err
	}

	return apartments, nil
}
