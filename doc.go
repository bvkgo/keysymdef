/*
Package keysymdef defines Go constants and unicode codepoint mappings for the
X11 key symbols from /usr/include/X11/keysymdef.h file.

Key symbol names are the same as their C language identifiers, including the
underscores, XK_ prefix and inconsistent capitalization (for example,
XK_Mode_switch, XK_script_switch, XK_Num_Lock, etc.).
*/
package keysymdef

//go:generate go run gen.go
