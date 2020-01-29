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
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

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

	fmt.Println("Loading data from ", dataDir)

	latLngs := finder.LoadLatLngs(path.Join(dataDir, "postcode.csv"))
	fmt.Printf("Loaded %d postcodes\n", len(latLngs))

	pharmacies := finder.LoadPharmacies("data/pharmacies.csv")
	fmt.Printf("Loaded %d pharmacies with lat/lng\n", len(pharmacies))

	pcode1 := "BD18 2DS"
	pcode2 := "M4 4BF"
	from := latLngs[pcode1]
	to := latLngs[pcode2]
	dist1 := finder.Distance(from, to)
	fmt.Printf("Distance from '%s' to '%s': %fm\n", pcode1, pcode2, dist1)

	find := finder.InMemFinder{LatLngs: latLngs, Pharmacies: pharmacies}
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter postcode in format M4 4BF: ")
		pcode, _ := reader.ReadString('\n')
		pcode = strings.Replace(pcode, "\n", "", -1)
		pcode = strings.ToUpper(pcode)
		if len(pcode) < 2 {
			fmt.Println("Goodbye!")
			return
		}

		results := find.ByPostcode(pcode)
		display(results)
	}
}

func searchGP(postcode string) {

}

func display(results []finder.FindResult) {
	for i, r := range results {
		fmt.Printf("%2d %7.2f %-30s %-30s %s\n", i, r.Distance, r.Pharmacy.Name, r.Pharmacy.Address.Line1, r.Pharmacy.Address.Postcode)
	}
}
