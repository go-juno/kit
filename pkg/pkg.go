package pkg

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"golang.org/x/xerrors"
)

type em struct{}

func GetRoot() (root string, err error) {
	pkgList := strings.Split(reflect.TypeOf(em{}).PkgPath(), "/")
	pkgName := pkgList[0]
	if strings.Contains(pkgName, "/") {
		ditList := strings.Split(pkgName, "/")
		pkgName = ditList[len(ditList)-1]
	}

	root, err = os.Getwd()
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	splitList := strings.Split(root, pkgName)

	if len(splitList) > 1 {
		root = fmt.Sprintf("%s%s", splitList[0], pkgName)
	}
	return
}
