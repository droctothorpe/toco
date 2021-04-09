package cmd

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var disableSidebar bool
var disableHome bool

// genCmd represents the gen command.
var genCmd = &cobra.Command{
	Use:    "gen",
	Short:  "Generate a table of contents and inject it into your wiki's homepage and sidebar",
	Long:   "Generate a table of contents and inject it into your wiki's homepage and sidebar.",
	PreRun: toggleDebug,
	RunE: func(cmd *cobra.Command, args []string) error {
		var targetFiles []string
		if !disableHome {
			targetFiles = append(targetFiles, "Home.md")
		}
		if !disableSidebar {
			targetFiles = append(targetFiles, "_Sidebar.md")
		}

		if err := initializeFilesIterator(targetFiles); err != nil {
			return err
		}

		files, err := globFiles("md")
		if err != nil {
			return err
		}

		d := filesToMap(files)

		log.Info("Generating table of contents.")
		toc := generateTOC(d)
		log.Debug(toc)

		log.Infof("Injecting table of contents into %s.", targetFiles)
		if err := updateFilesIterator(targetFiles, toc); err != nil {
			return err
		}

		log.Info("Injection complete. Run 'toco push' to push your changes.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().BoolVar(&disableSidebar, "disable-sidebar", false, "Pass this flag to disable _Sidebar.md TOC injection")
	genCmd.Flags().BoolVar(&disableHome, "disable-home", false, "Pass this flag to disable Home.md TOC injection")
}

func globFiles(extension string) ([]string, error) {
	var files []string

	exclusionList := []string{"Home.md", "_Sidebar.md", "README.md"}

	matches, err := filepath.Glob("./*." + extension)
	if err != nil {
		return nil, err
	}

out:
	for _, match := range matches {
		for _, exclusion := range exclusionList {
			if match == exclusion {
				continue out
			}
		}
		files = append(files, match)
	}

	return files, nil
}

func generateTOC(d map[string][][]string) string {
	var keys []string
	for key := range d {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	log.Debug("sorted keys: ", keys)

	var toc string

	for _, key := range keys {
		toc += fmt.Sprintf("**%s**  \n", key)
		for _, value := range d[key] {
			toc += fmt.Sprintf("• [%s](%s)  \n", value[0], value[1])
		}
	}

	return toc
}

func writeOutput(output, file string) error {
	err := ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}

func filesToMap(files []string) map[string][][]string {
	d := make(map[string][][]string)
	separator := ":"

	for _, file := range files {
		log.Debug("file: ", file)
		split := strings.Split(file, separator)
		re := regexp.MustCompile(`-|_|\.`)
		category := split[0]
		category = re.ReplaceAllString(category, " ")
		log.Debug("category: ", category)

		title := split[1]
		title = title[0 : len(title)-3]
		title = re.ReplaceAllString(title, " ")
		log.Debug("title: ", title)

		path := url.QueryEscape(file)
		path = "./" + path[0:len(path)-3]
		log.Debug("path: ", path)

		d[category] = append(d[category], []string{title, path})
	}

	return d
}

func updateFile(file, toc string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		// TODO: Handle file missing toc tags.
		return err
	}

	input := string(b)

	re := regexp.MustCompile("starttoc-->((.|\n)*)<!--endtoc")
	toc = "starttoc-->\n" + toc + "<!--endtoc"
	output := re.ReplaceAllString(input, toc)

	if err := writeOutput(output, file); err != nil {
		return err
	}

	return nil
}

func initializeFile(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		tocTags := `
<!--tocstart-->
<!--tocend-->
		`
		if err := ioutil.WriteFile(filename, []byte(tocTags), 0644); err != nil {
			return err
		}
	}

	return nil
}

func initializeFilesIterator(filenames []string) error {
	for _, filename := range filenames {
		if err := initializeFile(filename); err != nil {
			return err
		}
	}

	return nil
}

func updateFilesIterator(filenames []string, toc string) error {
	for _, filename := range filenames {
		if err := updateFile(filename, toc); err != nil {
			return err
		}
	}

	return nil
}
