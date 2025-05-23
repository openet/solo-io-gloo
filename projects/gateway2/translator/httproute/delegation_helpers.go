package httproute

import (
	"path"
	"slices"
	"strings"

	"github.com/solo-io/gloo/projects/gateway2/query"
	"github.com/solo-io/gloo/projects/gateway2/utils"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// filterDelegatedChildren filters the referenced children and their rules based
// on parent matchers, filters their hostnames, and applies parent matcher
// inheritance
//
// A child route may other match a parent rule matcher or may opt to use inherit the parent's matcher.
// In the case of rule inheritance, a child route does not need to match the parent matcher. By default.
// a child route must match the parent matcher.
//
// A child route matches a parent if all of the following conditions are met:
//
// 1. The child route matches the route selector specified by the parent's backendRef
//
// 2.The child route has a rule that matches the parent's route matcher:
//   - The child route's path must contain the the parent's path as a prefix
//   - The child route's headers must be a superset of the parent's headers
//   - The child route's query parameters must be a superset of the parent's query parameters
//
// If a child route's rule does not match the given parent match, it is not included in the route returned.
func filterDelegatedChildren(
	parentRef types.NamespacedName,
	parentMatch gwv1.HTTPRouteMatch,
	children []*query.RouteInfo,
) []*query.RouteInfo {
	// Select the child routes that match the parent
	var selected []*query.RouteInfo
	for _, c := range children {
		origChild, ok := c.Object.(*gwv1.HTTPRoute)
		if !ok {
			continue
		}

		// This check is redundant as it is already done in fetchRoutesByRef in projects/gateway2/query/httproute.go,
		// so this is just to be extra careful if there is a bug when the child routes are resolved earlier.
		if !utils.ChildRouteCanAttachToParentRef(origChild, parentRef) {
			continue
		}

		// make a copy; multiple parents can delegate to the same child so we can't modify a shared reference
		clone := c.Clone()

		child, ok := clone.Object.(*gwv1.HTTPRoute)
		if !ok {
			continue
		}

		inheritMatcher := utils.ShouldChildRouteInheritParentMatcher(child)

		// Check if the child route has a prefix that matches the parent.
		// Only rules matching the parent prefix are considered.
		//
		// We use validRules to store the rules in the child route that are valid
		// (matches in the rule match the parent route matcher). If a specific rule
		// in the child is not valid, then we discard it in the final child route
		// returned by this function.
		var validRules []gwv1.HTTPRouteRule
		for i, rule := range child.Spec.Rules {
			var validMatches []gwv1.HTTPRouteMatch

			// If the child route opts to inherit the parent's matcher and it does not specify its own matcher,
			// simply inherit the parent's matcher.
			if inheritMatcher && len(rule.Matches) == 0 {
				validMatches = append(validMatches, parentMatch)
			}

			for _, match := range rule.Matches {
				if inheritMatcher {
					// When inheriting the parent's matcher, all matches are valid.
					// In this case, the child inherits the parents matcher so we merge
					// the parent's matcher with the child's.
					mergeParentChildRouteMatch(&parentMatch, &match)
					validMatches = append(validMatches, match)
				} else if ok := utils.IsDelegatedRouteMatch(parentMatch, match); ok {
					// Non-inherited matcher delegation requires matching child matcher to parent matcher
					// to delegate from the parent route to the child.
					validMatches = append(validMatches, match)
				}
			}

			// Matchers in this rule match the parent route matcher, so consider the valid matchers on the child,
			child.Spec.Rules[i].Matches = validMatches
			// and discard rules on the child that do not match the parent route matcher.
			if len(validMatches) > 0 {
				validRule := child.Spec.Rules[i]
				validRule.Matches = validMatches
				validRules = append(validRules, validRule)
			}
		}
		if len(validRules) > 0 {
			child.Spec.Rules = validRules
			clone.Object = child
			selected = append(selected, clone)
		}
	}

	return selected
}

// mergeParentChildRouteMatch merges the parent route match into the child.
func mergeParentChildRouteMatch(
	parent *gwv1.HTTPRouteMatch,
	child *gwv1.HTTPRouteMatch,
) {
	if parent == nil || child == nil {
		return
	}

	if child.Path == nil {
		child.Path = &gwv1.HTTPPathMatch{
			Type:  ptr.To(gwv1.PathMatchPathPrefix),
			Value: ptr.To(""),
		}
	}
	child.Path.Value = ptr.To(path.Join(*parent.Path.Value, *child.Path.Value))

	// Inherit parent and child headers and query parameters while augmenting the merge
	// with additions specified on the child
	child.Headers = mergeHeaders(parent.Headers, child.Headers)
	child.QueryParams = mergeQueries(parent.QueryParams, child.QueryParams)

	// If parent specifies a method, inherit it
	if parent.Method != nil {
		child.Method = ptr.To(*parent.Method)
	}
}

func mergeHeaders(
	parent, child []gwv1.HTTPHeaderMatch,
) []gwv1.HTTPHeaderMatch {
	merged := make(map[gwv1.HTTPHeaderName]gwv1.HTTPHeaderMatch)
	for _, h := range parent {
		merged[h.Name] = h
	}
	for _, h := range child {
		key := h.Name
		// Only add the child if it does not conflict with the parent
		if _, ok := merged[key]; !ok {
			merged[key] = h
		}
	}
	var result []gwv1.HTTPHeaderMatch
	for _, h := range merged {
		result = append(result, h)
	}
	// Sort for deterministic ordering
	slices.SortFunc(result, func(a, b gwv1.HTTPHeaderMatch) int {
		return strings.Compare(string(a.Name), string(b.Name))
	})
	return result
}

func mergeQueries(
	parent, child []gwv1.HTTPQueryParamMatch,
) []gwv1.HTTPQueryParamMatch {
	merged := make(map[gwv1.HTTPHeaderName]gwv1.HTTPQueryParamMatch)
	for _, h := range parent {
		merged[h.Name] = h
	}
	for _, h := range child {
		key := h.Name
		// Only add the child if it does not conflict with the parent
		if _, ok := merged[key]; !ok {
			merged[key] = h
		}
	}
	var result []gwv1.HTTPQueryParamMatch
	for _, h := range merged {
		result = append(result, h)
	}
	// Sort for deterministic ordering
	slices.SortFunc(result, func(a, b gwv1.HTTPQueryParamMatch) int {
		return strings.Compare(string(a.Name), string(b.Name))
	})
	return result
}
