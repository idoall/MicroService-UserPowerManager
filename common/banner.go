package common

import (
	"fmt"
	"os"

	"github.com/agtorre/gocolorize"
)

const banner = `
 __  __   ____   _   _   _  __   _____ ___  ____  
|  \/  | / ___| | | | | | |/ /  |_   _/ _ \|  _ \ 
| |\/| | \___ \ | |_| | | ' /     | || | | | |_) |
| |  | |  ___) ||  _  | | . \   _ | || |_| |  __/ 
|_|  |_| \____/ |_| |_| |_|\_\ (_)|_| \___/|_|  

`

// OutBanner 输出 banner
func OutBanner() {
	fmt.Fprintf(os.Stdout, gocolorize.NewColor("green").Paint(banner))
}
