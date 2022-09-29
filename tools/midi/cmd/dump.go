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
	"unsafe"

	"github.com/spf13/cobra"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		files := args[0]
		if files == "" {
			return
		}
		smf := C.scc_smf_parse_file((C.CString)(args[0]))
		C.scc_smf_dump(smf)
		C.free(unsafe.Pointer(smf))
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
