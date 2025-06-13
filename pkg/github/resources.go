package github

import (
	"github.com/github/github-mcp-server/pkg/translations"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterResources(s *server.MCPServer, getClient GetClientFn, t translations.TranslationHelperFunc) {
	s.AddResourceTemplate(GetRepositoryResourceContent(getClient, t))
	s.AddResourceTemplate(GetRepositoryResourceBranchContent(getClient, t))
	s.AddResourceTemplate(GetRepositoryResourceCommitContent(getClient, t))
	s.AddResourceTemplate(GetRepositoryResourceTagContent(getClient, t))
	s.AddResourceTemplate(GetRepositoryResourcePrContent(getClient, t))
}

// RegisterMultiUserResources registers resources for multi-user mode
// For now, this is the same as RegisterResources since resources don't need 
// the same auth token wrapper as tools (resources are read-only and use URL patterns)
func RegisterMultiUserResources(s *server.MCPServer, getClient GetClientFn, t translations.TranslationHelperFunc) {
	RegisterResources(s, getClient, t)
}
