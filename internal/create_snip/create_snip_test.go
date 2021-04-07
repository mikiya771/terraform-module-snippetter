package create_snip

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func ExampleCreateSnip() {
	var b bytes.Buffer
	dir := "test"
	githubURL := "https://github.com/org/repo"
	err := CreateSnip(&b, githubURL, dir, "m_local_test", "module_local_test")
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("%s", b.String())

	// Unordered Output:
	// snippet m_local_test "module_local_test"
	// module "{$1}"{
	// 	source = "git::https://github.com/org/repo//test"
	// 	nyan = "string"
	// 	nyan1 = ""
	// 	nyan2 = ""
	// 	nyan3 = ""
	// 	# optional
	// 	# nyan4 = "hoge"
	// 	# nyan5 = "map[iruka:statoko tanuki:tagarogu]"
	// }
	// endsnippet

}

func useIoutilReadFile(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
