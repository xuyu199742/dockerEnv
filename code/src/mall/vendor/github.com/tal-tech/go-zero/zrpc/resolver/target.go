package resolver

import (
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/zrpc/resolver/internal"
)

// BuildDirectTarget returns a string that represents the given endpoints with direct schema.
func BuildDirectTarget(endpoints []string) string {
	return fmt.Sprintf("%s:///%s", internal.DirectScheme,
		strings.Join(endpoints, internal.EndpointSep))
}

// BuildDiscovTarget returns a string that represents the given endpoints with discov schema.
func BuildDiscovTarget(endpoints []string, key string) string {
	return fmt.Sprintf("%s://%s/%s", internal.DiscovScheme,
		strings.Join(endpoints, internal.EndpointSep), key)
}
