package pkg

import (
	"fmt"
	"path/filepath"
	"plugin"

	"github.com/anekkanti/go-codegen/pkg/types"
	"github.com/dave/jennifer/jen"
)

const (
	RenderInterfaceImplFuncName       = "RenderInterfaceImpl"
	RenderInterfaceMethodImplFuncName = "RenderInterfaceMethodImpl"
)

type (
	RenderInterfaceImplFunc func(
		ifc *types.GoInterface, // The interface to render implementations for
		file *jen.File, // The jen file to write the implementation to
	) error

	RenderInterfaceMethodImplFunc func(
		ifc *types.GoInterface, // The interface to render method implementations for
		method *types.GoMethod, // The method to render implementation for
		file *jen.File, // The jen file to write the implementation to
	) error
)

func renderIfc(p *plugin.Plugin, ifc *types.GoInterface, file *jen.File) error {
	f, err := p.Lookup(RenderInterfaceImplFuncName)
	if err == nil {
		fn, ok := f.(func(*types.GoInterface, *jen.File) error)
		if !ok {
			return fmt.Errorf("invalid render function %s", RenderInterfaceImplFuncName)
		}
		if err := fn(ifc, file); err != nil {
			return fmt.Errorf("failed to render interface: %w", err)
		}
	}

	f, err = p.Lookup(RenderInterfaceMethodImplFuncName)
	if err == nil {
		fnm, ok := f.(func(*types.GoInterface, *types.GoMethod, *jen.File) error)
		if !ok {
			return fmt.Errorf("invalid render function %s", RenderInterfaceMethodImplFuncName)
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

	for _, f := range files {

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
