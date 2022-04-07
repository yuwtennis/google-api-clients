package factories

import (
	"context"
	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/reports/v1"
	"google.golang.org/api/option"
	"log"
)

// SErvice is an interface of all google services
type Service interface {
	Create(ctx context.Context, creds *google.Credentials)
}

// AdminService holds Service object
type AdminService struct {
	Service *admin.Service
}

// Create return constructed ApiService including admin service
func (s *AdminService) Create(ctx context.Context, creds *google.Credentials) {
	service, err := admin.NewService(ctx, option.WithCredentials(creds))

	if err != nil {
		log.Fatalf("Failed to build service %v", err)
	}

	s.Service = service
}
