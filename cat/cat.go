package cat
/* Concatenate files in "stdout". */
import(
	"os"
	"io"
	"flag"
	"fmt"
)

func cat(p string) error {
	f, e := os.Open(p)
	if e != nil {
		return e
	}
	defer f.Close()
	fcat(f)
	return nil
}

func fcat(f *os.File) error {
	b := make([]byte, 512)
	for {
		n, e := f.Read(b)
		if n>0 {
			fmt.Print(string(b[:n]))
		}
		if e == io.EOF {
			break
		}else if e != nil {
			return e
		}
	}
	return nil
}

func Run(args []string) int {	
	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagSet.Parse(args[1:])
	arg0 := args[0]
	args = args[1:]
	status := 0
	if len(args)>0 {
		for _, p := range args {
			e := cat(p)
			if e != nil {
				fmt.Fprintf(os.Stderr, "%s: %s.\n", arg0, e)
			}
		}
	} else {
		fcat(os.Stdin)
	}
	return status
}