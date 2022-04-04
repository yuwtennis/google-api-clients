package factories

import (
	"context"
	"log"

	admin "google.golang.org/api/admin/reports/v1"
	"google.golang.org/api/option"
)

// APIService is an interface of all google services
type APIService interface {
	Create(ctx context.Context, client *Client)
}

// AdminService holds Service object
type AdminService struct {
	Service *admin.Service
}

// Create return constructed ApiService including admin service
func (s *AdminService) Create(ctx context.Context, client *Client) {
	service, err := admin.NewService(ctx, option.WithHTTPClient(client.Client))

	if err == nil {
		log.Fatalf("Failed to build service %v", err)
	}

	s.Service = service
}
