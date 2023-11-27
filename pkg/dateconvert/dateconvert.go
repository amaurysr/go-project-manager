package dateconvert 

import(
	"time" 
	"fmt"
)

func Date() string{
	dt := time.Now()
	formatdate := fmt.Sprintf(dt.Format("20060102150405"))
	return formatdate
}
