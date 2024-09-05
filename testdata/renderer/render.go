package main

import (
	"github.com/anekkanti/go-codegen/pkg/types"
	"github.com/dave/jennifer/jen"
)

func RenderInterfaceImpl(
	ifc *types.GoInterface, // The interface to render implementations for
	file *jen.File, // The jen file to write the implementation to
) error {
	// Render the interface implementation
	file.Comment("handler for " + ifc.Name)
	file.Type().Id(ifc.Name + "Handler").Struct()

	// Render the constructor
	file.Comment("create a new " + ifc.Name + "Handler")
	file.Func().Id("New"+ifc.Name+"Handler").Params().Params(jen.Qual(ifc.File.ImportPath, ifc.Name), jen.Error()).Block(
		jen.Return(jen.Op("&").Id(ifc.Name+"Handler").Block(), jen.Nil()),
	)
	file.Line()
	return nil
}

func RenderInterfaceMethodImpl(
	ifc *types.GoInterface, // The interface to render method implementations for
	method *types.GoMethod, // The method to render implementation for
	file *jen.File, // The jen file to write the implementation to
) error {
	// Render the method signature
	file.Func().Params(
		jen.Id("h").Op("*").Id(ifc.Name+"Handler"),
	).Id(method.Name).Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("req").Id(method.Params[1].Type),
	).Params(
		jen.Id(method.Results[0].Type),
		jen.Error(),
	).Block(
		jen.Return(
			jen.Nil(),
			jen.Qual("google.golang.org/grpc/status", "Errorf").Call(jen.Qual("google.golang.org/grpc/codes", "Unimplemented").Op(",").Lit("method not implemented"))),
	)
	file.Line()
	return nil
}
