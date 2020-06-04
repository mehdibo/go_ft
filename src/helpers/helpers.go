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