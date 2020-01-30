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

	finder "github.com/ayubmalik/pharmacyfinder"
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

func init() {

	searchCmd.AddCommand(searchPharmacyCmd)
	rootCmd.AddCommand(searchCmd)
}

func searchPharmacy(dataDir string, postcode string) {
	postcodeFile, err := os.Open(path.Join(dataDir, "postcode.csv"))
	if err != nil {
		exitError("could not open postcode file: %v\n", err)
	}
	defer postcodeFile.Close()

	latLngs, err := finder.LoadLatLngs(postcodeFile)
	if err != nil {
		exitError("could not load postcode data: %v\n", err)
	}
	fmt.Printf("Loaded %d postcodes with latlng\n", len(latLngs))

	pharmacyFile, err := os.Open(path.Join(dataDir, "pharmacy.csv"))
	if err != nil {
		exitError("could not open pharmacy file: %v", err)
	}
	defer pharmacyFile.Close()

	pharmacies, err := finder.LoadPharmacies(pharmacyFile)
	if err != nil {
		exitError("could not load pharmacy data: %v\n", err)
	}
	fmt.Printf("Loaded %d pharmacies with latlng\n", len(pharmacies))

	// create in mem finder
	finder := finder.InMemFinder{LatLngs: latLngs, Pharmacies: pharmacies}
	results := finder.ByPostcode(postcode)
	display(results)
}

func searchGP(postcode string) {

}

func display(results []finder.FindResult) {
	for i, r := range results {
		fmt.Printf("%2d %7.2f %-30s %-30s %s\n", i, r.Distance, r.Pharmacy.Name, r.Pharmacy.Address.Line1, r.Pharmacy.Address.Postcode)
	}
}

func exitError(msg string, err error) {
	fmt.Fprintln(os.Stderr, msg, err)
	os.Exit(1)
}
