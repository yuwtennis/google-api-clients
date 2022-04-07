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

	var nextPageToken *string = nil

	for {
		req := s.Service.Activities.List("all", "drive").
			StartTime(startTime)

		if nextPageToken != nil {
			req = req.PageToken(*nextPageToken)
		}

		resp, err := req.Do()

		if err != nil {
			log.Fatalf("Failed to paginate result: %v", err)
			break
		}

		log.Printf("Page: %v , Received %v items", pageCnt, len(resp.Items))
		for _, v := range resp.Items {
			log.Printf("Activity Record: Id.Time: %v , Id.ApplicationName: %v, Id.UniqueQualifier: %v, Num of Events: %v",
				v.Id.Time,
				v.Id.ApplicationName,
				v.Id.UniqueQualifier,
				len(v.Events))
		}

		if resp.NextPageToken == "" {
			break
		}

		pageCnt++
		*nextPageToken = resp.NextPageToken

		time.Sleep(5)
	}
}
