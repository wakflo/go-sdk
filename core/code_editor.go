package core

type CodeFrameworkProps struct {
	Files         map[string]string `json:"files"`
	SourcePath    string            `json:"sourcePath"`
	EntryPath     string            `json:"entryPath"`
	Framework     CodeFramework     `json:"framework"`
	InternalFiles []string          `json:"internalFiles"`
}

var CodeFrameworkDefaults = map[CodeFramework]CodeFrameworkProps{
	CodeFrameworkNode: {
		Files: map[string]string{
			"index.ts": "" +
				"import { Context } from '@wakflo/flow';\n" +
				"\n" +
				"export function execute(context: FlowContext) {\n" +
				"\n" +
				"     return {};\n" +
				"}\n",
			"package.json": "{\n" +
				"  \"name\": \"node-code\",\n" +
				"  \"dependencies\": {\n" +
				"    \"@wakflo/flow\": \"^0.0.1\"\n" +
				"  }\n" +
				"}\n",
		},
		SourcePath: "src",
		EntryPath:  "index.ts",
		Framework:  CodeFrameworkNode,
		InternalFiles: []string{
			"package.json",
			"index.ts",
		},
	},
	CodeFrameworkDeno: {
		Files: map[string]string{
			"index.ts": "" +
				"import { Context } from '@wakflo/flow';\n" +
				"\n" +
				"export function execute(context: FlowContext) {\n" +
				"\n" +
				"return true;\n" +
				"}\n",
			"deno.json": "{\n" +
				"  \"imports\": {\n" +
				"  }\n" +
				"}\n",
		},
		SourcePath: "src",
		EntryPath:  "index.ts",
		Framework:  CodeFrameworkDeno,
		InternalFiles: []string{
			"deno.json",
			"index.ts",
		},
	},
	CodeFrameworkGoLang: {
		Files: map[string]string{
			"lib.go": "" +
				"package main\n" +
				"\n" +
				"func Execute() any {\n" +
				"\n" +
				"  return true\n" +
				"}\n",
			"go.mod": "" +
				"module main\n" +
				"\n" +
				"go 1.23.4\n",
		},
		SourcePath: "src",
		EntryPath:  "lib.go",
		Framework:  CodeFrameworkGoLang,
		InternalFiles: []string{
			"go.mod",
			"lib.go",
		},
	},
}
