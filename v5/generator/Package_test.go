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
	fmt "fmt"
	mod "github.com/craterdog/go-class-model/v5"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

var inputDirectory = "./"
var outputDirectory = "../../../go-test-framework/v5/generator/"

func TestRoundTrips(t *tes.T) {
	fmt.Println("Round Trip Tests:")
	var filename = inputDirectory + "Package.go"
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
	filename = outputDirectory + "Package.go"
	err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done.")
}
