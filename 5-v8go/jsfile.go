/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     chromedp-rod-test
 * @Date:        2023-05-31 21:43
 * @Description:
 */

package main

import (
	_ "embed"
	"fmt"
)

//go:embed js/cheerio.min.js
var FILE_CHEERIO string

//go:embed js/jiexi.js
var FILE_JIEXI string

var TPL_CHEERIO_JS = func() string {
	return fmt.Sprintf("%s\n%s", FILE_CHEERIO, FILE_JIEXI)
}
