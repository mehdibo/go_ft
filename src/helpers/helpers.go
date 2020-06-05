/*
Copyright © 2020 Mehdi Bounya <mehdi.bounya@gmail.com>

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