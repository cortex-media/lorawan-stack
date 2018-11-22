// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

var _ClientFieldPaths = [...]string{
	"attributes",
	"contact_info",
	"created_at",
	"description",
	"endorsed",
	"grants",
	"ids",
	"ids.client_id",
	"name",
	"redirect_uris",
	"rights",
	"secret",
	"skip_authorization",
	"state",
	"updated_at",
}

func (*Client) FieldMaskPaths() []string {
	ret := make([]string, len(_ClientFieldPaths))
	copy(ret, _ClientFieldPaths[:])
	return ret
}

func (dst *Client) SetFields(src *Client, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "attributes":
			dst.Attributes = src.Attributes
		case "contact_info":
			dst.ContactInfo = src.ContactInfo
		case "created_at":
			dst.CreatedAt = src.CreatedAt
		case "description":
			dst.Description = src.Description
		case "endorsed":
			dst.Endorsed = src.Endorsed
		case "grants":
			dst.Grants = src.Grants
		case "ids":
			dst.ClientIdentifiers = src.ClientIdentifiers
		case "ids.client_id":
			dst.ClientIdentifiers.SetFields(&src.ClientIdentifiers, _pathsWithoutPrefix("ids", paths)...)
		case "name":
			dst.Name = src.Name
		case "redirect_uris":
			dst.RedirectURIs = src.RedirectURIs
		case "rights":
			dst.Rights = src.Rights
		case "secret":
			dst.Secret = src.Secret
		case "skip_authorization":
			dst.SkipAuthorization = src.SkipAuthorization
		case "state":
			dst.State = src.State
		case "updated_at":
			dst.UpdatedAt = src.UpdatedAt
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _ClientsFieldPaths = [...]string{
	"clients",
}

func (*Clients) FieldMaskPaths() []string {
	ret := make([]string, len(_ClientsFieldPaths))
	copy(ret, _ClientsFieldPaths[:])
	return ret
}

func (dst *Clients) SetFields(src *Clients, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "clients":
			dst.Clients = src.Clients
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _GetClientRequestFieldPaths = [...]string{
	"client_ids",
	"client_ids.client_id",
	"field_mask",
}

func (*GetClientRequest) FieldMaskPaths() []string {
	ret := make([]string, len(_GetClientRequestFieldPaths))
	copy(ret, _GetClientRequestFieldPaths[:])
	return ret
}

func (dst *GetClientRequest) SetFields(src *GetClientRequest, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "client_ids":
			dst.ClientIdentifiers = src.ClientIdentifiers
		case "client_ids.client_id":
			dst.ClientIdentifiers.SetFields(&src.ClientIdentifiers, _pathsWithoutPrefix("client_ids", paths)...)
		case "field_mask":
			dst.FieldMask = src.FieldMask
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _ListClientsRequestFieldPaths = [...]string{
	"collaborator",
	"collaborator.organization_ids",
	"collaborator.organization_ids.organization_id",
	"collaborator.user_ids",
	"collaborator.user_ids.email",
	"collaborator.user_ids.user_id",
	"field_mask",
	"limit",
	"order",
	"page",
}

func (*ListClientsRequest) FieldMaskPaths() []string {
	ret := make([]string, len(_ListClientsRequestFieldPaths))
	copy(ret, _ListClientsRequestFieldPaths[:])
	return ret
}

func (dst *ListClientsRequest) SetFields(src *ListClientsRequest, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "collaborator":
			dst.Collaborator = src.Collaborator
		case "collaborator.organization_ids":
			if dst.Collaborator == nil {
				dst.Collaborator = &OrganizationOrUserIdentifiers{}
			}
			dst.Collaborator.SetFields(src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.organization_ids.organization_id":
			if dst.Collaborator == nil {
				dst.Collaborator = &OrganizationOrUserIdentifiers{}
			}
			dst.Collaborator.SetFields(src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.user_ids":
			if dst.Collaborator == nil {
				dst.Collaborator = &OrganizationOrUserIdentifiers{}
			}
			dst.Collaborator.SetFields(src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.user_ids.email":
			if dst.Collaborator == nil {
				dst.Collaborator = &OrganizationOrUserIdentifiers{}
			}
			dst.Collaborator.SetFields(src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.user_ids.user_id":
			if dst.Collaborator == nil {
				dst.Collaborator = &OrganizationOrUserIdentifiers{}
			}
			dst.Collaborator.SetFields(src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "field_mask":
			dst.FieldMask = src.FieldMask
		case "limit":
			dst.Limit = src.Limit
		case "order":
			dst.Order = src.Order
		case "page":
			dst.Page = src.Page
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _CreateClientRequestFieldPaths = [...]string{
	"client",
	"client.attributes",
	"client.contact_info",
	"client.created_at",
	"client.description",
	"client.endorsed",
	"client.grants",
	"client.ids",
	"client.ids.client_id",
	"client.name",
	"client.redirect_uris",
	"client.rights",
	"client.secret",
	"client.skip_authorization",
	"client.state",
	"client.updated_at",
	"collaborator",
	"collaborator.organization_ids",
	"collaborator.organization_ids.organization_id",
	"collaborator.user_ids",
	"collaborator.user_ids.email",
	"collaborator.user_ids.user_id",
}

func (*CreateClientRequest) FieldMaskPaths() []string {
	ret := make([]string, len(_CreateClientRequestFieldPaths))
	copy(ret, _CreateClientRequestFieldPaths[:])
	return ret
}

func (dst *CreateClientRequest) SetFields(src *CreateClientRequest, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "client":
			dst.Client = src.Client
		case "client.attributes":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.contact_info":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.created_at":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.description":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.endorsed":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.grants":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.ids":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.ids.client_id":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.name":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.redirect_uris":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.rights":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.secret":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.skip_authorization":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.state":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.updated_at":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "collaborator":
			dst.Collaborator = src.Collaborator
		case "collaborator.organization_ids":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.organization_ids.organization_id":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.user_ids":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.user_ids.email":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.user_ids.user_id":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _UpdateClientRequestFieldPaths = [...]string{
	"client",
	"client.attributes",
	"client.contact_info",
	"client.created_at",
	"client.description",
	"client.endorsed",
	"client.grants",
	"client.ids",
	"client.ids.client_id",
	"client.name",
	"client.redirect_uris",
	"client.rights",
	"client.secret",
	"client.skip_authorization",
	"client.state",
	"client.updated_at",
	"field_mask",
}

func (*UpdateClientRequest) FieldMaskPaths() []string {
	ret := make([]string, len(_UpdateClientRequestFieldPaths))
	copy(ret, _UpdateClientRequestFieldPaths[:])
	return ret
}

func (dst *UpdateClientRequest) SetFields(src *UpdateClientRequest, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "client":
			dst.Client = src.Client
		case "client.attributes":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.contact_info":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.created_at":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.description":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.endorsed":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.grants":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.ids":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.ids.client_id":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.name":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.redirect_uris":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.rights":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.secret":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.skip_authorization":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.state":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "client.updated_at":
			dst.Client.SetFields(&src.Client, _pathsWithoutPrefix("client", paths)...)
		case "field_mask":
			dst.FieldMask = src.FieldMask
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}

var _SetClientCollaboratorRequestFieldPaths = [...]string{
	"client_ids",
	"client_ids.client_id",
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"collaborator.rights",
}

func (*SetClientCollaboratorRequest) FieldMaskPaths() []string {
	ret := make([]string, len(_SetClientCollaboratorRequestFieldPaths))
	copy(ret, _SetClientCollaboratorRequestFieldPaths[:])
	return ret
}

func (dst *SetClientCollaboratorRequest) SetFields(src *SetClientCollaboratorRequest, paths ...string) {
	for _, path := range _cleanPaths(paths) {
		switch path {
		case "client_ids":
			dst.ClientIdentifiers = src.ClientIdentifiers
		case "client_ids.client_id":
			dst.ClientIdentifiers.SetFields(&src.ClientIdentifiers, _pathsWithoutPrefix("client_ids", paths)...)
		case "collaborator":
			dst.Collaborator = src.Collaborator
		case "collaborator.ids":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.ids.organization_ids":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.ids.organization_ids.organization_id":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.ids.user_ids":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.ids.user_ids.email":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.ids.user_ids.user_id":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		case "collaborator.rights":
			dst.Collaborator.SetFields(&src.Collaborator, _pathsWithoutPrefix("collaborator", paths)...)
		default:
			panic(fmt.Errorf("invalid field path: '%s'", path))
		}
	}
}