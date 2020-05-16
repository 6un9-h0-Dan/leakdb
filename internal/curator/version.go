package curator

/*
	---------------------------------------------------------------------
	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
	----------------------------------------------------------------------
*/

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	// Version - The semantic version of the program
	Version string

	// CompiledAt - Unix timestamp of compiled binary
	CompiledAt string

	// GitCommit - Most recent commit from which the binary was compiled
	GitCommit string

	// GitDirty - Set if the compiled from uncommitted code
	GitDirty string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	Long:  `Print the version number of leakdb-curator and exit`,
	Run: func(cmd *cobra.Command, args []string) {
		details, _ := cmd.Flags().GetBool(detailsFlagStr)
		if details {
			timeStamp, _ := strconv.Atoi(CompiledAt)
			compiledAt := time.Unix(int64(timeStamp), 0)
			fmt.Printf("LeakDB Curator v%s - Compiled %s - %s\n", Version, compiledAt, GitCommit)
		} else {
			fmt.Printf("LeakDB Curator v%s\n", Version)
		}
	},
}
