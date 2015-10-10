# xmlformat
Streaming formatting of xml files

A simple streaming XML formatter. Not as fast and versatile as xmllint, but doesn't run into memory issues for very big files.
It obviously also works for small files.

Xmlformat is useful for example to pre-format huge files before piping them through some filter. 

All whitespace-only tokens are assumed to be 'ignorable'. It is not possible to use a schema.

# Usage examples

xmlformat -indent=tab -outfile huge-formatted.xml huge.xml   

or streaming directly from archive:

gunzip -c huge.xml.gz | xmlformat | grep '<SomeTag>' | sort -u | wc -l