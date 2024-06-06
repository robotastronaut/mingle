package mpackage

type Package interface {
	Read() error
	Write() error
}

/*

	Readme
	Icon
	Deps?
	outputFile??
	If readme exists, will replace MFile description

	compile to ./build
	config.lua



Package Structure


/
mfile
readme.md
.gitignore
.gitattributes
Taskfile.yml
src/
	resources/   <--- This is moved raw. No PKGNAME replace, etc.
*/
