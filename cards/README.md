# Setting up VSCODE to work with Go and Go Modules

1. Install GoLang via Brew
`brew install go`
2. Set the GOROOT path in bash_profile
`export GOROOT="$(brew --prefix golang)/libexec"`
3. Enable GO Modules in bash_profile
`export GO111MODULE=on`
4. Install GO extension for visual studio code
> Name: Go
Id: ms-vscode.go
Description: Rich Go language support for Visual Studio Code
Version: 0.11.9
Publisher: Microsoft
VS Marketplace Link: https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go`
1. Install all GO Tools for visual studio code.
    1. CTRL + SHIFT + P
    2. Search for GO: Install/Update Tools
    3. Select all tools
    4. VSCode Terminal will show progress of successful installation
2. In the working directory execute `go mod init {name_of_package}`. This will generate a `go.mod` file.
3. In the working directory create a vscode workspace. File > Save as Workspace. This will generate a `{name_of_project}.code-workspace` file.
4. In the workspace file. Copy and paste the following settings underneath the file array.
```
"settings": {
		"go.gotoSymbol.includeGoroot": true,
		"go.gotoSymbol.includeImports": true,
		"go.useLanguageServer": true,
		"go.lintOnSave": "package",
		"go.languageServerExperimentalFeatures": {
			"format": false,
			"autoComplete": true,
			"rename": true,
			"goToDefinition": true,
			"hover": true,
			"signatureHelp": true,
			"goToTypeDefinition": true,
			"goToImplementation": true,
			"documentSymbols": true,
			"workspaceSymbols": true,
			"findReferences": true,
			"diagnostics": true,
			"documentLink": true,
			"highlight": true
		},
		"go.lintFlags": [
			"--fast"
		],
		"[go]": {
			"editor.snippetSuggestions": "none",
			"editor.formatOnSave": true,
			"editor.codeActionsOnSave": {
				"source.organizeImports": false
			}
		},
		"files.eol": "\n", // formatting only supports LF line endings
	}
```
