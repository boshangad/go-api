package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
)
// 正在表达式 - 解析路由配置器
//var rangeCompile = regexp.MustCompile(`(?i)@router\s+([\w\/\:\*\-\_]+)(?:\s+\[([\w\t ,]+)\])?`)

type genStruct struct {
	fest *token.FileSet
	// 访问包路径
	PackagePath string

}

func (g genStruct) getFunctionByFileDecl(funcDecl *ast.FuncDecl) {

}

func (g genStruct) getFunctions(pkgs map[string]*ast.Package) {
	//var packageName string = ""
	for _,pkg := range pkgs {
		//packageName = pkg.Name
		files := pkg.Files
		for _, file := range files {
			decls := file.Decls
			for _, decl := range decls {
				// 判断是否是一个Func的配置数据
				funcDecl, ok := decl.(*ast.FuncDecl)
				if ok {
					g.getFunctionByFileDecl(funcDecl)
				}

				fmt.Println(funcDecl)
			}
		}

	}
}

func (g genStruct) getPackages() (pkgs map[string]*ast.Package) {
	fest := token.NewFileSet()
	pkgs, err := parser.ParseDir(fest, g.PackagePath, func(info fs.FileInfo) bool {
		return true
	}, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
		return
	}
	g.fest = fest
	//ast.Print(fest, f)
	return
}
