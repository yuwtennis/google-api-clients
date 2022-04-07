package main

import (
	"context"
	helpers "gihub.com/google-api-tutorials/internal"
	apiservice "gihub.com/google-api-tutorials/internal/factories"
	"log"
	"time"

	admin "google.golang.org/api/admin/reports/v1"
)

func main() {
	// https://pkg.go.dev/google.golang.org/api/admin/reports/v1
	// https://developers.google.com/admin-sdk/reports/v1/quickstart/go
	ctx := context.Background()
	duration, _ := time.ParseDuration("-96h")
	startTime := time.Now().UTC().Add(duration).Format(time.RFC3339)

	log.Printf("Initializing client...")
	creds := helpers.LoadDefaultCredentials(ctx, admin.AdminReportsAuditReadonlyScope)
	s := new(apiservice.AdminService)

	s.Create(ctx, creds)

	log.Printf("Prepare API object.")
	resp, err := s.Service.Activities.List("all", "drive").StartTime(startTime).Do()

	if err != nil {
		log.Fatalf("Failed retrieve result from api %v", err)
	}

	log.Printf("Items: %v", resp.Items)

	nextPageToken := &resp.NextPageToken
	count := 1

	for nextPageToken != nil {
		resp, err := s.Service.Activities.List("all", "drive").
			StartTime(startTime).
			PageToken(*nextPageToken).
			Do()

		log.Printf("Page: %v Items: %v", count, resp.Items)

		if err != nil {
			log.Fatalf("Failed to paginate result %v", err)
		}

		nextPageToken := &resp.NextPageToken

		if nextPageToken == nil {
			break
		}

		count++
	}
}
