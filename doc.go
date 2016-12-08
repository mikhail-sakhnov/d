/*
package d provides quick and dirty debugging output for tired programmers.

d.D() is a fast way to pretty-print variables. It's easier than typing
fmt.Printf("%#v", whatever). The output will be colorized and nicely formatted.
The output goes to stdout or $OUT env var if defined.

d exports a single D() function. This is how you use it:
    import "github.com/soider/d"
    ...
    d.D(a, b, c)
*/
package d
