/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
Package "module" defines type aliases for the commonly used types defined in the
packages contained in this module.  It also provides a universal constructor for
the commonly used classes that are exported by this module.  Each constructor
delegates the actual construction process to its corresponding concrete class
defined in a package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-class-model/wiki
*/
package module

import (
	gen "github.com/craterdog/go-class-generation/v5/generator"
	mod "github.com/craterdog/go-class-model/v5"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// TYPE ALIASES

type (
	ClassesLike = gen.ClassesLike
)

// UNIVERSAL CONSTRUCTORS

func Classes(args ...any) ClassesLike {
	if len(args) > 0 {
		panic("The \"classes\" constructor does not take any arguments.")
	}
	var classes = gen.Classes().Make()
	return classes
}

// GLOBAL FUNCTIONS

func GenerateModelClasses(model mod.ModelLike) abs.CatalogLike[string, string] {
	var generator = Classes()
	var catalog = generator.GenerateModelClasses(model)
	return catalog
}
