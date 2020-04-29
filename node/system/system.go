package system

import(
	"os"
	"fmt"
	"time"
	"github.com/SaKu2110/chain_dev/chain"
)

var(
	logfilePath string
)

func intit() {
	logfilePath = os.Getenv("LOGFILE_PATH")
	if logfilePath == "" {
		logfilePath = "./"
	}
}

func EditLogFile(data chain.Block, hash [32]byte) (error) {
	file, err := os.OpenFile(logfilePath + "chain.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Fprintln(file, time.Unix(time.Now().Unix(), 0))
	fmt.Fprintln(file, data)
	fmt.Fprintln(file, hash)
	return nil
}