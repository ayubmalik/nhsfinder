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
	"io/ioutil"
	"os"
	"path"
	"strings"

	finder "github.com/ayubmalik/pharmacyfinder"
	"github.com/mholt/archiver"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO: move to config file/viper
const (
	pharmacyCSV  = "http://media.nhschoices.nhs.uk/data/foi/Pharmacy.csv"
	gpCSV        = "http://media.nhschoices.nhs.uk/data/foi/GP.csv"
	postcodesZip = "https://www.getthedata.com/downloads/open_postcode_geo.csv.zip"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download source data required by this application",
	Long: `Download source data required by this application.
The NHS data is downloaded from NHS Choices. (TODO: use ODS datasets).
The UK postcode data is downloaded from http://getthedata.com.
The data is also sanitised and simplified where required.
`,
}

var downloadPharmacyCmd = &cobra.Command{
	Use:   "pharmacy",
	Short: "download pharmacy data",
	Run: func(cmd *cobra.Command, args []string) {
		dataDir := viper.GetString("data")
		downloadODS(&finder.HTTPDownloader{}, pharmacyCSV, path.Join(dataDir, "pharmacy.csv"))
	},
}

var downloadGPCmd = &cobra.Command{
	Use:   "gp",
	Short: "download GP data",
	Run: func(cmd *cobra.Command, args []string) {
		dataDir := viper.GetString("data")
		downloadODS(&finder.HTTPDownloader{}, gpCSV, path.Join(dataDir, "gp.csv"))
	},
}

var downloadPostcodeCmd = &cobra.Command{
	Use:   "postcode",
	Short: "download postcode data",
	Run: func(cmd *cobra.Command, args []string) {
		dataDir := viper.GetString("data")
		downloadPostcodes(&finder.HTTPDownloader{}, path.Join(dataDir, "postcode.csv"))
	},
}

func init() {
	downloadCmd.AddCommand(downloadPharmacyCmd)
	downloadCmd.AddCommand(downloadGPCmd)
	downloadCmd.AddCommand(downloadPostcodeCmd)
	rootCmd.AddCommand(downloadCmd)
}

func downloadODS(d finder.Downloader, srcURL, outputFile string) {
	tmpDir, err := ioutil.TempDir("", "finder-")
	if err != nil {
		panic(err)
	}
	defer func() { os.RemoveAll(tmpDir) }()

	base := path.Base(srcURL)
	tmpFile := path.Join(tmpDir, base)

	d.Download(srcURL, tmpFile)
	finder.SimplifyODS(tmpFile, outputFile)
}

func downloadPostcodes(d finder.Downloader, outputFile string) {
	tmpDir, err := ioutil.TempDir("", "finder-")
	if err != nil {
		panic(err)
	}
	defer func() { os.RemoveAll(tmpDir) }()

	base := path.Base(postcodesZip)
	tmpFile := path.Join(tmpDir, base)

	d.Download(postcodesZip, tmpFile)
	archiver.Unarchive(tmpFile, tmpDir)
	csvFile := path.Join(tmpDir, strings.Replace(base, ".zip", "", 1))

	cf, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}
	defer cf.Close()

	of, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer of.Close()

	finder.SimplifyPostcodes(cf, of)
}
