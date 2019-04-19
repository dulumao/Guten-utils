package env

import (
	"os"
)

func All() []string {
	return os.Environ()
}

// 获取环境变量，并可以指定当环境变量不存在时的默认值
func Get(k string, def...string) string {
	v, ok := os.LookupEnv(k)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

func Set(k, v string) error {
	return os.Setenv(k, v)
}

func Remove(k string) error {
	return os.Unsetenv(k)
}
