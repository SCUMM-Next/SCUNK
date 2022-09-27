/*
Copyright © 2022 Mark "Happy-Ferret" Bauermeister <warfan2007@gmail.com>

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
package cmd

// #cgo CFLAGS: -I../../..
// #cgo LDFLAGS: -L../../..
// #include "../../../include/pendian.h"
// #if __BYTE_ORDER == __BIG_ENDIAN
// #define IS_BIG_ENDIAN
// #elif __BYTE_ORDER == __LITTLE_ENDIAN
// #define IS_LITTLE_ENDIAN
// #endif
// #include <stdlib.h>
// #include <stdio.h>
// #include <string.h>
// #include <inttypes.h>
// #include <errno.h>
// #include <sys/types.h>
// #include <sys/stat.h>
// #include <fcntl.h>
// #include "../../../scc_util.h"
// #include "../../../scc_fd.h"
// #include "../../../scc_param.h"
// #include "../../../scc_smf.h"
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/spf13/cobra"
)

type CParam struct {
	Param C.scc_param_t
}

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var files *C.scc_cl_arg_t
		// var smf *C.scc_ksmf_t

		params := CParam{Param: C.scc_param_t{(*C.char)(unsafe.Pointer(&args[0])), 0, 0, 0, nil}}
		cParams := *(*C.scc_param_t)(unsafe.Pointer(&params))
		// params := C.scc_param_t((*C.char)(unsafe.Pointer(&args[0])))
		argc := C.int(len(args) - 1)
		argv := (*C.char)(unsafe.Pointer(&args[1]))
		files = (*C.scc_cl_arg_t)(unsafe.Pointer(C.scc_param_parse_argv(&cParams, argc, &argv)))
		// smf := C.scc_smf_parse_file(C.scc_smf_t(("test.midi")))
		if files == nil {
			fmt.Println(cParams)
			fmt.Println(argc)
			fmt.Println(argv)
			// fmt.Println(smf)
		}
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dumpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
