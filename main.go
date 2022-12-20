package main

import (
	"github.com/teris-io/cli"
	"github.com/daviddengcn/go-colortext"
	"os"
	
	// "os/exec"
	"fmt"
	// "strings"
	"io/ioutil"

	"context"
	// "math/rand"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/barchart"
)

func countFileInDir(path string,minusBy int)(countFileInDir_out int){
    files,_ := ioutil.ReadDir(path)
    myInt := len(files)-minusBy
    return myInt
}


func playBarChart(ctx context.Context, bc *barchart.BarChart, delay time.Duration) {
	max := 100

	

	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for {
		out_ready := countFileInDir("./ภายนอก/เตรียมแนบ",1)
		out_done := countFileInDir("./ภายนอก/เตรียมแนบ/แนบแล้ว",0)
		out_sum := out_ready + out_done
		in_ready := countFileInDir("./ภายใน/เตรียมแนบ",1)
		in_done :=  countFileInDir("./ภายใน/เตรียมแนบ/แนบแล้ว",0)
		in_sum := in_ready + in_done
		// in_out := out_sum + in_sum
		if in_sum > out_sum {
			max = in_sum
		}else{
			max = out_sum
		}
		
		select {
		case <-ticker.C:
			// var brBarColor []cell.Color
			var values []int
			for i := 0; i < 4; i++ {
				switch i {
				case 0:
					// if out_done == out_sum {
					// 	brBarColor = append(brBarColor,  cell.ColorGreen)
					// }else{
					// 	brBarColor = append(brBarColor,  cell.ColorRed)
					// }
				
					values = append(values,  out_done)
				case 1:
					values = append(values, out_sum )
				case 2:
					// if in_done == in_sum {
					// 	brBarColor = append(brBarColor,  cell.ColorGreen)
					// }else{
					// 	brBarColor = append(brBarColor,  cell.ColorRed)
					// }
					values = append(values, in_done )
				case 3:
					values = append(values, in_sum)
				}
				// values = append(values, int(rand.Int31n(max+1)))
			}

			if err := bc.Values(values, max); err != nil {
				panic(err)
			}

			// bc.BarColors(brBarColor)


		case <-ctx.Done():
			return
		}
	}
}

func showLog (comment string , data []byte) {
	fmt.Println("")
	ct.Foreground(ct.Green, false)
	fmt.Println(comment)
	ct.ResetColor()
	fmt.Printf("%s\n",data)
	fmt.Println("")
}

func showLogString (comment string , data string) {
	fmt.Println("")
	ct.Foreground(ct.Green, false)
	fmt.Println(comment)
	ct.ResetColor()
	fmt.Println(data)
	fmt.Println("")
}



func showErr (comment string , data error) {
	fmt.Println("")
	ct.Foreground(ct.Red, false)
	fmt.Println(comment)
	ct.ResetColor()
	fmt.Printf("%s\n",data)
	fmt.Println("")
}

func main() {

	// now := time.Now().Format("2006-01-02 15:04:05")
	// dir_name_with_nl_bytes , _ := exec.Command("basename",os.Getenv("PWD")).Output()
	// dir_name_with_nl_string := string(dir_name_with_nl_bytes)
	// dir_name := strings.Replace(dir_name_with_nl_string, "\n", "",-1)
	
	showPath := cli.NewCommand("showPath", "show current path ").
	WithShortcut("sp").
	WithAction(func(args []string, options map[string]string) int {
		PWD , _ := os.Getwd()
		showLogString("showPath",PWD)
		return 0
	})

	dirDate := cli.NewCommand("dirDate", "mkdir name=YYMMDD").
	WithShortcut("dd").
	WithAction(func(args []string, options map[string]string) int {
		YYMMDD := string(time.Now().Format("060102"))
		mkdirYYMMDD := os.Mkdir(YYMMDD,0666)
		if mkdirYYMMDD != nil {
			showErr("error :",mkdirYYMMDD)
		}else{
			showLogString("dirDate : ",YYMMDD)
		}
			// showLog("mkdir YYMMDD",mkdirYYMMDD)
		return 0
	})

	dirSarabanDocScan := cli.NewCommand("dirSarabanDocScan", "...").
	WithShortcut("dds").
	WithAction(func(args []string, options map[string]string) int {
		YYMMDD := string(time.Now().Format("06.01.02"))
		mkdirYYMMDD := os.Mkdir(YYMMDD,0666)
		if mkdirYYMMDD != nil {
			showErr("error :",mkdirYYMMDD)
		}else{
			showLogString("dirDate : ",YYMMDD)
		}
		os.Mkdir(YYMMDD+"/ภายนอก",0666)
		os.Mkdir(YYMMDD+"/ภายใน",0666)
		os.Mkdir(YYMMDD+"/ไม่เสนอ",0666)
		os.Mkdir(YYMMDD+"/ภายนอก/เตรียมแนบ",0666)
		os.Mkdir(YYMMDD+"/ภายใน/เตรียมแนบ",0666)
		os.Mkdir(YYMMDD+"/ไม่เสนอ/เตรียมแนบ",0666)
		os.Mkdir(YYMMDD+"/ภายนอก/เตรียมแนบ/แนบแล้ว",0666)
		os.Mkdir(YYMMDD+"/ภายใน/เตรียมแนบ/แนบแล้ว",0666)
		os.Mkdir(YYMMDD+"/ไม่เสนอ/เตรียมแนบ/แนบแล้ว",0666)
		
		return 0
	})

	dirWatch:= cli.NewCommand("dirWatch", "...").
	WithShortcut("dw").
	WithAction(func(args []string, options map[string]string) int {

			t, err := tcell.New()
			if err != nil {
				panic(err)
			}
			defer t.Close()


		
			ctx, cancel := context.WithCancel(context.Background())
			bc, err := barchart.New(
				
				barchart.BarColors([]cell.Color{
					cell.ColorYellow,
					cell.ColorYellow,
					cell.ColorYellow,
					cell.ColorYellow,
					// cell.ColorGreen,
				}),
				barchart.ValueColors([]cell.Color{
					cell.ColorRed,
					cell.ColorRed,
					cell.ColorRed,
					cell.ColorRed,
				}),
				barchart.ShowValues(),
				barchart.BarWidth(20),
				barchart.Labels([]string{
					"ภายนอก - แนบแล้ว",
					"ภายนอก - ทั้งหมด",
					"ภายใน - แนบแล้ว",
					"ภายใน - ทั้งหมด",
				}),
			)
			if err != nil {
				panic(err)
			}
			go playBarChart(ctx, bc, 1*time.Second)
		
			c, err := container.New(
				t,
				container.Border(linestyle.Light),
				container.BorderTitle("PRESS Q TO QUIT"),
				container.PlaceWidget(bc),
			)
			if err != nil {
				panic(err)
			}
		
			quitter := func(k *terminalapi.Keyboard) {
				if k.Key == 'q' || k.Key == 'Q' {
					cancel()
				}
			}
		
			if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter)); err != nil {
				panic(err)
			}

		return 0
	})

	
	
	

	

	app := cli.New("sara is toolbox for manage sarabun duc work writing in golang").
	WithCommand(dirDate).
	WithCommand(dirSarabanDocScan).
	WithCommand(dirWatch).
	WithCommand(showPath)

	

	os.Exit(app.Run(os.Args, os.Stdout))
}