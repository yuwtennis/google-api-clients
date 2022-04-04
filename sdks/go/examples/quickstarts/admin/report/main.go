package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/oauth2/google"
	"golang.org/x/vuln/client"
	admin "google.golang.org/api/admin/reports/v1"
	"google.golang.org/api/option"
apsrvice
	"internal/factories/client"lint
	"internal/factories/apiservice"
)

func main() {
	// https://pkg.go.dev/google.golang.org/api/admin/reports/v1
	// https://developers.google.com/admin-sdk/reports/v1/quickstart/go
	ctx := context.Background()
	duration, _ := time.ParseDuration("-96h")
	starttime := time.Now().UTC().Add(duration).String()

	c := client.NewClient()
	s := new(apiservice.AdminService)

	client := c.Create(ctx,admin.AdminReportsAuditReadonlyScope)
	srv = s.Create(ctx, client.Client)

	resp, err := srv.
		Activities.List("all", "drive").
		StartTime(starttime).
		Do()

	if err != nil {
		log.Fatalf("Failed retrieve result from api %v", err)
	}

	log.Printf("Items: %v", resp.Items)

	next_page_token := &resp.NextPageToken
	count := 1

	for next_page_token != nil {
		resp, err := srv.Activities.List("all", "drive").
			StartTime(starttime).
			PageToken(*next_page_token).
			Do()

		log.Printf("Page: %v Items: %v", count, resp.Items)

		if err != nil {
			log.Fatalf("Failed to paginate result %v", err)
		}

		next_page_token := &resp.NextPageToken

		if next_page_token == nil {
			break
		}

		count++
	}
}
