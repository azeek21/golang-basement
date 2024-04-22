package utils

import (
	"fmt"

	"server/types"
)

func GetFindClosestRestaurantQuery(
	pagination types.PagingIncoming,
	lat, lon float64,
	distance int,
) string {
	query := fmt.Sprintf(`
{
  "from": %d,
  "size": %d,
  "query": {
    "bool": {
      "must": {
        "match_all": {}
      },
      "filter": {
        "geo_distance": {
          "distance": "%vm",
          "location": {
            "lat": %v,
            "lon": %v
          }
        }
      }
    }
  },
  "sort": [
    {
      "_geo_distance": {
        "location": {
          "lat": %v,
          "lon": %v
        },
        "order": "asc",
        "unit": "km",
        "mode": "min",
        "distance_type": "arc",
        "ignore_unmapped": true
      }
    }
  ]
}
`, (pagination.PageNumber-1)*pagination.PageSize, pagination.PageSize, distance, lat, lon, lat, lon)
	return query
}
