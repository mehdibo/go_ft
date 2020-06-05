/*
Copyright © 2020 Mehdi Bounya <mehdi.bounya@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package helpers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

/**
 * This function will return the body as a []byte
 * If any errors are encountered, it will print a message and exit with status 1
 */
func GetRespBody(resp *http.Response) []byte {
	defer resp.Body.Close()

	respBody, ioErr := ioutil.ReadAll(resp.Body)

	if ioErr != nil && ioErr != io.EOF {
		fmt.Fprintf(os.Stderr, "There was an error reading the response body")
		fmt.Fprintln(os.Stderr, ioErr)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "The API returned the following status code: %d\n", resp.StatusCode)
		fmt.Fprintf(os.Stderr, "And the following body: \n%s\n", string(respBody))
		os.Exit(1)
	}
	return respBody
}

func WriteToFile(outputFile string, content string) {
	resultsFile, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Couldn't create file: \n%s\n", err)
		os.Exit(1)
	}

	_, err = resultsFile.WriteString(content)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Couldn't write data to file: \n%s\n", err)
		os.Exit(1)
	}
}