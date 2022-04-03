package main

import (
  "context"
  "log"
  "time"

  "golang.org/x/oauth2/google"
  admin "google.golang.org/api/admin/reports/v1"
  "google.golang.org/api/option"
)

func main() {
  // https://pkg.go.dev/google.golang.org/api/admin/reports/v1
  // https://developers.google.com/admin-sdk/reports/v1/quickstart/go
  ctx := context.Background()
  duration, _ := time.ParseDuration("-96h")

  client, err := google.DefaultClient(
    ctx,
    admin.AdminReportsAuditReadonlyScope)

  if err != nil {
    log.Fatalf("Client credential error %v", err)
  }

  src, err := admin.NewService(ctx, option.NewHTTPClient(client))

  if err != nil {
    log.Fatalf("Failed to build client using discovery service %v", err)
  }

  resp, err := src.
    Activities.List("all", "drive").
      StartTime(time.Now().UTC().Add(duration)).
      Do()

  if err != nil {
    log.Fatalf("Failed retrieve result from api %v", err)
  }

  log.Printf("Items: %v", resp.Items)

  next_page_token := resp.NextPageToken
  count := 1

  for next_page_token != nil {
    resp, err := src.Activities.List("all", "drive").
      StartTime(time.Now().UTC().Add(duration)).
      PageToken(next_page_token).
      Do()

    log.Printf("Page: %v Items: %v", count, resp.Items)

    next_page_token := resp.NextPageToken

    if next_page_token == nil { break }
    count++
  }
}
