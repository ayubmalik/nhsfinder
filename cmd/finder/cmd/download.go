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
	"strings"

	finder "github.com/ayubmalik/pharmacyfinder"
	"github.com/mholt/archiver"
	"github.com/spf13/cobra"
)

// TODO: move to config file/viper
const (
	dataDir      = "data"
	pharmacyCSV  = "http://media.nhschoices.nhs.uk/data/foi/Pharmacy.csv"
	gpCSV        = "http://media.nhschoices.nhs.uk/data/foi/GP.csv"
	postcodesZip = "https://www.getthedata.com/downloads/open_postcode_geo.csv.zip"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download pharmacy|gp|postcode",
	Short: "Download NHS pharmacy or GP data or UK postcode data",
	Long: `Download NHS pharmacy or GP data or UK postcode data.
The NHS data is downloaded from the NHS Choices dataset for now. (TODO: use ODS datasets).
The UK postcode data is downloaded from getthedata.com.
The data is also sanitised and simplified where required.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires at least one of: %v", cmd.ValidArgs)
		}
		return cobra.OnlyValidArgs(cmd, args)
	},
	ValidArgs: []string{"pharmacy", "gp", "postcode"},
	Run: func(cmd *cobra.Command, args []string) {
		org := args[0]
		switch org {
		case "pharmacy":
			downloadODS(&finder.HTTPDownloader{}, path.Join(dataDir, "pharmacies.csv"))
		case "gps":
			downloadODS(&finder.HTTPDownloader{}, path.Join(dataDir, "gps.csv"))
		case "postcode":
			downloadPostcodes(&finder.HTTPDownloader{}, path.Join(dataDir, "ukpostcodes.csv"))
		default:
			cmd.Usage()
		}
	},
}

func downloadODS(d finder.Downloader, outputFile string) {
	tmpDir, err := ioutil.TempDir("", "finder-")
	if err != nil {
		panic(err)
	}
	defer func() { os.RemoveAll(tmpDir) }()

	base := path.Base(gpCSV)
	tmpFile := path.Join(tmpDir, base)

	d.Download(gpCSV, tmpFile)
	finder.SimplifyODS(tmpFile, outputFile)
}

func downloadPostcodes(d finder.Downloader, outputFile string) {
	tmpDir, err := ioutil.TempDir("", "finder-")
	if err != nil {
		panic(err)
	}
	//defer func() { os.RemoveAll(tmpDir) }()

	base := path.Base(postcodesZip)
	tmpFile := path.Join(tmpDir, base)

	d.Download(postcodesZip, tmpFile)
	archiver.Unarchive(tmpFile, tmpDir)
	csvFile := path.Join(tmpDir, strings.Replace(base, ".zip", "", 1))
	fmt.Println(csvFile)

	cf, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}
	defer cf.Close()

	of, err := os.Open(outputFile)
	if err != nil {
		panic(err)
	}
	defer of.Close()

	fmt.Println(outputFile)
	finder.SimplifyPostcodes(cf, of)
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
