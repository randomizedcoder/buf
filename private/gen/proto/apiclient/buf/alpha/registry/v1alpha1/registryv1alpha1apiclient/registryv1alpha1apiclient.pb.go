// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-apiclient. DO NOT EDIT.

package registryv1alpha1apiclient

import (
	context "context"
	registryv1alpha1api "github.com/bufbuild/buf/private/gen/proto/api/buf/alpha/registry/v1alpha1/registryv1alpha1api"
)

// Provider provides all the types in registryv1alpha1apiclient.
type Provider interface {
	AuditLogsServiceProvider
	AuthnServiceProvider
	AuthzServiceProvider
	DocServiceProvider
	DownloadServiceProvider
	GenerateServiceProvider
	ImageServiceProvider
	LocalResolveServiceProvider
	OrganizationServiceProvider
	OwnerServiceProvider
	PluginServiceProvider
	PushServiceProvider
	RecommendationServiceProvider
	ReferenceServiceProvider
	RepositoryBranchServiceProvider
	RepositoryCommitServiceProvider
	RepositoryServiceProvider
	RepositoryTagServiceProvider
	RepositoryTrackServiceProvider
	ResolveServiceProvider
	SearchServiceProvider
	TokenServiceProvider
	UserServiceProvider
}

// AuditLogsServiceProvider provides a client-side AuditLogsService for an address.
type AuditLogsServiceProvider interface {
	NewAuditLogsService(ctx context.Context, address string) (registryv1alpha1api.AuditLogsService, error)
}

// AuthnServiceProvider provides a client-side AuthnService for an address.
type AuthnServiceProvider interface {
	NewAuthnService(ctx context.Context, address string) (registryv1alpha1api.AuthnService, error)
}

// AuthzServiceProvider provides a client-side AuthzService for an address.
type AuthzServiceProvider interface {
	NewAuthzService(ctx context.Context, address string) (registryv1alpha1api.AuthzService, error)
}

// DocServiceProvider provides a client-side DocService for an address.
type DocServiceProvider interface {
	NewDocService(ctx context.Context, address string) (registryv1alpha1api.DocService, error)
}

// DownloadServiceProvider provides a client-side DownloadService for an address.
type DownloadServiceProvider interface {
	NewDownloadService(ctx context.Context, address string) (registryv1alpha1api.DownloadService, error)
}

// GenerateServiceProvider provides a client-side GenerateService for an address.
type GenerateServiceProvider interface {
	NewGenerateService(ctx context.Context, address string) (registryv1alpha1api.GenerateService, error)
}

// ImageServiceProvider provides a client-side ImageService for an address.
type ImageServiceProvider interface {
	NewImageService(ctx context.Context, address string) (registryv1alpha1api.ImageService, error)
}

// LocalResolveServiceProvider provides a client-side LocalResolveService for an address.
type LocalResolveServiceProvider interface {
	NewLocalResolveService(ctx context.Context, address string) (registryv1alpha1api.LocalResolveService, error)
}

// OrganizationServiceProvider provides a client-side OrganizationService for an address.
type OrganizationServiceProvider interface {
	NewOrganizationService(ctx context.Context, address string) (registryv1alpha1api.OrganizationService, error)
}

// OwnerServiceProvider provides a client-side OwnerService for an address.
type OwnerServiceProvider interface {
	NewOwnerService(ctx context.Context, address string) (registryv1alpha1api.OwnerService, error)
}

// PluginServiceProvider provides a client-side PluginService for an address.
type PluginServiceProvider interface {
	NewPluginService(ctx context.Context, address string) (registryv1alpha1api.PluginService, error)
}

// PushServiceProvider provides a client-side PushService for an address.
type PushServiceProvider interface {
	NewPushService(ctx context.Context, address string) (registryv1alpha1api.PushService, error)
}

// RecommendationServiceProvider provides a client-side RecommendationService for an address.
type RecommendationServiceProvider interface {
	NewRecommendationService(ctx context.Context, address string) (registryv1alpha1api.RecommendationService, error)
}

// ReferenceServiceProvider provides a client-side ReferenceService for an address.
type ReferenceServiceProvider interface {
	NewReferenceService(ctx context.Context, address string) (registryv1alpha1api.ReferenceService, error)
}

// RepositoryBranchServiceProvider provides a client-side RepositoryBranchService for an address.
type RepositoryBranchServiceProvider interface {
	NewRepositoryBranchService(ctx context.Context, address string) (registryv1alpha1api.RepositoryBranchService, error)
}

// RepositoryCommitServiceProvider provides a client-side RepositoryCommitService for an address.
type RepositoryCommitServiceProvider interface {
	NewRepositoryCommitService(ctx context.Context, address string) (registryv1alpha1api.RepositoryCommitService, error)
}

// RepositoryServiceProvider provides a client-side RepositoryService for an address.
type RepositoryServiceProvider interface {
	NewRepositoryService(ctx context.Context, address string) (registryv1alpha1api.RepositoryService, error)
}

// RepositoryTagServiceProvider provides a client-side RepositoryTagService for an address.
type RepositoryTagServiceProvider interface {
	NewRepositoryTagService(ctx context.Context, address string) (registryv1alpha1api.RepositoryTagService, error)
}

// RepositoryTrackServiceProvider provides a client-side RepositoryTrackService for an address.
type RepositoryTrackServiceProvider interface {
	NewRepositoryTrackService(ctx context.Context, address string) (registryv1alpha1api.RepositoryTrackService, error)
}

// ResolveServiceProvider provides a client-side ResolveService for an address.
type ResolveServiceProvider interface {
	NewResolveService(ctx context.Context, address string) (registryv1alpha1api.ResolveService, error)
}

// SearchServiceProvider provides a client-side SearchService for an address.
type SearchServiceProvider interface {
	NewSearchService(ctx context.Context, address string) (registryv1alpha1api.SearchService, error)
}

// TokenServiceProvider provides a client-side TokenService for an address.
type TokenServiceProvider interface {
	NewTokenService(ctx context.Context, address string) (registryv1alpha1api.TokenService, error)
}

// UserServiceProvider provides a client-side UserService for an address.
type UserServiceProvider interface {
	NewUserService(ctx context.Context, address string) (registryv1alpha1api.UserService, error)
}
