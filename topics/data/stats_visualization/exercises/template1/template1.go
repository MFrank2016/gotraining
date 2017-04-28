// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./template1

// Sample program to generate a box plot of diabetes bmi values.
package main

import (
	"bytes"
	"log"

	"github.com/kniren/gota/dataframe"
	"github.com/pachyderm/pachyderm/src/client"
)

func main() {

	// Connect to Pachyderm on our localhost.  By default
	// Pachyderm will be exposed on port 30650.
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Get the diabetes dataset from Pachyderm's data
	// versioning at the latest commit.
	var b bytes.Buffer
	if err := c.GetFile("diabetes", "master", "diabetes.csv", 0, 0, &b); err != nil {
		log.Fatal()
	}

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	diabetesDF := dataframe.ReadCSV(bytes.NewReader(b.Bytes()))

	// Create the plot and set its title and axis label.

	// Create the box for our data.

	// Create a plotter.Values value and fill it with the
	// values from the respective column of the dataframe.

	// Add the data to the plot.

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1, etc.

	// Save the plot.
}
