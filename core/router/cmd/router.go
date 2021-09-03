package main

import (
	"errors"
	"fmt"
	"github.com/boshangad/go-api/utils"
	"github.com/boshangad/go-api/utils/str"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"time"
	"unicode"
)
// 正在表达式 - 解析路由配置器
var rangeCompile = regexp.MustCompile(`(?i)@ROUTE\s+([\w\/\:\*\-\_]+)(?:\s+\[([\w\t ,]+)\])?`)

// 一个函数表明是一个路由组，生成函数
type routeRegister struct {
	// 控制器名称
	FunName string
	// 注册函数名称
	Name string
	// 组的路由地址
	Path string
	// 可访问的类型方法
	Methods []string
	// 入参
	Params map[string]string
}

// 一个函数表明是一个路由组，生成函数
type routeRegisterFun struct {
	// 函数名称
	Name string
	// 组的路由地址
	Path string
	// 注册的路由方法
	Items []routeRegister
}

// Router 路由结构
type Router struct {
	RouterDir string
	Filename string
	ControllerPath string
}

// 创建Ast方法函数
func (r *Router) buildAstFunDecl(funcDecl *ast.FuncDecl) (*routeRegister, error) {
	// 判断首字母是否大写，也就是公共的方法
	if unicode.IsUpper([]rune(funcDecl.Name.Name)[0]) == false {
		return nil, fmt.Errorf("[%s] controller action must be public", funcDecl.Name.Name)
	}
	// 获取方法所属类的名称，只允许存在一个类名称
	declRecvList := funcDecl.Recv.List
	if len(declRecvList) != 1 {
		return nil, fmt.Errorf("func %s formatter error", funcDecl.Name.Name)
	}
	objectName := ""
	declRecvType, ok := declRecvList[0].Type.(*ast.Ident)
	if ok {
		objectName = declRecvType.Name
	} else {
		declRecvType, ok := declRecvList[0].Type.(*ast.StarExpr)
		if !ok {
			return nil, fmt.Errorf("func %s formatter error, not is ast.StarExpr", funcDecl.Name.Name)
		}
		recvTypeX, ok := declRecvType.X.(*ast.Ident)
		if !ok {
			return nil, fmt.Errorf("func %s formatter error, not is ast.Ident", funcDecl.Name.Name)
		}
		objectName = recvTypeX.Name
	}

	objectPath := ""
	// 获取注释
	var funcMethods []string
	anyMethod := "Any"
	if funcDecl.Doc != nil && len(funcDecl.Doc.List) > 0 {
		docList := funcDecl.Doc.List
		for _, doc := range docList {
			// 忽略掉没有@符号的数据
			if strings.Index(doc.Text, "@") == -1 {
				continue
			}
			methods := rangeCompile.FindStringSubmatch(doc.Text)
			methodsLen := len(methods)
			if methodsLen < 2 {
				methods = []string{"", "", anyMethod}
			} else if methodsLen == 2 {
				methods = append(methods, anyMethod)
			}
			objectPath = methods[1]
			if objectPath == "-" {
				return nil, errors.New("")
			}
			methods = strings.Split(strings.ToUpper(strings.Replace(methods[2], " ", "", -1)), ",")
			if methods == nil || len(methods) < 1 {
				methods = []string{anyMethod}
			}
			sort.Strings(methods)
			for _, method := range methods {
				if method == "" {
					continue
				} else if method == strings.ToUpper(anyMethod) {
					funcMethods= []string{anyMethod}
					break
				}
				funcMethods = append(funcMethods, method)
			}
		}
	}
	if len(funcMethods) < 1 {
		funcMethods = []string{anyMethod}
	}
	if objectPath == "" {
		objectPath = str.SnakeString(funcDecl.Name.Name)
	}
	// 获取方法入参
	if funcDecl.Type != nil && funcDecl.Type.Params != nil {
		typeList := funcDecl.Type.Params.List
		for _, funType := range typeList {
			ft := reflect.TypeOf(funType.Type).Name()
			if ft == "*ast.Ident" {
				_ = funType.Type.(*ast.Ident).Name
			} else if ft == "*ast.Ellipsis" {
				ftc := funType.Type.(*ast.Ellipsis)
				_ = ftc.Elt.(*ast.Ident).Name
				_ = true
			}
			if funType.Names != nil {
				for _, fn := range funType.Names {
					_ = fn.Name
				}
			}
		}
	}

	return &routeRegister{
		FunName: objectName,
		Name: funcDecl.Name.Name,
		Path: objectPath,
		Methods: funcMethods,
	}, nil
}

// 获取GO源码的AST码
func (r *Router) getAstPackages(dir string) (map[string]*ast.Package, error) {
	if !utils.IsDir(dir) {
		return nil, fmt.Errorf("%s目录不存在", dir)
	}
	fest := token.NewFileSet()
	f, err := parser.ParseDir(fest, dir, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	//ast.Print(fest, f)
	//log.Fatal("")
	return f, nil
}

// 生成路由配置文件
func (r *Router) buildAstPackages(packages map[string]*ast.Package) {
	// 开始循环获取值
	for _,pkg := range packages {
		routeRegisterFuns := make(map[string]routeRegisterFun)
		files := pkg.Files
		for _, file := range files {
			decls := file.Decls
			for _, decl := range decls {
				declType := reflect.TypeOf(decl).String()
				// 如果是一个函数结构函数结构
				if declType == "*ast.FuncDecl" {
					funcDecl := decl.(*ast.FuncDecl)
					routeReg, err := r.buildAstFunDecl(funcDecl)
					if err != nil {
						fmt.Println(err)
						continue
					}

					regItem, ok := routeRegisterFuns[routeReg.FunName]
					if !ok {
						// 如果不存在要创建的方法名
						routeRegisterFuns[routeReg.FunName] = routeRegisterFun{
							Name: routeReg.FunName,
							Path: str.SnakeString(routeReg.FunName),
							Items: []routeRegister{
								*routeReg,
							},
						}
					} else {
						// 如果有创建的方法
						regItem.Items = append(regItem.Items, *routeReg)
						routeRegisterFuns[routeReg.FunName] = regItem
					}
				}
			}
		}
		// 这里就要开始注册数据了
		if len(routeRegisterFuns) < 1{
			continue
		}
		// 定义模板,并写入模板
		genTemplate, err := template.New("template").Parse(getTemplateStr())
		if err != nil {
			log.Fatalln(err)
		}
		packageName := filepath.Base(r.RouterDir)
		fileName := filepath.Base(r.Filename)
		if fileName == "controllers" {
			fileName = "routers"
		}
		filePath := fmt.Sprintf("./%s/%s_gen.go", packageName, fileName)
		f, err := os.Create(filePath)
		if err != nil {
			log.Fatalln(err)
		}
		t := templateStruct{
			PackageName: filepath.Dir(filePath),
			QuoteName: pkg.Name,
			Functions: routeRegisterFuns,
		}
		err = genTemplate.Execute(f, t)
		//err = genTemplate.Execute(f, struct {
		//	PackageName string
		//	QuoteName string
		//	Funs map[string]routeRegisterFun
		//}{filepath.Dir(filePath), pkg.Name, routeRegisterFuns})
		_ = f.Close()
		if err != nil {
			log.Fatalln(err)
		}

	}
	fmt.Printf("\n%s 完成操作\n", time.Now().Format("2006年01月02日15:04:05"))
}

// Build 生成路由配置器
func (r *Router) Build(path string, relativePath string)  {
	dir := r.ControllerPath + "/" + path
	fileInfoList, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln("目录文件不存在", err)
		return
	}
	for _, fileInfo := range fileInfoList {
		if fileInfo.IsDir() {
			r.Build(path + fileInfo.Name(), relativePath)
			continue
		}
	}
	fmt.Println("正在加载文件：" + dir)
	packages, e := r.getAstPackages(dir)
	if e != nil {
		fmt.Println(e)
		return
	}
	r.Filename = dir
	r.buildAstPackages(packages)
}