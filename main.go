package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/davyxu/tabtoy/compiler"
	"github.com/davyxu/tabtoy/gen"
	"github.com/davyxu/tabtoy/gen/bindata"
	"github.com/davyxu/tabtoy/gen/cssrc"
	"github.com/davyxu/tabtoy/gen/gosrc"
	"github.com/davyxu/tabtoy/gen/javasrc"
	"github.com/davyxu/tabtoy/gen/jsondata"
	"github.com/davyxu/tabtoy/gen/jsontype"
	"github.com/davyxu/tabtoy/gen/luasrc"
	"github.com/davyxu/tabtoy/gen/pbdata"
	"github.com/davyxu/tabtoy/gen/pbsrc"
	helper2 "github.com/davyxu/tabtoy/helper"
	model2 "github.com/davyxu/tabtoy/model"
	"github.com/davyxu/tabtoy/report"

	"github.com/davyxu/tabtoy/build"
	"github.com/pkg/profile"
)

// 生成器入口定义
type V3GenEntry struct {
	name          string
	genSingleFile gen.GenSingleFile
	genCustom     gen.GenCustom
	param         *string
}

// 命令行参数定义
type cmdFlags struct {
	// 基础参数
	enableProfile *bool
	version       *bool
	mode          *string
	para          *bool
	cacheDir      *string
	useCache      *bool
	modifyList    *string
	language      *string

	// 输出参数
	packageName       *string
	combineStructName *string
	protoOut          *string
	pbBinaryOut       *string
	pbtOut            *string
	luaOut            *string
	jsonOut           *string
	jsonTypeOut       *string
	csharpOut         *string
	goOut             *string
	binaryOut         *string
	typeOut           *string
	cppOut            *string
	javaOut           *string
	jsonDir           *string
	luaDir            *string
	binaryDir         *string
	pbBinaryDir       *string
	indexFile         *string
	tagAction         *string
}

// 初始化命令行参数
func initFlags() *cmdFlags {
	f := &cmdFlags{}

	// 基础参数
	f.enableProfile = flag.Bool("profile", false, "Enable profiling")
	f.version = flag.Bool("version", false, "Show version")
	f.mode = flag.String("mode", "", "v3")
	f.para = flag.Bool("para", false, "parallel export by your cpu count")
	f.cacheDir = flag.String("cachedir", "./.tabtoycache", "cache file output dir")
	f.useCache = flag.Bool("usecache", false, "use cache file enhanced exporting speed")
	f.modifyList = flag.String("modlistfile", "", "output list to file, include not using cache input file list")
	f.language = flag.String("lan", "en_us", "set output language")

	// 输出参数
	f.packageName = flag.String("package", "", "override the package name in table @Types")
	f.combineStructName = flag.String("combinename", "Table", "combine struct name")
	f.protoOut = flag.String("proto_out", "", "output protobuf define (*.proto)")
	f.pbBinaryOut = flag.String("pbbin_out", "", "output protobuf binary (*.pbb)")
	f.pbtOut = flag.String("pbt_out", "", "output proto text format (*.pbt)")
	f.luaOut = flag.String("lua_out", "", "output lua code (*.lua)")
	f.jsonOut = flag.String("json_out", "", "output json format (*.json)")
	f.jsonTypeOut = flag.String("jsontype_out", "", "output json type (*.json)")
	f.csharpOut = flag.String("csharp_out", "", "output c# class and deserialize code (*.cs)")
	f.goOut = flag.String("go_out", "", "output golang code (*.go)")
	f.binaryOut = flag.String("binary_out", "", "output binary format(*.bin)")
	f.typeOut = flag.String("type_out", "", "output table types(*.json)")
	f.cppOut = flag.String("cpp_out", "", "output c++ format (*.cpp)")
	f.javaOut = flag.String("java_out", "", "output java code (*.java)")
	f.jsonDir = flag.String("json_dir", "", "output json format (*.json) to dir")
	f.luaDir = flag.String("lua_dir", "", "output lua format (*.lua) to dir")
	f.binaryDir = flag.String("binary_dir", "", "output binary format (*.bin) to dir")
	f.pbBinaryDir = flag.String("pbbin_dir", "", "output binary format (*.pbb) to dir")
	f.indexFile = flag.String("index", "", "input multi-files configs")
	f.tagAction = flag.String("tag_action", "", "do action by tag selected target")
	return f
}

// 生成器列表
func initGenList(flags *cmdFlags) []V3GenEntry {
	return []V3GenEntry{
		{name: "gosrc", genSingleFile: gosrc.Generate, param: flags.goOut},
		{name: "jsondata", genSingleFile: jsondata.Generate, param: flags.jsonOut},
		{name: "jsontype", genSingleFile: jsontype.Generate, param: flags.jsonTypeOut},
		{name: "luasrc", genSingleFile: luasrc.Generate, param: flags.luaOut},
		{name: "cssrc", genSingleFile: cssrc.Generate, param: flags.csharpOut},
		{name: "bindata", genSingleFile: bindata.Generate, param: flags.binaryOut},
		{name: "javasrc", genSingleFile: javasrc.Generate, param: flags.javaOut},
		{name: "pbsrc", genSingleFile: pbsrc.Generate, param: flags.protoOut},
		{name: "pbdata", genSingleFile: pbdata.Generate, param: flags.pbBinaryOut},
		{name: "jsondir", genCustom: jsondata.Output, param: flags.jsonDir},
		{name: "luadir", genCustom: luasrc.Output, param: flags.luaDir},
		{name: "binarydir", genCustom: bindata.Output, param: flags.binaryDir},
		{name: "pbdatadir", genCustom: pbdata.Output, param: flags.pbBinaryDir},
	}
}

// 生成单个文件
func genFile(ctx context.Context, globals *model2.Globals, entry V3GenEntry) error {
	filename := *entry.param

	if entry.genSingleFile != nil {
		data, err := entry.genSingleFile(globals)
		if err != nil {
			return err
		}

		report.Log.Infof("  [%s] %s", entry.name, filename)
		return helper2.WriteFile(filename, data)
	}

	if entry.genCustom != nil {
		if err := entry.genCustom(globals, filename); err != nil {
			return err
		}
		report.Log.Infof("  [%s] %s", entry.name, filename)
	}
	return nil
}

// GenFileByList 并发生成所有文件
func GenFileByList(ctx context.Context, globals *model2.Globals, genList []V3GenEntry) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(genList))

	for _, entry := range genList {
		if *entry.param == "" {
			continue
		}

		wg.Add(1)
		go func(entry V3GenEntry) {
			defer wg.Done()
			if err := genFile(ctx, globals, entry); err != nil {
				errChan <- err
			}
		}(entry)
	}

	// 等待所有生成任务完成
	wg.Wait()
	close(errChan)

	// 检查是否有错误
	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flags := initFlags()
	flag.Parse()

	// 设置性能分析
	if *flags.enableProfile {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	}

	// 创建带取消的上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 处理中断信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cancel()
	}()

	// 初始化全局配置
	globals := model2.NewGlobals()
	globals.Version = build.Version
	globals.ParaLoading = *flags.para
	if *flags.useCache {
		globals.CacheDir = *flags.cacheDir
		if err := os.MkdirAll(globals.CacheDir, 0755); err != nil {
			report.Log.Errorln(err)
			os.Exit(1)
		}
	}
	globals.IndexFile = *flags.indexFile
	globals.PackageName = *flags.packageName
	globals.CombineStructName = *flags.combineStructName
	globals.GenBinary = *flags.binaryOut != "" || *flags.binaryDir != ""

	// 设置文件加载器
	globals.IndexGetter = helper2.NewFileLoader(true, globals.CacheDir)

	// 解析标签动作
	if *flags.tagAction != "" {
		tagActions, err := model2.ParseTagAction(*flags.tagAction)
		if err != nil {
			report.Log.Errorln(err)
			os.Exit(1)
		}
		globals.TagActions = tagActions
	}

	// 编译
	if err := compiler.Compile(globals); err != nil {
		report.Log.Errorln(err)
		os.Exit(1)
	}

	// 生成文件
	report.Log.Debugln("Generate files...")
	genList := initGenList(flags)
	if err := GenFileByList(ctx, globals, genList); err != nil {
		report.Log.Errorln(err)
		os.Exit(1)
	}
}
