# xmlformat
Streaming formatting of xml files

A simple streaming XML formatter. Not as fast and versatile as xmllint, but doesn't run into memory issues for very big files.
It obviously also works for small files.

Xmlformat is useful for example to pre-format huge files before piping them through some filter. 

All whitespace-only tokens are assumed to be 'ignorable'. It is not possible to use a schema.

# Build from source

```bash
go get -u github.com/bertbaron/xmlformat
```

# Usage examples

```bash
xmlformat -indent=tab -outfile huge-formatted.xml huge.xml   

or streaming directly from archive:

```bash
gunzip -c huge.xml.gz | xmlformat | grep '<SomeTag>' | sort -u | wc -l
```

# Compatibility

Due to limitations of the go endocing/xml library files starting with non ASCII-compatible encodings like UTF-16 are not supported. Other input encodings are supported. The output will always be UTF-8 encoded.

