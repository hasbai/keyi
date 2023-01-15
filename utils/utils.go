package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type CanPreprocess interface {
	Preprocess(c *fiber.Ctx) error
}

type numbers interface {
	int | uint | int8 | uint8 |
	int16 | uint16 | int32 | uint32 |
	int64 | uint64 | float32 | float64
}

func Min[T numbers](x T, y T) T {
	if x > y {
		return y
	} else {
		return x
	}
}

// GetAbsPath get absolute path
func GetAbsPath() string {
	dir := getAbsPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getAbsPathByCaller()
	}
	return dir
}

// getAbsPathByExecutable 获取当前执行文件绝对路径
func getAbsPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// getAbsPathByCaller 获取当前执行文件绝对路径（go run）
func getAbsPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// ToMap struct to map, skips zero value and panics if in is not a struct
func ToMap(in any) map[string]any {
	out := make(map[string]any)
	const tagName = "json"

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic(fmt.Sprintf("ToMap only accepts struct; got %T", in))
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		ti := t.Field(i)
		vi := v.Field(i)

		if vi.IsZero() {
			continue
		}

		key := ti.Tag.Get(tagName)
		if key == "-" || key == "" {
			continue
		}
		// for a ptr, get its value
		if ti.Type.Kind() == reflect.Ptr {
			if vi.IsNil() {
				continue
			}
			out[key] = vi.Elem().Interface()
		} else {
			out[key] = vi.Interface()
		}
	}

	return out
}
