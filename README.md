# FIleLinker
A simple GO library to merge and unmerge multiple files. Used to include parts of code/text into other files without thinking about dependencies.
Bringing the love of PHP's `include` to every system.

## Dialect Specification
Files are scanned for import flags, these must be on their own line. Example of an import flag:
```
//import<file>
```
or, alternatively
```
#import<file>
```
depending on the language it's used in.
The file can also include a relative path, like
```
//import<utils/rest.js>
```
This means that the file `rest.js` will be loaded from the `utils` directory, scanned for imports in itself (since it could be a tree of other files). get compiled and then replaces the import statement in the source file with
```
//FL:START:utils/rest.js
compiled file content...
//FL:END
```

The first line of the file should always be a root note containing the filename, example
```
//FL:ROOT:main.js
```

###Terminology:
 - **Root File** The top-level output file after compiling
 - **Source File** One of the imported files by the root file or another source file.