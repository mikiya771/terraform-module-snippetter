package module_repo_analyzer

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

type ModuleInfo struct {
	OrgName    string
	RepoName   string
	ModuleName string
	ModulePath string
}

func (m *ModuleInfo) GetDescription() string {
	return fmt.Sprintf("module %s in %s/%s", m.ModuleName, m.OrgName, m.RepoName)
}

func (m *ModuleInfo) GetTitle() string {
	return fmt.Sprintf("m_%s_%s_%s", m.OrgName, m.RepoName, m.ModuleName)
}

func ModuleRepoAnalyzer(cacheDir string, orgName string, repoName string) []*ModuleInfo {
	moduleInfos := []*ModuleInfo{}
	dirs := searchDirs(path.Join(cacheDir, orgName, repoName), "")
	for _, dir := range dirs {
		if !isIgnored(path.Join(cacheDir, orgName, repoName, dir)) {
			pathSplit := strings.Split(dir, "/")
			m := ModuleInfo{
				OrgName:    orgName,
				RepoName:   repoName,
				ModuleName: pathSplit[len(pathSplit)-1],
				ModulePath: dir,
			}
			moduleInfos = append(moduleInfos, &m)
		}
	}
	return moduleInfos
}

func isIgnored(dirPath string) bool {
	// if dir has ignore pattern file (*_test.go), we ignore the
	patterns := []string{"_test.go"}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			for _, pattern := range patterns {
				if strings.Contains(file.Name(), pattern) {
					return true
				}
			}
		}
	}
	return false
}

func searchDirs(absRootPath, relativePath string) []string {
	files, err := ioutil.ReadDir(path.Join(absRootPath, relativePath))
	dirPatterns := []string{".git", "example"}
	ret := []string{}
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			ignored := false
			for _, pattern := range dirPatterns {
				if strings.Contains(file.Name(), pattern) {
					ignored = true
				}
			}
			if !ignored {
				nextPath := path.Join(relativePath, file.Name())
				ret = append(ret, nextPath)
				ret = append(ret, searchDirs(absRootPath, nextPath)...)
			}
		}
	}
	return ret
}
