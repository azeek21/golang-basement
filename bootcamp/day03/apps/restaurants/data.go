package restaurants

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/google/uuid"
)

type RestaurantLocation struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Restaurant struct {
	Id       string             `json:"id"`
	Name     string             `json:"name"     csv:"Name"`
	Address  string             `json:"address"  csv:"Address"`
	Phone    string             `json:"phone"    csv:"Phone"`
	Location RestaurantLocation `json:"location"`
}

type RestaurantReader struct {
	cvsReader *csv.Reader
	fid       *os.File
}

// 0:id 1:name 2:address 3:phone 4:Longitude 5:Latitude
func (restaurantReader *RestaurantReader) Read(tar *Restaurant) error {
	slices, err := restaurantReader.cvsReader.Read()
	if nil != err {
		return err
	}
	rid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	tar.Id = rid.String()
	tar.Name = slices[1]
	tar.Address = slices[2]
	tar.Phone = slices[3]
	tar.Location.Lat, err = strconv.ParseFloat(slices[4], 64)
	if err != nil {
		return err
	}
	tar.Location.Lon, err = strconv.ParseFloat(slices[5], 64)
	if err != nil {
		return err
	}
	return nil
}

func newRestaurantsReader(path string) (*RestaurantReader, error) {
	reader := RestaurantReader{}
	fid, err := os.Open(path)
	if err != nil {
		return &reader, err
	}
	reader.fid = fid
	scvReader := csv.NewReader(fid)
	scvReader.Comma = rune('\t')
	reader.cvsReader = scvReader
	scvReader.Read()
	return &reader, nil
}

func loadRestaurants(client *elastic.Client, filePath string, indexName string) error {
	ctx := context.Background()
	reader, err := newRestaurantsReader(filePath)
	defer reader.fid.Close()
	if err != nil {
		return err
	}
	restaurant := Restaurant{}
	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      indexName,
		Client:     client,
		NumWorkers: 5,
	})
	if err != nil {
		return err
	}

	for {
		err := reader.Read(&restaurant)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		body, err := json.Marshal(restaurant)
		if err != nil {
			return err
		}

		err = bulkIndexer.Add(ctx, esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: restaurant.Id,
			Body:       bytes.NewReader(body),
		})
	}
	if err != nil {
		return err
	}
	err = bulkIndexer.Close(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("+ Indexed %d documents\n", bulkIndexer.Stats().NumIndexed)
	return nil
}
