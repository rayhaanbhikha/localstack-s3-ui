package main

import (
	"github.com/rayhaanbhikha/localstack-s3-ui/db"
	"github.com/rayhaanbhikha/localstack-s3-ui/s3"
)

func seed(db *db.DB) error {

	// parse file.
	s3Resources, err := s3.Parse("./recorded_api_calls.json")
	if err != nil {
		return err
	}

	for _, s3Resource := range s3Resources {

		if s3Resource.Name == "index.html" && len(s3Resource.ParentDirs) == 3 {
			s3.EmptyDirResources(s3Resource)
		}

		// if s3Resource.Type == "Bucket" {
		// 	switch s3Resource.Method {
		// 	case "PUT":
		// 		_, err := db.AddBucket(s3Resource.BucketName)
		// 		if err != nil {
		// 			log.Fatal(err)
		// 		}
		// 	case "DELETE":
		// 		_, err := db.DeleteBucket(s3Resource.BucketName)
		// 		if err != nil {
		// 			log.Fatal(err)
		// 		}
		// 	}
		// } else {
		// 	fmt.Println(s3Resource)
		// 	switch s3Resource.Method {
		// 	case "PUT":
		// 		_, err := db.AddResource(s3Resource)
		// 		if err != nil {
		// 			log.Fatal(err)
		// 		}
		// 	case "DELETE":
		// 		_, err := db.DeleteResource(s3Resource)
		// 		if err != nil {
		// 			log.Fatal(err)
		// 		}
		// 	}
		// }

		// fmt.Println(s3Resource)
	}

	return nil
}
