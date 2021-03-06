package web

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
	"github.com/W3-Engineers-Ltd/Radiant/core/utils"
)

var (
	radiantTplFuncMap             = make(template.FuncMap)
	radicalViewPathTemplateLocked = false
	// radicalViewPathTemplates caching map and supported template file extensions per view
	radicalViewPathTemplates = make(map[string]map[string]*template.Template)
	templatesLock            sync.RWMutex
	// radicalTemplateExt stores the template extension which will build
	radicalTemplateExt = []string{"tpl", "html", "gohtml"}
	// radicalTemplatePreprocessors stores associations of extension -> preprocessor handler
	radicalTemplateEngines = map[string]templatePreProcessor{}
	radicalTemplateFS      = defaultFSFunc
)

// ExecuteTemplate applies the template with name  to the specified data object,
// writing the output to wr.
// A template will be executed safely in parallel.
func ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return ExecuteViewPathTemplate(wr, name, BConfig.WebConfig.ViewsPath, data)
}

// ExecuteViewPathTemplate applies the template with name and from specific viewPath to the specified data object,
// writing the output to wr.
// A template will be executed safely in parallel.
func ExecuteViewPathTemplate(wr io.Writer, name string, viewPath string, data interface{}) error {
	if BConfig.RunMode == DEV {
		templatesLock.RLock()
		defer templatesLock.RUnlock()
	}
	if radicalTemplates, ok := radicalViewPathTemplates[viewPath]; ok {
		if t, ok := radicalTemplates[name]; ok {
			var err error
			if t.Lookup(name) != nil {
				err = t.ExecuteTemplate(wr, name, data)
			} else {
				err = t.Execute(wr, data)
			}
			if err != nil {
				logs.Trace("template Execute err:", err)
			}
			return err
		}
		panic("can't find templatefile in the path:" + viewPath + "/" + name)
	}
	panic("Unknown view path:" + viewPath)
}

func init() {
	radiantTplFuncMap["dateformat"] = DateFormat
	radiantTplFuncMap["date"] = Date
	radiantTplFuncMap["compare"] = Compare
	radiantTplFuncMap["compare_not"] = CompareNot
	radiantTplFuncMap["not_nil"] = NotNil
	radiantTplFuncMap["not_null"] = NotNil
	radiantTplFuncMap["substr"] = Substr
	radiantTplFuncMap["html2str"] = HTML2str
	radiantTplFuncMap["str2html"] = Str2html
	radiantTplFuncMap["htmlquote"] = Htmlquote
	radiantTplFuncMap["htmlunquote"] = Htmlunquote
	radiantTplFuncMap["renderform"] = RenderForm
	radiantTplFuncMap["assets_js"] = AssetsJs
	radiantTplFuncMap["assets_css"] = AssetsCSS
	radiantTplFuncMap["config"] = GetConfig
	radiantTplFuncMap["map_get"] = MapGet

	// Comparisons
	radiantTplFuncMap["eq"] = eq // ==
	radiantTplFuncMap["ge"] = ge // >=
	radiantTplFuncMap["gt"] = gt // >
	radiantTplFuncMap["le"] = le // <=
	radiantTplFuncMap["lt"] = lt // <
	radiantTplFuncMap["ne"] = ne // !=

	radiantTplFuncMap["urlfor"] = URLFor // build a URL to match a Controller and it's method
}

// AddFuncMap let user to register a func in the template.
func AddFuncMap(key string, fn interface{}) error {
	radiantTplFuncMap[key] = fn
	return nil
}

type templatePreProcessor func(root, path string, funcs template.FuncMap) (*template.Template, error)

type templateFile struct {
	root  string
	files map[string][]string
}

// visit will make the paths into two part,the first is subDir (without tf.root),the second is full path(without tf.root).
// if tf.root="views" and
// paths is "views/errors/404.html",the subDir will be "errors",the file will be "errors/404.html"
// paths is "views/admin/errors/404.html",the subDir will be "admin/errors",the file will be "admin/errors/404.html"
func (tf *templateFile) visit(paths string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	if f.IsDir() || (f.Mode()&os.ModeSymlink) > 0 {
		return nil
	}
	if !HasTemplateExt(paths) {
		return nil
	}

	replace := strings.NewReplacer("\\", "/")
	file := strings.TrimLeft(replace.Replace(paths[len(tf.root):]), "/")
	subDir := filepath.Dir(file)

	tf.files[subDir] = append(tf.files[subDir], file)
	return nil
}

// HasTemplateExt return this path contains supported template extension of radiant or not.
func HasTemplateExt(paths string) bool {
	for _, v := range radicalTemplateExt {
		if strings.HasSuffix(paths, "."+v) {
			return true
		}
	}
	return false
}

// AddTemplateExt add new extension for template.
func AddTemplateExt(ext string) {
	for _, v := range radicalTemplateExt {
		if v == ext {
			return
		}
	}
	radicalTemplateExt = append(radicalTemplateExt, ext)
}

// AddViewPath adds a new path to the supported view paths.
// Can later be used by setting a controller ViewPath to this folder
// will panic if called after radiant.Run()
func AddViewPath(viewPath string) error {
	if radicalViewPathTemplateLocked {
		if _, exist := radicalViewPathTemplates[viewPath]; exist {
			return nil // Ignore if viewpath already exists
		}
		panic("Can not add new view paths after radiant.Run()")
	}
	radicalViewPathTemplates[viewPath] = make(map[string]*template.Template)
	return BuildTemplate(viewPath)
}

func lockViewPaths() {
	radicalViewPathTemplateLocked = true
}

// BuildTemplate will build all template files in a directory.
// it makes radiant can render any template file in view directory.
func BuildTemplate(dir string, files ...string) error {
	var err error
	fs := radicalTemplateFS()
	f, err := fs.Open(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return errors.New("dir open err")
	}
	defer f.Close()

	radicalTemplates, ok := radicalViewPathTemplates[dir]
	if !ok {
		panic("Unknown view path: " + dir)
	}
	self := &templateFile{
		root:  dir,
		files: make(map[string][]string),
	}
	err = Walk(fs, dir, self.visit)
	if err != nil {
		fmt.Printf("Walk() returned %v\n", err)
		return err
	}
	buildAllFiles := len(files) == 0
	for _, v := range self.files {
		for _, file := range v {
			if buildAllFiles || utils.InSlice(file, files) {
				templatesLock.Lock()
				ext := filepath.Ext(file)
				var t *template.Template
				if len(ext) == 0 {
					t, err = getTemplate(self.root, fs, file, v...)
				} else if fn, ok := radicalTemplateEngines[ext[1:]]; ok {
					t, err = fn(self.root, file, radiantTplFuncMap)
				} else {
					t, err = getTemplate(self.root, fs, file, v...)
				}
				if err != nil {
					logs.Error("parse template err:", file, err)
					templatesLock.Unlock()
					return err
				}
				radicalTemplates[file] = t
				templatesLock.Unlock()
			}
		}
	}
	return nil
}

func getTplDeep(root string, fs http.FileSystem, file string, parent string, t *template.Template) (*template.Template, [][]string, error) {
	var fileAbsPath string
	var rParent string
	var err error
	if strings.HasPrefix(file, "../") {
		rParent = filepath.Join(filepath.Dir(parent), file)
		fileAbsPath = filepath.Join(root, filepath.Dir(parent), file)
	} else {
		rParent = file
		fileAbsPath = filepath.Join(root, file)
	}
	f, err := fs.Open(fileAbsPath)
	if err != nil {
		panic("can't find template file:" + file)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, [][]string{}, err
	}
	t, err = t.New(file).Parse(string(data))
	if err != nil {
		return nil, [][]string{}, err
	}
	reg := regexp.MustCompile(BConfig.WebConfig.TemplateLeft + "[ ]*template[ ]+\"([^\"]+)\"")
	allSub := reg.FindAllStringSubmatch(string(data), -1)
	for _, m := range allSub {
		if len(m) == 2 {
			tl := t.Lookup(m[1])
			if tl != nil {
				continue
			}
			if !HasTemplateExt(m[1]) {
				continue
			}
			_, _, err = getTplDeep(root, fs, m[1], rParent, t)
			if err != nil {
				return nil, [][]string{}, err
			}
		}
	}
	return t, allSub, nil
}

func getTemplate(root string, fs http.FileSystem, file string, others ...string) (t *template.Template, err error) {
	t = template.New(file).Delims(BConfig.WebConfig.TemplateLeft, BConfig.WebConfig.TemplateRight).Funcs(radiantTplFuncMap)
	var subMods [][]string
	t, subMods, err = getTplDeep(root, fs, file, "", t)
	if err != nil {
		return nil, err
	}
	t, err = _getTemplate(t, root, fs, subMods, others...)

	if err != nil {
		return nil, err
	}
	return
}

func _getTemplate(t0 *template.Template, root string, fs http.FileSystem, subMods [][]string, others ...string) (t *template.Template, err error) {
	t = t0
	for _, m := range subMods {
		if len(m) == 2 {
			tpl := t.Lookup(m[1])
			if tpl != nil {
				continue
			}
			// first check filename
			for _, otherFile := range others {
				if otherFile == m[1] {
					var subMods1 [][]string
					t, subMods1, err = getTplDeep(root, fs, otherFile, "", t)
					if err != nil {
						logs.Trace("template parse file err:", err)
					} else if len(subMods1) > 0 {
						t, err = _getTemplate(t, root, fs, subMods1, others...)
					}
					break
				}
			}
			// second check define
			for _, otherFile := range others {
				var data []byte
				fileAbsPath := filepath.Join(root, otherFile)
				f, err := fs.Open(fileAbsPath)
				if err != nil {
					f.Close()
					logs.Trace("template file parse error, not success open file:", err)
					continue
				}
				data, err = ioutil.ReadAll(f)
				f.Close()
				if err != nil {
					logs.Trace("template file parse error, not success read file:", err)
					continue
				}
				reg := regexp.MustCompile(BConfig.WebConfig.TemplateLeft + "[ ]*define[ ]+\"([^\"]+)\"")
				allSub := reg.FindAllStringSubmatch(string(data), -1)
				for _, sub := range allSub {
					if len(sub) == 2 && sub[1] == m[1] {
						var subMods1 [][]string
						t, subMods1, err = getTplDeep(root, fs, otherFile, "", t)
						if err != nil {
							logs.Trace("template parse file err:", err)
						} else if len(subMods1) > 0 {
							t, err = _getTemplate(t, root, fs, subMods1, others...)
							if err != nil {
								logs.Trace("template parse file err:", err)
							}
						}
						break
					}
				}
			}
		}
	}
	return
}

type templateFSFunc func() http.FileSystem

func defaultFSFunc() http.FileSystem {
	return FileSystem{}
}

// SetTemplateFSFunc set default filesystem function
func SetTemplateFSFunc(fnt templateFSFunc) {
	radicalTemplateFS = fnt
}

// SetViewsPath sets view directory path in radiant application.
func SetViewsPath(path string) *HttpServer {
	BConfig.WebConfig.ViewsPath = path
	return RadicalApp
}

// SetStaticPath sets static directory path and proper url pattern in radiant application.
// if radiant.SetStaticPath("static","public"), visit /static/* to load static file in folder "public".
func SetStaticPath(url string, path string) *HttpServer {
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}
	if url != "/" {
		url = strings.TrimRight(url, "/")
	}
	BConfig.WebConfig.StaticDir[url] = path
	return RadicalApp
}

// DelStaticPath removes the static folder setting in this url pattern in radiant application.
func DelStaticPath(url string) *HttpServer {
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}
	if url != "/" {
		url = strings.TrimRight(url, "/")
	}
	delete(BConfig.WebConfig.StaticDir, url)
	return RadicalApp
}

// AddTemplateEngine add a new templatePreProcessor which support extension
func AddTemplateEngine(extension string, fn templatePreProcessor) *HttpServer {
	AddTemplateExt(extension)
	radicalTemplateEngines[extension] = fn
	return RadicalApp
}
