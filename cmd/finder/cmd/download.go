/*
Copyright © 2020 Ayub Malik <ayub.malik@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	finder "github.com/ayubmalik/pharmacyfinder"
	"github.com/spf13/cobra"
)

// TODO: move to config file/viper
const (
	dataDir     = "data"
	pharmacyCSV = "http://media.nhschoices.nhs.uk/data/foi/Pharmacy.csv"
	gpCSV       = "http://media.nhschoices.nhs.uk/data/foi/GP.csv"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download pharmacy|gp",
	Short: "Download NHS pharmacy or GP data",
	Long: `Download NHS pharmacy or GP data.
	The CSV data is downloaded from the NHS Choices dataset for now. (TODO: use ODS datasets)
The data is also sanitised and simplified where required.
Valid options are 'pharmacy' or 'gp'.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires at least one of: %v", cmd.ValidArgs)
		}
		return cobra.OnlyValidArgs(cmd, args)
	},
	ValidArgs: []string{"pharmacy", "gp"},
	Run: func(cmd *cobra.Command, args []string) {
		org := args[0]
		switch org {
		case "gps":
			downloadGP()
		default:
			downloadPharmacy()
		}
	},
}

func downloadGP() {
}

func downloadPharmacy() {
	tmpDir, err := ioutil.TempDir("", "finder-")
	if err != nil {
		panic(err)
	}
	defer func() { os.RemoveAll(tmpDir) }()

	base := path.Base(pharmacyCSV)
	destFile := path.Join(tmpDir, base)

	downloader := finder.HTTPDownloader{}
	downloader.Download(pharmacyCSV, destFile)

	finder.SimplifyPharmacies(destFile, path.Join(dataDir, "pharmacies.csv"))
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

}
