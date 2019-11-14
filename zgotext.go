// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p := messageKeyToIndex[key]
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en": &dictionary{index: enIndex, data: enData},
		"sl": &dictionary{index: slIndex, data: slData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%.2f\u00a0GiB":                         21,
	"%.2f\u00a0KiB":                         19,
	"%.2f\u00a0MiB":                         20,
	"%.2f\u00a0TiB":                         22,
	"%d day(s)":                             13,
	"%d hour(s)":                            14,
	"%d minute(s)":                          15,
	"%d second(s)":                          16,
	"%d tunnels were unable to be removed.": 157,
	"%d year(s)":                            12,
	"%d\u00a0B":                             18,
	"%s\n\nPlease consult the log for more information.": 110,
	"%s (out of date)":                        111,
	"%s (unsigned build, no updates)":         162,
	"%s You cannot undo this action.":         153,
	"%s ago":                                  17,
	"%s received, %s sent":                    70,
	"%s: %q":                                  23,
	"&About WireGuard…":                       108,
	"&Activate":                               57,
	"&Block untunneled traffic (kill-switch)": 81,
	"&Configuration:":                         85,
	"&Copy":                                   101,
	"&Deactivate":                             56,
	"&Edit":                                   132,
	"&Import tunnel(s) from file…":            118,
	"&Manage tunnels…":                        117,
	"&Name:":                                  78,
	"&Public key:":                            79,
	"&Remove selected tunnel(s)":              140,
	"&Save":                                   83,
	"&Save to file…":                          103,
	"&Toggle":                                 137,
	"(no argument): elevate and install manager service": 1,
	"(unknown)":                             80,
	"A name is required.":                   87,
	"A tunnel was unable to be removed: %s": 155,
	"About WireGuard":                       50,
	"Activating":                            96,
	"Active":                                95,
	"Add &empty tunnel…":                    133,
	"Add Tunnel":                            134,
	"Addresses:":                            61,
	"Addresses: %s":                         123,
	"Addresses: None":                       116,
	"All peers must have public keys":       44,
	"Allowed IPs:":                          64,
	"An Update is Available!":               128,
	"An interface must have a private key":  42,
	"An update to WireGuard is available. It is highly advisable to update without delay.":            165,
	"An update to WireGuard is now available. You are advised to update as soon as possible.":         130,
	"Another tunnel already exists with the name ‘%s’":                                                143,
	"Another tunnel already exists with the name ‘%s’.":                                               91,
	"App version: %s\nGo backend version: %s\nGo version: %s\nOperating system: %s\nArchitecture: %s": 52,
	"Are you sure you would like to delete %d tunnels?":                                               150,
	"Are you sure you would like to delete tunnel ‘%s’?":                                              152,
	"Brackets must contain an IPv6 address":                                                           28,
	"Cancel":                                                                                          84,
	"Close":                                                                                           53,
	"Command Line Options":                                                                            3,
	"Configuration Files (*.zip, *.conf)|*.zip;*.conf|All Files (*.*)|*.*":                            158,
	"Configuration ZIP Files (*.zip)|*.zip":                                                           160,
	"Could not enumerate existing tunnels: %v":                                                        142,
	"Could not import selected configuration: %v":                                                     141,
	"Create new tunnel":                                                                               76,
	"DNS servers:":                                                                                    62,
	"Deactivating":                                                                                    98,
	"Delete %d tunnels":                                                                               149,
	"Delete tunnel ‘%s’":                                                                              151,
	"E&xit":                                                                                           119,
	"Edit &selected tunnel…":                                                                          139,
	"Edit tunnel":                                                                                     77,
	"Endpoint:":                                                                                       65,
	"Error":                                                                                           0,
	"Error Exiting WireGuard":                                                                         163,
	"Error in getting configuration":                                                                  45,
	"Error: %v. Please try again.":                                                                    169,
	"Export all tunnels to &zip…":                                                                     138,
	"Export all tunnels to zip":                                                                       136,
	"Export log to file":                                                                              107,
	"Export tunnels to zip":                                                                           161,
	"Failed to activate tunnel":                                                                       72,
	"Failed to deactivate tunnel":                                                                     73,
	"Failed to determine tunnel state":                                                                71,
	"File ‘%s’ already exists.\n\nDo you want to overwrite it?":                                       94,
	"Import tunnel(s) from file":                                                                      159,
	"Imported %d of %d tunnels":                                                                       147,
	"Imported %d tunnels":                                                                             146,
	"Imported tunnels":                                                                                145,
	"Inactive":                                                                                        97,
	"Interface: %s":                                                                                   74,
	"Invalid IP address":                                                                              24,
	"Invalid MTU":                                                                                     29,
	"Invalid config key is missing an equals separator":                                               38,
	"Invalid endpoint host":                                                                           27,
	"Invalid key for [Interface] section":                                                             40,
	"Invalid key for [Peer] section":                                                                  41,
	"Invalid key for interface section":                                                               46,
	"Invalid key for peer section":                                                                    48,
	"Invalid key: %v":                                                                                 32,
	"Invalid name":                                                                                    86,
	"Invalid network prefix length":                                                                   25,
	"Invalid persistent keepalive":                                                                    31,
	"Invalid port":                                                                                    30,
	"Key must have a value":                                                                           39,
	"Keys must decode to exactly 32 bytes":                                                            33,
	"Latest handshake:":                                                                               67,
	"Line must occur in a section":                                                                    37,
	"Listen port:":                                                                                    59,
	"Log":                                                                                             100,
	"Log message":                                                                                     105,
	"MTU:":                                                                                            60,
	"Missing port from endpoint":                                                                      26,
	"Now":                                                                                             10,
	"Number must be a number between 0 and 2^64-1: %v":                                                34,
	"Peer":                                75,
	"Persistent keepalive:":               66,
	"Preshared key:":                      63,
	"Protocol version must be 1":          47,
	"Public key:":                         58,
	"Remove selected tunnel(s)":           135,
	"Select &all":                         102,
	"Status:":                             55,
	"Status: %s":                          122,
	"Status: Complete!":                   170,
	"Status: Unknown":                     115,
	"Status: Waiting for updater service": 168,
	"Status: Waiting for user":            166,
	"System clock wound backward!":        11,
	"Text Files (*.txt)|*.txt|All Files (*.*)|*.*": 106,
	"The %s tunnel has been activated.":            125,
	"The %s tunnel has been deactivated.":          127,
	"Time":                                         104,
	"Transfer:":                                    68,
	"Tunnel Error":                                 109,
	"Tunnel already exists":                        90,
	"Tunnel name is not valid":                     36,
	"Tunnel name ‘%s’ is invalid.":                 88,
	"Tunnels":                                      131,
	"Two commas in a row":                          35,
	"Unable to create new configuration":           92,
	"Unable to create tunnel":                      148,
	"Unable to delete tunnel":                      154,
	"Unable to delete tunnels":                     156,
	"Unable to determine whether the process is running under WOW64: %v":                          4,
	"Unable to exit service due to: %v. You may want to stop WireGuard from the service manager.": 164,
	"Unable to import configuration: %v":                                                          144,
	"Unable to list existing tunnels":                                                             89,
	"Unable to open current process token: %v":                                                    6,
	"Unable to wait for WireGuard window to appear: %v":                                           113,
	"Unknown state":    99,
	"Update Now":       167,
	"Usage: %s [\n%s]": 2,
	"When a configuration has exactly one peer, and that peer has an allowed IPs containing at least one of 0.0.0.0/0 or ::/0, then the tunnel service engages a firewall ruleset to block all traffic that is neither to nor from the tunnel interface, with special exceptions for DHCP and NDP.": 82,
	"WireGuard Activated":        124,
	"WireGuard Deactivated":      126,
	"WireGuard Detection Error":  112,
	"WireGuard Tunnel Error":     120,
	"WireGuard Update Available": 129,
	"WireGuard is running, but the UI is only accessible from desktops of the Builtin %s group.": 8,
	"WireGuard logo image": 51,
	"WireGuard may only be used by users who are a member of the Builtin %s group.": 7,
	"WireGuard system tray icon did not appear after 30 seconds.":                   9,
	"WireGuard: %s":          121,
	"WireGuard: Deactivated": 114,
	"Writing file failed":    93,
	"You must use the 64-bit version of WireGuard on this computer.": 5,
	"[EnumerationSeparator]": 49,
	"[none specified]":       43,
	"enabled":                69,
	"http2: Framer %p: failed to decode just-written frame": 171,
	"http2: Framer %p: read %v":                             173,
	"http2: Framer %p: wrote %v":                            172,
	"http2: decoded hpack field %+v":                        174,
	"♥ &Donate!":                                            54,
}

var enIndex = []uint32{ // 176 elements
	// Entry 0 - 1F
	0x00000000, 0x00000006, 0x00000039, 0x0000004f,
	0x00000064, 0x000000aa, 0x000000e9, 0x00000115,
	0x00000166, 0x000001c4, 0x00000200, 0x00000204,
	0x00000221, 0x00000241, 0x0000025f, 0x0000027f,
	0x000002a3, 0x000002c7, 0x000002d1, 0x000002da,
	0x000002e7, 0x000002f4, 0x00000301, 0x0000030e,
	0x0000031b, 0x0000032e, 0x0000034c, 0x00000367,
	0x0000037d, 0x000003a3, 0x000003af, 0x000003bc,
	// Entry 20 - 3F
	0x000003d9, 0x000003ec, 0x00000411, 0x00000445,
	0x00000459, 0x00000472, 0x0000048f, 0x000004c1,
	0x000004d7, 0x000004fb, 0x0000051a, 0x0000053f,
	0x00000550, 0x00000570, 0x0000058f, 0x000005b1,
	0x000005cc, 0x000005e9, 0x000005ec, 0x000005fc,
	0x00000611, 0x0000067c, 0x00000682, 0x0000068f,
	0x00000697, 0x000006a3, 0x000006ad, 0x000006b9,
	0x000006c6, 0x000006cb, 0x000006d6, 0x000006e3,
	// Entry 40 - 5F
	0x000006f2, 0x000006ff, 0x00000709, 0x0000071f,
	0x00000731, 0x0000073b, 0x00000743, 0x0000075e,
	0x0000077f, 0x00000799, 0x000007b5, 0x000007c6,
	0x000007cb, 0x000007dd, 0x000007e9, 0x000007f0,
	0x000007fd, 0x00000807, 0x0000082f, 0x0000094d,
	0x00000953, 0x0000095a, 0x0000096a, 0x00000977,
	0x0000098b, 0x000009af, 0x000009cf, 0x000009e5,
	0x00000a1e, 0x00000a41, 0x00000a55, 0x00000a94,
	// Entry 60 - 7F
	0x00000a9b, 0x00000aa6, 0x00000aaf, 0x00000abc,
	0x00000aca, 0x00000ace, 0x00000ad4, 0x00000ae0,
	0x00000af1, 0x00000af6, 0x00000b02, 0x00000b2f,
	0x00000b42, 0x00000b56, 0x00000b63, 0x00000b97,
	0x00000bab, 0x00000bc5, 0x00000bfa, 0x00000c11,
	0x00000c21, 0x00000c31, 0x00000c44, 0x00000c63,
	0x00000c69, 0x00000c80, 0x00000c91, 0x00000c9f,
	0x00000cb0, 0x00000cc4, 0x00000ce9, 0x00000cff,
	// Entry 80 - 9F
	0x00000d26, 0x00000d3e, 0x00000d59, 0x00000db1,
	0x00000db9, 0x00000dbf, 0x00000dd4, 0x00000ddf,
	0x00000df9, 0x00000e13, 0x00000e1b, 0x00000e39,
	0x00000e52, 0x00000e6d, 0x00000e9c, 0x00000ec8,
	0x00000f00, 0x00000f26, 0x00000f37, 0x00000f6d,
	0x00000fb4, 0x00000fcc, 0x00000ffe, 0x00001070,
	0x0000108a, 0x000010c4, 0x000010e7, 0x000010ff,
	0x00001128, 0x00001141, 0x0000119a, 0x000011df,
	// Entry A0 - BF
	0x000011fa, 0x00001220, 0x00001236, 0x00001259,
	0x00001271, 0x000012d0, 0x00001325, 0x0000133e,
	0x00001349, 0x0000136d, 0x0000138d, 0x0000139f,
	0x000013d8, 0x000013f9, 0x00001419, 0x0000143b,
} // Size: 728 bytes

const enData string = "" + // Size: 5179 bytes
	"\x02Error\x02(no argument): elevate and install manager service\x02Usage" +
	": %[1]s [\x0a%[2]s]\x02Command Line Options\x02Unable to determine wheth" +
	"er the process is running under WOW64: %[1]v\x02You must use the 64-bit " +
	"version of WireGuard on this computer.\x02Unable to open current process" +
	" token: %[1]v\x02WireGuard may only be used by users who are a member of" +
	" the Builtin %[1]s group.\x02WireGuard is running, but the UI is only ac" +
	"cessible from desktops of the Builtin %[1]s group.\x02WireGuard system t" +
	"ray icon did not appear after 30 seconds.\x02Now\x02System clock wound b" +
	"ackward!\x14\x01\x81\x01\x00\x02\x0b\x02%[1]d year\x00\x0c\x02%[1]d year" +
	"s\x14\x01\x81\x01\x00\x02\x0a\x02%[1]d day\x00\x0b\x02%[1]d days\x14\x01" +
	"\x81\x01\x00\x02\x0b\x02%[1]d hour\x00\x0c\x02%[1]d hours\x14\x01\x81" +
	"\x01\x00\x02\x0d\x02%[1]d minute\x00\x0e\x02%[1]d minutes\x14\x01\x81" +
	"\x01\x00\x02\x0d\x02%[1]d second\x00\x0e\x02%[1]d seconds\x02%[1]s ago" +
	"\x02%[1]d\u00a0B\x02%.2[1]f\u00a0KiB\x02%.2[1]f\u00a0MiB\x02%.2[1]f" +
	"\u00a0GiB\x02%.2[1]f\u00a0TiB\x02%[1]s: %[2]q\x02Invalid IP address\x02I" +
	"nvalid network prefix length\x02Missing port from endpoint\x02Invalid en" +
	"dpoint host\x02Brackets must contain an IPv6 address\x02Invalid MTU\x02I" +
	"nvalid port\x02Invalid persistent keepalive\x02Invalid key: %[1]v\x02Key" +
	"s must decode to exactly 32 bytes\x02Number must be a number between 0 a" +
	"nd 2^64-1: %[1]v\x02Two commas in a row\x02Tunnel name is not valid\x02L" +
	"ine must occur in a section\x02Invalid config key is missing an equals s" +
	"eparator\x02Key must have a value\x02Invalid key for [Interface] section" +
	"\x02Invalid key for [Peer] section\x02An interface must have a private k" +
	"ey\x02[none specified]\x02All peers must have public keys\x02Error in ge" +
	"tting configuration\x02Invalid key for interface section\x02Protocol ver" +
	"sion must be 1\x02Invalid key for peer section\x02, \x02About WireGuard" +
	"\x02WireGuard logo image\x02App version: %[1]s\x0aGo backend version: %[" +
	"2]s\x0aGo version: %[3]s\x0aOperating system: %[4]s\x0aArchitecture: %[5" +
	"]s\x02Close\x02♥ &Donate!\x02Status:\x02&Deactivate\x02&Activate\x02Publ" +
	"ic key:\x02Listen port:\x02MTU:\x02Addresses:\x02DNS servers:\x02Preshar" +
	"ed key:\x02Allowed IPs:\x02Endpoint:\x02Persistent keepalive:\x02Latest " +
	"handshake:\x02Transfer:\x02enabled\x02%[1]s received, %[2]s sent\x02Fail" +
	"ed to determine tunnel state\x02Failed to activate tunnel\x02Failed to d" +
	"eactivate tunnel\x02Interface: %[1]s\x02Peer\x02Create new tunnel\x02Edi" +
	"t tunnel\x02&Name:\x02&Public key:\x02(unknown)\x02&Block untunneled tra" +
	"ffic (kill-switch)\x02When a configuration has exactly one peer, and tha" +
	"t peer has an allowed IPs containing at least one of 0.0.0.0/0 or ::/0, " +
	"then the tunnel service engages a firewall ruleset to block all traffic " +
	"that is neither to nor from the tunnel interface, with special exception" +
	"s for DHCP and NDP.\x02&Save\x02Cancel\x02&Configuration:\x02Invalid nam" +
	"e\x02A name is required.\x02Tunnel name ‘%[1]s’ is invalid.\x02Unable to" +
	" list existing tunnels\x02Tunnel already exists\x02Another tunnel alread" +
	"y exists with the name ‘%[1]s’.\x02Unable to create new configuration" +
	"\x02Writing file failed\x02File ‘%[1]s’ already exists.\x0a\x0aDo you wa" +
	"nt to overwrite it?\x02Active\x02Activating\x02Inactive\x02Deactivating" +
	"\x02Unknown state\x02Log\x02&Copy\x02Select &all\x02&Save to file…\x02Ti" +
	"me\x02Log message\x02Text Files (*.txt)|*.txt|All Files (*.*)|*.*\x02Exp" +
	"ort log to file\x02&About WireGuard…\x02Tunnel Error\x02%[1]s\x0a\x0aPle" +
	"ase consult the log for more information.\x02%[1]s (out of date)\x02Wire" +
	"Guard Detection Error\x02Unable to wait for WireGuard window to appear: " +
	"%[1]v\x02WireGuard: Deactivated\x02Status: Unknown\x02Addresses: None" +
	"\x02&Manage tunnels…\x02&Import tunnel(s) from file…\x02E&xit\x02WireGua" +
	"rd Tunnel Error\x02WireGuard: %[1]s\x02Status: %[1]s\x02Addresses: %[1]s" +
	"\x02WireGuard Activated\x02The %[1]s tunnel has been activated.\x02WireG" +
	"uard Deactivated\x02The %[1]s tunnel has been deactivated.\x02An Update " +
	"is Available!\x02WireGuard Update Available\x02An update to WireGuard is" +
	" now available. You are advised to update as soon as possible.\x02Tunnel" +
	"s\x02&Edit\x02Add &empty tunnel…\x02Add Tunnel\x02Remove selected tunnel" +
	"(s)\x02Export all tunnels to zip\x02&Toggle\x02Export all tunnels to &zi" +
	"p…\x02Edit &selected tunnel…\x02&Remove selected tunnel(s)\x02Could not " +
	"import selected configuration: %[1]v\x02Could not enumerate existing tun" +
	"nels: %[1]v\x02Another tunnel already exists with the name ‘%[1]s’\x02Un" +
	"able to import configuration: %[1]v\x02Imported tunnels\x14\x01\x81\x01" +
	"\x00\x02\x16\x02Imported %[1]d tunnel\x00\x17\x02Imported %[1]d tunnels" +
	"\x14\x02\x80\x01\x02\x1f\x02Imported %[1]d of %[2]d tunnel\x00 \x02Impor" +
	"ted %[1]d of %[2]d tunnels\x02Unable to create tunnel\x14\x01\x81\x01" +
	"\x00\x02\x14\x02Delete %[1]d tunnel\x00\x15\x02Delete %[1]d tunnels\x14" +
	"\x01\x81\x01\x00\x024\x02Are you sure you would like to delete %[1]d tun" +
	"nel?\x005\x02Are you sure you would like to delete %[1]d tunnels?\x02Del" +
	"ete tunnel ‘%[1]s’\x02Are you sure you would like to delete tunnel ‘%[1]" +
	"s’?\x02%[1]s You cannot undo this action.\x02Unable to delete tunnel\x02" +
	"A tunnel was unable to be removed: %[1]s\x02Unable to delete tunnels\x14" +
	"\x01\x81\x01\x00\x02'\x02%[1]d tunnel was unable to be removed.\x00)\x02" +
	"%[1]d tunnels were unable to be removed.\x02Configuration Files (*.zip, " +
	"*.conf)|*.zip;*.conf|All Files (*.*)|*.*\x02Import tunnel(s) from file" +
	"\x02Configuration ZIP Files (*.zip)|*.zip\x02Export tunnels to zip\x02%[" +
	"1]s (unsigned build, no updates)\x02Error Exiting WireGuard\x02Unable to" +
	" exit service due to: %[1]v. You may want to stop WireGuard from the ser" +
	"vice manager.\x02An update to WireGuard is available. It is highly advis" +
	"able to update without delay.\x02Status: Waiting for user\x02Update Now" +
	"\x02Status: Waiting for updater service\x02Error: %[1]v. Please try agai" +
	"n.\x02Status: Complete!\x02http2: Framer %[1]p: failed to decode just-wr" +
	"itten frame\x02http2: Framer %[1]p: wrote %[2]v\x02http2: Framer %[1]p: " +
	"read %[2]v\x02http2: decoded hpack field %+[1]v"

var slIndex = []uint32{ // 176 elements
	// Entry 0 - 1F
	0x00000000, 0x00000007, 0x00000058, 0x00000070,
	0x00000089, 0x000000c1, 0x00000107, 0x0000013e,
	0x00000190, 0x000001f6, 0x0000023a, 0x0000023f,
	0x0000025e, 0x00000296, 0x000002cd, 0x00000301,
	0x00000341, 0x00000385, 0x00000391, 0x0000039a,
	0x000003a7, 0x000003b4, 0x000003c1, 0x000003ce,
	0x000003db, 0x000003ee, 0x00000412, 0x00000434,
	0x0000045d, 0x00000483, 0x00000490, 0x0000049f,
	// Entry 20 - 3F
	0x000004c3, 0x000004da, 0x0000050b, 0x0000053f,
	0x00000554, 0x0000056b, 0x00000586, 0x000005be,
	0x000005d9, 0x000005fe, 0x0000061e, 0x00000640,
	0x0000064e, 0x00000675, 0x00000695, 0x000006b7,
	0x000006d5, 0x000006f7, 0x000006fa, 0x00000707,
	0x00000725, 0x0000079a, 0x000007a0, 0x000007ae,
	0x000007b6, 0x000007c3, 0x000007ce, 0x000007dc,
	0x000007ef, 0x000007f4, 0x000007fd, 0x0000080d,
	// Entry 40 - 5F
	0x00000823, 0x00000834, 0x00000844, 0x00000860,
	0x00000872, 0x0000087a, 0x00000885, 0x000008a2,
	0x000008c6, 0x000008e4, 0x00000904, 0x00000913,
	0x0000091b, 0x0000092d, 0x00000939, 0x0000093f,
	0x0000094e, 0x00000958, 0x00000982, 0x00000aa3,
	0x00000aab, 0x00000ab5, 0x00000ac5, 0x00000ad2,
	0x00000ae2, 0x00000b04, 0x00000b34, 0x00000b46,
	0x00000b71, 0x00000b98, 0x00000bb6, 0x00000bf1,
	// Entry 60 - 7F
	0x00000bf9, 0x00000c05, 0x00000c0f, 0x00000c1d,
	0x00000c2c, 0x00000c34, 0x00000c3d, 0x00000c49,
	0x00000c61, 0x00000c66, 0x00000c7c, 0x00000cb4,
	0x00000cce, 0x00000ce1, 0x00000cef, 0x00000d1e,
	0x00000d34, 0x00000d51, 0x00000d8c, 0x00000da3,
	0x00000db2, 0x00000dc0, 0x00000dd7, 0x00000df6,
	0x00000dfd, 0x00000e15, 0x00000e26, 0x00000e34,
	0x00000e43, 0x00000e57, 0x00000e75, 0x00000e8b,
	// Entry 80 - 9F
	0x00000eab, 0x00000ec4, 0x00000ee4, 0x00000f29,
	0x00000f30, 0x00000f37, 0x00000f50, 0x00000f5c,
	0x00000f74, 0x00000f8c, 0x00000f96, 0x00000fb4,
	0x00000fcd, 0x00000fe6, 0x00001014, 0x00001047,
	0x0000106c, 0x00001092, 0x000010a2, 0x00001106,
	0x00001192, 0x000011ae, 0x00001213, 0x00001300,
	0x0000131b, 0x00001356, 0x00001381, 0x0000139b,
	0x000013c3, 0x000013de, 0x00001492, 0x000014df,
	// Entry A0 - BF
	0x000014f8, 0x00001523, 0x00001537, 0x00001566,
	0x00001586, 0x000015f3, 0x00001646, 0x00001662,
	0x00001670, 0x00001697, 0x000016b9, 0x000016cb,
	0x00001713, 0x00001737, 0x0000175b, 0x00001780,
} // Size: 728 bytes

const slData string = "" + // Size: 6016 bytes
	"\x02Napaka\x02(brez argumenta): povzdigni na skrbniške pravice in namest" +
	"i skrbniško storitev\x02Uporaba: %[1]s [\x0a%[2]s]\x02Možnosti ukazne vr" +
	"stice\x02Napaka pri določanju ali proces teče kot WOW64: %[1]v\x02Na tem" +
	"u računalniku morate uporabiti 64-bitno različico WireGuarda.\x02Napaka " +
	"pri odpiranju žetona trenutnega procesa: %[1]v\x02WireGuard lahko uporab" +
	"ljajo samo uporabniki, ki so člani vgrajene skupine %[1]s.\x02WireGuard " +
	"je zagnan, vendar je up. vmesnik dostopen samo z namizij uporabnikov čla" +
	"nov skupine %[1]s.\x02Ikona WireGuarda se po 30 sekundah ni pojavila v s" +
	"istemski vrstici.\x02Zdaj\x02Sistemska ura prevrtena nazaj!\x14\x01\x81" +
	"\x01\x00\x04\x0b\x02%[1]d leta\x02\x0b\x02%[1]d leto\x03\x0b\x02%[1]d le" +
	"ti\x00\x0a\x02%[1]d let\x14\x01\x81\x01\x00\x04\x0a\x02%[1]d dni\x02\x0a" +
	"\x02%[1]d dan\x03\x0c\x02%[1]d dneva\x00\x0a\x02%[1]d dni\x14\x01\x81" +
	"\x01\x00\x04\x0a\x02%[1]d ure\x02\x0a\x02%[1]d uro\x03\x0a\x02%[1]d uri" +
	"\x00\x09\x02%[1]d ur\x14\x01\x81\x01\x00\x04\x0d\x02%[1]d minute\x02\x0d" +
	"\x02%[1]d minuto\x03\x0d\x02%[1]d minuti\x00\x0c\x02%[1]d minut\x14\x01" +
	"\x81\x01\x00\x04\x0e\x02%[1]d sekunde\x02\x0e\x02%[1]d sekundo\x03\x0e" +
	"\x02%[1]d sekundi\x00\x0d\x02%[1]d sekund\x02%[1]s nazaj\x02%[1]d\u00a0B" +
	"\x02%.2[1]f\u00a0KiB\x02%.2[1]f\u00a0MiB\x02%.2[1]f\u00a0GiB\x02%.2[1]f" +
	"\u00a0TiB\x02%[1]s: %[2]q\x02Napačen naslov IP\x02Napačna dolžina predpo" +
	"ne omrežja\x02Pri končni točki manjkajo vrata\x02Pri končni točki je gos" +
	"titelj napačen\x02Oklepaji morajo vsebovati naslov IPv6\x02Napačen MTU" +
	"\x02Napačna vrata\x02Napačno trajno ohranjanje povezave\x02Napačen ključ" +
	": %[1]v\x02Dekodirani ključi morajo biti natanko 32 bajtov\x02Številka m" +
	"ora biti število med 0 in 2^64-1: %[1]v\x02Dve zaporedni vejici\x02Ime t" +
	"unela ni veljavno\x02Vrstica mora biti v odseku\x02Napačnemu ključu konf" +
	"iguracije manjka ločilo-enačaj\x02Ključ mora imeti vrednost\x02Napačen k" +
	"ljuč za odsek [Interface]\x02Napačen ključ za odsek [Peer]\x02Vmesnik mo" +
	"ra imeti zasebni ključ\x02[ni navedeno]\x02Vsi vrstniki morajo imeti jav" +
	"ni ključ\x02Napaka pri branju konfiguracije\x02Napačen ključ za odsek vm" +
	"esnika\x02Verzija protokola mora biti 1\x02Napačen ključ za odsek vrstni" +
	"ka\x02, \x02O WireGuardu\x02Slika WireGuardovega logotipa\x02Verzija apl" +
	"ikacije: %[1]s\x0aVerzija wireguard-go: %[2]s\x0aVerzija Go: %[3]s\x0aOp" +
	"eracijski sistem: %[4]s\x0aArhitektura: %[5]s\x02Zapri\x02♥ &Doniraj!" +
	"\x02Status:\x02&Deaktiviraj\x02&Aktiviraj\x02Javni ključ:\x02Vrata poslu" +
	"šanja:\x02MTU:\x02Naslovi:\x02Strežniki DNS:\x02Ključ v skupni rabi:" +
	"\x02Dovoljeni IP-ji:\x02Končna točka:\x02Trajno ohranjanje povezave:\x02" +
	"Zadnje rokovanje:\x02Prenos:\x02omogočeno\x02%[1]s prejeto, %[2]s poslan" +
	"o\x02Napaka pri določanju stanja tunela\x02Napaka pri aktiviranju tunela" +
	"\x02Napaka pri deaktiviranju tunela\x02Vmesnik: %[1]s\x02Vrstnik\x02Ustv" +
	"ari nov tunel\x02Uredi tunel\x02&Ime:\x02&Javni ključ:\x02(neznano)\x02&" +
	"Blokiraj promet izven tunela (varovalka)\x02Kadar ima konfiguracija nata" +
	"nko enega vrstnika in njegov spisek dovoljenih IP-jev vsebuje vsaj enega" +
	" izmed 0.0.0.0/0 ali ::/0, bo storitev tunela vzpostavila pravila požarn" +
	"ega zidu, ki bodo blokirala ves promet, ki ni niti za niti iz vmesnika t" +
	"unela s posebnimi izjemami za DHCP and NDP.\x02&Shrani\x02Prekliči\x02&K" +
	"onfiguracija:\x02Napačno ime\x02Ime je obvezno.\x02Ime tunela »%[1]s« ni" +
	" veljavno.\x02Napaka pri pripravi seznama obstoječih tunelov\x02Tunel že" +
	" obstaja\x02Drug tunel z imenom »%[1]s« že obstaja.\x02Napaka pri izdela" +
	"vi nove konfiguracije\x02Napaka pri pisanju v datoteko\x02Datoteka »%[1]" +
	"s« že obstaja.\x0a\x0aAli jo želite prepisati?\x02Aktiven\x02Se aktivira" +
	"\x02Neaktiven\x02Se deaktivira\x02Neznano stanje\x02Dnevnik\x02&Kopiraj" +
	"\x02&Izberi vse\x02&Shrani v datoteko\u00a0…\x02Čas\x02Sporočilo v dnevn" +
	"iku\x02Tekstovne datoteke (*.txt)|*.txt|Vse datoteke (*.*)|*.*\x02Izvozi" +
	" dnevnik v datoteko\x02O WireGu&ardu\u00a0…\x02Napaka tunela\x02%[1]s" +
	"\x0a\x0aDodatne informacije najdete v dnevniku.\x02%[1]s (neposodobljen)" +
	"\x02Napaka zaznavanja WireGuarda\x02Čakanje, da se pojavi WireGuardovo o" +
	"kno, ni možno: %[1]v\x02WireGuard: Deaktiviran\x02Status: Neznan\x02Nasl" +
	"ovi: Brez\x02&Upravljaj tunele\u00a0…\x02Uvoz&i tunele iz datoteke\u00a0" +
	"…\x02I&zhod\x02Napaka tunela WireGuard\x02WireGuard: %[1]s\x02Status: " +
	"%[1]s\x02Naslovi: %[1]s\x02WireGuard aktiviran\x02Tunel %[1]s je bil akt" +
	"iviran.\x02WireGuard deaktiviran\x02Tunel %[1]s je bil deaktiviran.\x02N" +
	"a voljo je posodobitev!\x02Posodobitev WireGuarda na voljo\x02Posodobite" +
	"v WireGuarda je na voljo. Svetujemo posodobitev čim prej.\x02Tuneli\x02U" +
	"r&edi\x02Dodaj praz&en tunel\u00a0…\x02Dodaj tunel\x02Odstrani izbrane t" +
	"unele\x02Izvozi vse tunele v zip\x02&Preklopi\x02Izvozi vse tunele v &zi" +
	"p\u00a0…\x02Uredi i&zbran tunel\u00a0…\x02Odst&rani izbrane tunele\x02Na" +
	"paka pri uvozu izbrane konfiguracije: %[1]v\x02Napaka pri preštevanju ob" +
	"stoječih tunelov: %[1]v\x02Tunel z imenom »%[1]s« že obstaja\x02Napaka p" +
	"ri uvozu konfiguracije: %[1]v\x02Uvoženi tuneli\x14\x01\x81\x01\x00\x04" +
	"\x16\x02Uvoženi %[1]d tuneli\x02\x14\x02Uvožen %[1]d tunel\x03\x16\x02Uv" +
	"ožena %[1]d tunela\x00\x17\x02Uvoženo %[1]d tunelov\x14\x01\x81\x01\x00" +
	"\x04 \x02Uvoženi %[1]d od %[2]d tunelov\x02\x1f\x02Uvožen %[1]d od %[2]d" +
	" tunelov\x03 \x02Uvožena %[1]d od %[2]d tunelov\x00 \x02Uvoženo %[1]d od" +
	" %[2]d tunelov\x02Napaka pri stvaritvi tunela\x14\x01\x81\x01\x00\x04" +
	"\x16\x02Izbriši %[1]d tunele\x02\x15\x02Izbriši %[1]d tunel\x03\x16\x02I" +
	"zbriši %[1]d tunela\x00\x17\x02Izbriši %[1]d tunelov\x14\x01\x81\x01\x00" +
	"\x048\x02Ali ste prepričani, da želite izbrisati %[1]d tunele?\x027\x02A" +
	"li ste prepričani, da želite izbrisati %[1]d tunel?\x038\x02Ali ste prep" +
	"ričani, da želite izbrisati %[1]d tunela?\x009\x02Ali ste prepričani, da" +
	" želite izbrisati %[1]d tunelov?\x02Izbriši tunel ‘%[1]s’\x02Ali ste pre" +
	"pričani, da želite izbrisati tunel »%[1]s«?\x02%[1]s Tega dejanja ne mor" +
	"ete razveljaviti.\x02Napaka pri izbrisu tunela\x02Napaka pri odstranjeva" +
	"nju tunela: %[1]s\x02Napaka pri izbrisu tunelov\x14\x01\x81\x01\x00\x04*" +
	"\x02%[1]d tunelov ni bilo mogoče odstraniti.\x02)\x02%[1]d tunela ni bil" +
	"o mogoče odstraniti.\x03*\x02%[1]d tunelov ni bilo mogoče odstraniti." +
	"\x00*\x02%[1]d tunelov ni bilo mogoče odstraniti.\x02Konfiguracijske dat" +
	"oteke (*.zip, *.conf)|*.zip;*.conf|Vse datoteke (*.*)|*.*\x02Uvozi tunel" +
	"e iz datoteke\x02Konfiguracijske datoteke ZIP (*.zip)|*.zip\x02Izvozi tu" +
	"nele v zip\x02%[1]s (nepodpisane izdelave, brez posodobitev)\x02Napaka p" +
	"ri izhodu iz WireGuarda\x02Storitve ni bilo mogoče zaustaviti, ker: %[1]" +
	"v. Poskusite zaustaviti WireGuard z uporabo programa Storitve.\x02Posodo" +
	"bitev WireGuarda je na voljo. Zelo priporočamo posodobitev brez odlašanj" +
	"a.\x02Status: Čaka na uporabnika\x02Posodobi zdaj\x02Status: Čaka na ser" +
	"vis za posodobitev\x02Napaka: %[1]v. Poskusite ponovno.\x02Status: Konča" +
	"no!\x02http2: Framer %[1]p: napaka pri dekodiranju ravnokar zapisanega o" +
	"kvirja\x02http2: Framer %[1]p: zapisano %[2]v\x02http2: Framer %[1]p: pr" +
	"ebrano %[2]v\x02http2: dekodirano polje hpack %+[1]v"

	// Total table size 12651 bytes (12KiB); checksum: 7F514737
