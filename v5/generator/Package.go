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
Package "generator" provides the ability to generate Go class files based on a
Go Package.go file that follows the format shown in the following code template:
  - https://github.com/craterdog/go-class-model/blob/main/models/Package.go

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package generator

import (
	mod "github.com/craterdog/go-class-model/v5"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// Class Definitions

/*
ClassesGeneratorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete classes-like class.
*/
type ClassesGeneratorClassLike interface {
	// Constructor Methods
	Make() ClassesGeneratorLike
}

// Instance Definitions

/*
ClassesGeneratorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete classes-like class.
*/
type ClassesGeneratorLike interface {
	// Primary Methods
	GetClass() ClassesGeneratorClassLike
	GenerateModelClasses(
		model mod.ModelLike,
	) abs.CatalogLike[string, string]
}
