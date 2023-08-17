package route

import (
	"bytes"
	"fmt"
	"github.com/Senhnn/shlhttp/app"
)

// MethodTrees 方法数
type MethodTrees []*router

func (m MethodTrees) get(method string) *router {
	for _, tree := range m {
		if tree.method == method {
			return tree
		}
	}
	return nil
}

// path中的获取参数个数
func getParamsNum(path string) int {
	var ret int
	b := []byte(path)
	ret += bytes.Count(b, []byte(":"))
	ret += bytes.Count(b, []byte("*"))
	return ret
}

type router struct {
	method     string
	root       *node
	hasHandler map[string]bool
}

// radix tree的子节点
type children []*node

type node struct {
	kind       kind
	label      byte
	prefix     string
	parent     *node
	children   children
	parentPath string
	paramNames []string
	handlers   app.HandlersChain
	paramChild *node
	anyChild   *node
	isLeaf     bool // 是否是叶子，如果是叶子则children不为空
}

// 检查path
func checkPathValid(path string) {
	if path == "" {
		panic("path is empty")
	}
	if path[0] != '/' {
		panic("first char should be '/'")
	}
	for i, c := range path {
		switch c {
		case ':': // : 只捕获后面的一个参数
			if (i < len(path)-1 && path[i+1] == '/') || i == len(path)-1 {
				panic("wildcards mut be named with no-empty string in path:" + path)
			}
			i++
			for ; i < len(path)-1; i++ {
				if path[i] == ':' || path[i] == '*' {
					panic("in one path only should be held one wildcard, multi in path:" + path)
				}
			}
		case '*': // * 捕获‘/’以及‘/’后面的所有参数
			if i == len(path)-1 {
				panic("wildcards must be named with a non-empty name in path:" + path)
			}
			if i > 0 && path[i-1] != '/' {
				panic(" no / before wildcard in path:" + path)
			}
			for ; i < len(path); i++ {
				if path[i] == '/' {
					panic("catch-all routes are only allowed at the end of the path in path:" + path)
				}
			}
		}
	}
}

// 向一个router中加入新的路由规则
func (r *router) addRoute(path string, chain app.HandlersChain) {
	// 检查path是否合法
	checkPathValid(path)
	var paramNames []string
	var rawPath string = path

	if chain == nil {
		panic(fmt.Sprintf("addRoute path:%s chain is nil", path))
	}

	// 添加非静态路由的前端静态部分
	for i1, i2 := 0, len(path); i1 < i2; i1++ {
		if path[i1] == ':' {
			j := i1 + 1

		}
	}
}

func (r *router) insert(path string, chain app.HandlersChain, t kind, ppath string, paramNames []string) {
	currentNode := r.root
	if currentNode == nil {
		panic("router.root should not empty")
	}
	search := path

	for {
		searchLen := len(search)
		prefixLen := len(currentNode.prefix)
		lcpLen := 0

		max := prefixLen
		if searchLen < max {
			max = searchLen
		}

		// 从头开始匹配公共部分
		for ; lcpLen < max && search[lcpLen] == currentNode.prefix[lcpLen]; lcpLen++ {
		}

		if lcpLen == 0 {
			// 此时为根节点
			currentNode.label = search[0]

		} else if lcpLen < prefixLen {

		} else if lcpLen < searchLen {

		}
	}
}
