package pkg

import (
	"fmt"
	"path/filepath"
	"plugin"

	"github.com/anekkanti/go-codegen/pkg/types"
	"github.com/dave/jennifer/jen"
)

const (
	RenderFuncName                = "Render"
	RenderFileFuncName            = "RenderFile"
	RenderInterfaceFuncName       = "RenderInterface"
	RenderInterfaceMethodFuncName = "RenderInterfaceMethod"
)

type (
	// Will be invoked only once per generation
	RenderFunc func(
		files []*types.GoFile, // The files to render implementations for
		file *jen.File, // The jen file to write the implementation to
	) error

	// Will be invoked for each file parsed
	RenderFileFunc func(
		f *types.GoFile, // The file to render implementations for
		jenFile *jen.File, // The jen file to write the implementation to
	) error

	// Will be invoked for each interface parsed in each file
	RenderInterfaceFunc func(
		ifc *types.GoInterface, // The interface to render implementations for
		file *jen.File, // The jen file to write the implementation to
	) error

	// Will be invoked for each method parsed in each interface in each file
	RenderInterfaceMethodFunc func(
		ifc *types.GoInterface, // The interface to render method implementations for
		method *types.GoMethod, // The method to render implementation for
		file *jen.File, // The jen file to write the implementation to
	) error
)

func render(p *plugin.Plugin, files []*types.GoFile, file *jen.File) error {
	f, err := p.Lookup(RenderFuncName)
	if err == nil {
		fn, ok := f.(func([]*types.GoFile, *jen.File) error)
		if !ok {
			return fmt.Errorf("invalid render function %s", RenderFuncName)
		}
		if err := fn(files, file); err != nil {
			return fmt.Errorf("failed to render files: %w", err)
		}
	}
	return nil
}

func renderFile(p *plugin.Plugin, file *types.GoFile, jenFile *jen.File) error {
	f, err := p.Lookup(RenderFileFuncName)
	if err == nil {
		fn, ok := f.(func(*types.GoFile, *jen.File) error)
		if !ok {
			return fmt.Errorf("invalid render function %s", RenderFileFuncName)
		}
		if err := fn(file, jenFile); err != nil {
			return fmt.Errorf("failed to render file: %w", err)
		}
	}
	return nil
}

func renderIfc(p *plugin.Plugin, ifc *types.GoInterface, file *jen.File) error {
	f, err := p.Lookup(RenderInterfaceFuncName)
	if err == nil {
		fn, ok := f.(func(*types.GoInterface, *jen.File) error)
		if !ok {
			return fmt.Errorf("invalid render function %s", RenderInterfaceFuncName)
		}
		if err := fn(ifc, file); err != nil {
			return fmt.Errorf("failed to render interface: %w", err)
		}
	}

	f, err = p.Lookup(RenderInterfaceMethodFuncName)
	if err == nil {
		fnm, ok := f.(func(*types.GoInterface, *types.GoMethod, *jen.File) error)
		if !ok {
			return fmt.Errorf("invalid render function %s", RenderInterfaceMethodFuncName)
		}
		for _, method := range ifc.Methods {
			if err := fnm(ifc, method, file); err != nil {
				return fmt.Errorf("failed to render interface method: %w", err)
			}
		}
	}
	return nil
}

func Render(
	pluginPath string,
	outputPath string,
	nameFilters []string,
	files []*types.GoFile,
) error {
	// load the plugin
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return fmt.Errorf("failed to open plugin %s: %w", pluginPath, err)
	}

	outPkg, err := parseImportPath(filepath.Dir(outputPath))
	if err != nil {
		return fmt.Errorf("failed to parse output package: %w", err)
	}
	// create the jen file
	file := jen.NewFilePath(outPkg)

	// Render the files
	if err := render(p, files, file); err != nil {
		return fmt.Errorf("failed to render: %w", err)
	}

	for _, f := range files {

		// Render the file
		if err := renderFile(p, f, file); err != nil {
			return fmt.Errorf("failed to render file: %w", err)
		}

		// Iterate over the interfaces
		for _, ifc := range f.Interfaces {
			if len(nameFilters) > 0 {
				found := false
				for _, name := range nameFilters {
					if name == ifc.Name {
						found = true
						break
					}
				}
				if !found {
					continue
				}
			}
			// Render the interface
			if err := renderIfc(p, ifc, file); err != nil {
				return fmt.Errorf("failed to render interface: %w", err)
			}
		}

		// TODO: Render the structs
		// TODO: Render the methods
	}

	return file.Save(outputPath)
}
