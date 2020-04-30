package main

import (
	"fmt"
	"strings"

	"github.com/rayhaanbhikha/localstack-s3-ui/api"
	"github.com/rayhaanbhikha/localstack-s3-ui/db"
)

func seed(db *db.DB) error {

	// parse file.
	apiRequests, err := api.Parse("./recorded_api_calls.json")
	if err != nil {
		return err
	}

	for _, apiRequest := range apiRequests {

		splitFn := func(c rune) bool {
			return c == '/'
		}
		path := strings.FieldsFunc(apiRequest.Path, splitFn)

		if len(path) == 1 {
			res, err := db.AddBucket(path[0])
			if err != nil {
				return err
			}
			fmt.Println(res.LastInsertId())
		}
		// fmt.Println(apiRequest)
	}


	return nil
}