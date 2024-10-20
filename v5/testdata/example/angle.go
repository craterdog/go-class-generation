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

package example

// CLASS INTERFACE

// Access Function

func Angle() AngleClassLike {
	return angleReference()
}

// Constructor Methods

func (c *angleClass_) Make(
	intrinsic float64,
) AngleLike {
	var instance = angle_(intrinsic)
	return instance

}

func (c *angleClass_) MakeFromString(
	value string,
) AngleLike {
	var instance AngleLike
	// TBD - Add the constructor implementation.
	return instance

}

// Constant Methods

func (c *angleClass_) Pi() AngleLike {
	return c.pi_
}

func (c *angleClass_) Tau() AngleLike {
	return c.tau_
}

// Function Methods

func (c *angleClass_) Apply(
	function TrigonometricFunction,
	angle AngleLike,
) float64 {
	var result_ float64
	// TBD - Add the function implementation.
	return result_
}

func (c *angleClass_) Sine(
	angle AngleLike,
) float64 {
	var result_ float64
	// TBD - Add the function implementation.
	return result_
}

func (c *angleClass_) Cosine(
	angle AngleLike,
) float64 {
	var result_ float64
	// TBD - Add the function implementation.
	return result_
}

func (c *angleClass_) Tangent(
	angle AngleLike,
) float64 {
	var result_ float64
	// TBD - Add the function implementation.
	return result_
}

// INSTANCE INTERFACE

// Primary Methods

func (v angle_) GetClass() AngleClassLike {
	return angleReference()
}

func (v angle_) GetIntrinsic() float64 {
	return float64(v)
}

func (v angle_) IsZero() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

// Angular Methods

func (v angle_) AsNormalized() AngleLike {
	var result_ AngleLike
	// TBD - Add the method implementation.
	return result_
}

func (v angle_) InUnits(
	units Units,
) float64 {
	var result_ float64
	// TBD - Add the method implementation.
	return result_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type angle_ float64

// Class Structure

type angleClass_ struct {
	// Declare the class constants.
	pi_  AngleLike
	tau_ AngleLike
}

// Class Reference

func angleReference() *angleClass_ {
	return angleReference_
}

var angleReference_ = &angleClass_{
	// Initialize the class constants.
	// pi_: constantValue,
	// tau_: constantValue,
}
