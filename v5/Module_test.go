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

package module_test

import (
	fmt "fmt"
	gen "github.com/craterdog/go-class-generation/v5"
	mod "github.com/craterdog/go-class-model/v5"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var testDirectories = []string{
	"./testdata/ast/",
	"./testdata/grammar/",
	"./testdata/example/",
}

func TestGeneration(t *tes.T) {
	fmt.Println("Generating concrete classes for the following class models:")
	for _, directory := range testDirectories {

		// Validate the class model.
		var filename = directory + "Package.go"
		fmt.Printf("   %v\n", filename)
		var bytes, err = osx.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		var source = string(bytes)
		var model = mod.ParseSource(source)
		mod.ValidateModel(model)
		var actual = mod.FormatModel(model)
		ass.Equal(t, source, actual)

		// Recreate the package directory.
		err = osx.RemoveAll(directory)
		if err != nil {
			panic(err)
		}
		err = osx.Mkdir(directory, 0755)
		if err != nil {
			panic(err)
		}
		err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}

		// Generate the concrete class files for the class model.
		var classes = gen.GenerateModelClasses(model).GetIterator()
		for classes.HasNext() {
			var association = classes.GetNext()
			var className = association.GetKey()
			var classSource = association.GetValue()
			bytes = []byte(classSource)
			var filename = directory + className + ".go"
			err = osx.WriteFile(filename, bytes, 0644)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("Done.")
}
