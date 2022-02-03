// Copyright 2016 radiant authors. All Rights Reserved.
//

package web

import (
	"strings"

	"github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

// PolicyFunc defines a policy function which is invoked before the controller handler is executed.
type PolicyFunc func(*context.Context)

// FindPolicy Find Router info for URL
func (p *ControllerRegister) FindPolicy(cont *context.Context) []PolicyFunc {
	urlPath := cont.Input.URL()
	if !BConfig.RouterCaseSensitive {
		urlPath = strings.ToLower(urlPath)
	}
	httpMethod := cont.Input.Method()
	isWildcard := false
	// Find policy for current method
	t, ok := p.policies[httpMethod]
	// If not found - find policy for whole controller
	if !ok {
		t, ok = p.policies["*"]
		isWildcard = true
	}
	if ok {
		runObjects := t.Match(urlPath, cont)
		if r, ok := runObjects.([]PolicyFunc); ok {
			return r
		} else if !isWildcard {
			// If no policies found and we checked not for "*" method - try to find it
			t, ok = p.policies["*"]
			if ok {
				runObjects = t.Match(urlPath, cont)
				if r, ok = runObjects.([]PolicyFunc); ok {
					return r
				}
			}
		}
	}
	return nil
}

func (p *ControllerRegister) addToPolicy(method, pattern string, r ...PolicyFunc) {
	method = strings.ToUpper(method)
	p.enablePolicy = true
	if !BConfig.RouterCaseSensitive {
		pattern = strings.ToLower(pattern)
	}
	if t, ok := p.policies[method]; ok {
		t.AddRouter(pattern, r)
	} else {
		t := NewTree()
		t.AddRouter(pattern, r)
		p.policies[method] = t
	}
}

// Policy Register new policy in radiant
func Policy(pattern, method string, policy ...PolicyFunc) {
	RadicalApp.Handlers.addToPolicy(method, pattern, policy...)
}

// Find policies and execute if were found
func (p *ControllerRegister) execPolicy(cont *context.Context, urlPath string) (started bool) {
	if !p.enablePolicy {
		return false
	}
	// Find Policy for method
	policyList := p.FindPolicy(cont)
	if len(policyList) > 0 {
		// Run policies
		for _, runPolicy := range policyList {
			runPolicy(cont)
			if cont.ResponseWriter.Started {
				return true
			}
		}
		return false
	}
	return false
}
