package main

import (
	"context"
	"gihub.com/google-api-tutorials/internal/factories"
	admin "google.golang.org/api/admin/reports/v1"
	"log"
	"os"
	"time"
)

func main() {
	// https://pkg.go.dev/google.golang.org/api/admin/reports/v1
	// https://developers.google.com/admin-sdk/reports/v1/quickstart/go
	ctx := context.Background()

	pageCnt := 1
	duration, _ := time.ParseDuration("-24h")
	startTime := time.Now().UTC().Add(duration).Format(time.RFC3339)

	scopes := []string{admin.AdminReportsAuditReadonlyScope}
	subjectEmail := os.Getenv("GOOGLE_SUBJECT_EMAIL")

	log.Printf("Initializing client...")
	c := factories.NewCredential()
	c.Create(ctx, scopes, subjectEmail)

	s := new(factories.AdminService)
	s.Create(ctx, c.Credential)

	log.Printf("Prepare API object.")
	resp, err := s.Service.Activities.List("all", "drive").
		StartTime(startTime).Do()

	if err != nil {
		log.Fatalf("Failed retrieve result from api %v", err)
	}

	log.Printf("Page: #{pageCnt} , Received #{len(resp.Items)} events")
	for _, v := range resp.Items {
		log.Printf("Record: %v", v)
	}

	nextPageToken := &resp.NextPageToken

	for nextPageToken != nil {
		pageCnt++

		resp, err := s.Service.Activities.List("all", "drive").
			StartTime(startTime).
			PageToken(*nextPageToken).
			Do()

		log.Printf("Page: #{pageCnt} , Received #{len(resp.Items)} events")
		for _, v := range resp.Items {
			log.Printf("Record: %v", v)
		}

		if err != nil {
			log.Fatalf("Failed to paginate result %v", err)
		}

		nextPageToken := &resp.NextPageToken

		if nextPageToken == nil {
			break
		}
	}
}
