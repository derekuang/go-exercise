// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvins.
func CToK(c Celsius) Kelvins { return Kelvins(c + 273.15) }

// KToC converts a Kelvins temperature to Celsius.
func KToC(k Kelvins) Celsius { return Celsius(k - 273.15) }

// FToK converts a Fahrenheit temperature to Kelvins.
func FToK(f Fahrenheit) Kelvins { return CToK(FToC(f)) }

// KToF converts a Kelvins temperature to Fahrenheit.
func KToF(k Kelvins) Fahrenheit { return CToF(KToC(k)) }

//!-
