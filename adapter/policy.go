// Copyright 2016 radiant authors. All Rights Reserved.
//

package adapter

import (
	"github.com/W3-Engineers-Ltd/Radiant/adapter/context"
	"github.com/W3-Engineers-Ltd/Radiant/server/web"
	radicalcontext "github.com/W3-Engineers-Ltd/Radiant/server/web/context"
)

// PolicyFunc defines a policy function which is invoked before the controller handler is executed.
type PolicyFunc func(*context.Context)

// FindPolicy Find Router info for URL
func (p *ControllerRegister) FindPolicy(cont *context.Context) []PolicyFunc {
	pf := (*web.ControllerRegister)(p).FindPolicy((*radicalcontext.Context)(cont))
	npf := newToOldPolicyFunc(pf)
	return npf
}

func newToOldPolicyFunc(pf []web.PolicyFunc) []PolicyFunc {
	npf := make([]PolicyFunc, 0, len(pf))
	for _, f := range pf {
		npf = append(npf, func(c *context.Context) {
			f((*radicalcontext.Context)(c))
		})
	}
	return npf
}

func oldToNewPolicyFunc(pf []PolicyFunc) []web.PolicyFunc {
	npf := make([]web.PolicyFunc, 0, len(pf))
	for _, f := range pf {
		npf = append(npf, func(c *radicalcontext.Context) {
			f((*context.Context)(c))
		})
	}
	return npf
}

// Policy Register new policy in radiant
func Policy(pattern, method string, policy ...PolicyFunc) {
	pf := oldToNewPolicyFunc(policy)
	web.Policy(pattern, method, pf...)
}
