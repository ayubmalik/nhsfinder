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
	"os"
	"path"
	"time"

	finder "github.com/ayubmalik/nhsfinder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for pharmacy or GP by postcode",
	Long: `Search for pharmacy or GP by postcode. Only the nearest N are returned.
`,
}

var searchPharmacyCmd = &cobra.Command{
	Use:   "pharmacy [postcode]",
	Short: "search for pharmacy by postcode",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := viper.GetString("data")
		searchPharmacy(data, args[0])
	},
}

var searchGPCmd = &cobra.Command{
	Use:   "gp [postcode]",
	Short: "search for GP by postcode",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := viper.GetString("data")
		searchGP(data, args[0])
	},
}

func init() {
	searchCmd.AddCommand(searchPharmacyCmd)
	searchCmd.AddCommand(searchGPCmd)
	rootCmd.AddCommand(searchCmd)
}

func searchPharmacy(dataDir, postcode string) {
	postcodeFile, err := os.Open(path.Join(dataDir, "postcode.csv"))
	if err != nil {
		exitError("could not open postcode file: %v\n", err)
	}
	defer postcodeFile.Close()

	latLngs, err := finder.LoadLatLngs(postcodeFile)
	if err != nil {
		exitError("could not load postcode data: %v\n", err)
	}

	pharmacyFile, err := os.Open(path.Join(dataDir, "pharmacy.csv"))
	if err != nil {
		exitError("could not open pharmacy file: %v", err)
	}
	defer pharmacyFile.Close()

	pharmacies, err := finder.LoadPharmacies(pharmacyFile)
	if err != nil {
		exitError("could not load pharmacy data: %v\n", err)
	}

	// create in mem finder
	finder := finder.InMemFinder{LatLngs: latLngs, Pharmacies: pharmacies}
	start := time.Now()
	results := finder.FindPharmacies(postcode)
	end := time.Now().Sub(start)

	fmt.Printf("Calculated GP distances from %s in %s\n", postcode, end)
	for i, r := range results {
		fmt.Printf("%2d %7.2f %-40s %-30s %8s\n", i+1, r.Distance, r.Pharmacy.Name, r.Pharmacy.Address.Line1, r.Pharmacy.Address.Postcode)
	}
}

func searchGP(dataDir, postcode string) {
	postcodeFile, err := os.Open(path.Join(dataDir, "postcode.csv"))
	if err != nil {
		exitError("could not open postcode file: %v\n", err)
	}
	defer postcodeFile.Close()

	latLngs, err := finder.LoadLatLngs(postcodeFile)
	if err != nil {
		exitError("could not load postcode data: %v\n", err)
	}

	gpFile, err := os.Open(path.Join(dataDir, "gp.csv"))
	if err != nil {
		exitError("could not open GP file: %v", err)
	}
	defer gpFile.Close()

	gps, err := finder.LoadGPs(gpFile)
	if err != nil {
		exitError("could not load GP data: %v\n", err)
	}

	// create in mem finder
	finder := finder.InMemFinder{LatLngs: latLngs, GPs: gps}
	start := time.Now()
	results := finder.FindGPs(postcode)
	end := time.Now().Sub(start)

	fmt.Printf("Calculated GP distances from %s in %s\n", postcode, end)
	for i, r := range results {
		fmt.Printf("%2d %7.2f %-40s %-30s %8s\n", i+1, r.Distance, r.GP.Name, r.GP.Address.Line1, r.GP.Address.Postcode)
	}
}

func exitError(msg string, err error) {
	fmt.Fprintln(os.Stderr, msg, err)
	os.Exit(1)
}
