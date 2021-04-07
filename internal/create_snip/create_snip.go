package create_snip

import (
	"fmt"
	"io"
	"path"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func CreateSnip(w io.Writer, githubURL string, modDir string, dir string, title string, desc string) error {
	module, diags := tfconfig.LoadModule(path.Join(modDir, dir))
	if fmt.Sprintf("%v", diags) != "no problems" {
		return fmt.Errorf("%v", diags)
	}
	fmt.Fprintf(w, "snippet %s \"%s\"\n", title, desc)
	fmt.Fprintf(w, "module \"%s\"{\n", "{$1}")
	fmt.Fprintf(w, "\tsource = \"git::%s//%s\"\n", githubURL, dir)
	for key, rp := range module.RequiredProviders {
		fmt.Fprintf(w, "\t# required_providers = %s from %s \n", key, rp.Source)
	}

	for _, conf := range module.Variables {
		if conf.Required {
			fmt.Fprintf(w, "\t%s = \"%s\"\n", conf.Name, conf.Type)
		}
	}
	fmt.Fprintf(w, "\t# optional\n")
	for _, conf := range module.Variables {
		if !conf.Required {
			fmt.Fprintf(w, "\t# %s = \"%s\"\n", conf.Name, conf.Default)
		}
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "endsnippet")
	return nil
}
