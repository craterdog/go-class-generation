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

package generator_test

import (
	gen "github.com/craterdog/go-class-generation/v5/generator"
	mod "github.com/craterdog/go-class-model/v5"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var directory = "../testdata/generator/"

func TestGeneration(t *tes.T) {
	// Validate the class model.
	var filename = "./Package.go"
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
	filename = directory + "Package.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}

	// Generate the class files.
	var generator = gen.ClassesGenerator().Make()
	var classes = generator.GenerateModelClasses(model).GetIterator()
	for classes.HasNext() {
		var class = classes.GetNext()
		var className = class.GetKey()
		filename = directory + className + ".go"
		source = class.GetValue()
		bytes = []byte(source)
		err = osx.WriteFile(filename, bytes, 0644)
		if err != nil {
			panic(err)
		}
	}
}
