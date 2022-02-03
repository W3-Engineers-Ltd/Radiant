// Copyright 2021 radiant
//

package berror

import (
	"fmt"
)

// pre define code

// Unknown indicates got some error which is not defined
var Unknown = DefineCode(5000001, "error", "Unknown", fmt.Sprintf(`
Unknown error code. Usually you will see this code in three cases:
1. You forget to define Code or function DefineCode not being executed;
2. This is not Radiant's error but you call FromError();
3. Radiant got unexpected error and don't know how to handle it, and then return Unknown error

A common practice to DefineCode looks like:
%s

In this way, you may forget to import this package, and got Unknown error. 

Sometimes, you believe you got Radiant error, but actually you don't, and then you call FromError(err)

`, goCodeBlock(`
import your_package

func init() {
    DefineCode(5100100, "your_module", "detail")
    // ...
}
`)))

func goCodeBlock(code string) string {
	return codeBlock("go", code)
}

func codeBlock(lan string, code string) string {
	return fmt.Sprintf("```%s\n%s\n```", lan, code)
}
