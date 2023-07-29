package generator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/google/go-github/v53/github"
)

func NewGenerator(baseDirectory, version string) *Generator {
	return &Generator{
		baseDirectory: baseDirectory,
		version:       version,
	}
}

type Generator struct {
	baseDirectory string
	version       string
}

func (g *Generator) Handle() {
	g.run(g.version)
}

func (g *Generator) path(file string) string {
	return filepath.Join(g.baseDirectory, file)
}

func (g *Generator) run(version string) {
	// Ensure the wrapper directory exists
	err := os.MkdirAll(g.path("wrapper"), os.ModePerm)
	g.panicOnError(err)

	// Create the base client file
	g.createClientFile(version)

	client := github.NewClient(nil)

	t := reflect.TypeOf(client).Elem() // get type of struct

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Check if the field is a struct and is exported
		if field.Type.Kind() == reflect.Ptr && field.PkgPath == "" {
			// Create a new file in the wrapper directory
			file, err := os.Create(g.path(filepath.Join("wrapper", fmt.Sprintf("%s.go", strings.ToLower(field.Name)))))
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer file.Close()

			// Write the package name and imports
			_, err = io.WriteString(file, "package wrapper\n\n")
			g.panicOnError(err)
			_, err = io.WriteString(file, fmt.Sprintf("import \"github.com/google/go-github/%s/github\"\n\n", version))
			g.panicOnError(err)
			// Write the method
			_, err = io.WriteString(file, fmt.Sprintf("func (c *Client) %s() *github.%sService {\n", field.Name, field.Name))
			g.panicOnError(err)
			_, err = io.WriteString(file, fmt.Sprintf("\treturn c.client.%s\n", field.Name))
			g.panicOnError(err)
			_, err = io.WriteString(file, "}\n")
			g.panicOnError(err)
		}
	}
}

func (g *Generator) createClientFile(version string) {
	file, err := os.Create(g.path("wrapper/client.go"))
	g.panicOnError(err)
	defer file.Close()

	// Write the package name and imports
	_, err = io.WriteString(file, "package wrapper\n\n")
	g.panicOnError(err)
	_, err = io.WriteString(file, fmt.Sprintf("import \"github.com/google/go-github/%s/github\"\n\n", version))
	g.panicOnError(err)

	// Write the Client struct
	_, err = io.WriteString(file, "type Client struct {\n")
	g.panicOnError(err)
	_, err = io.WriteString(file, "\tclient *github.Client\n")
	g.panicOnError(err)
	_, err = io.WriteString(file, "}\n\n")
	g.panicOnError(err)

	// Write the NewClient function
	_, err = io.WriteString(file, "func NewClient() *Client {\n")
	g.panicOnError(err)
	_, err = io.WriteString(file, "\treturn &Client{client: github.NewClient(nil)}\n")
	g.panicOnError(err)
	_, err = io.WriteString(file, "}\n")
	g.panicOnError(err)
}

func (g *Generator) panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
