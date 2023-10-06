package main

import (
	"github.com/teris-io/cli"
	"github.com/daviddengcn/go-colortext"
	"os"
	
	// "os/exec"
	"fmt"
	"strings"
	"io/ioutil"
	"path/filepath"
	"context"
	// "math/rand"
	"time"
	"encoding/csv"
	"github.com/skratchdot/open-golang/open"	

	"github.com/mum4k/termdash"
	// "github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	// "github.com/mum4k/termdash/widgets/barchart"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/mum4k/termdash/widgets/gauge"
	"unicode/utf8"

)



func countFileInDir(path string,minusBy int)(countFileInDir_out int){
    files,_ := ioutil.ReadDir(path)
    myInt := len(files)-minusBy
    return myInt
}
type playType int

const (
	playTypePercent playType = iota
	playTypeAbsolute
)

// func playBarChart(ctx context.Context, bc *barchart.BarChart, delay time.Duration) {
// 	max := 100

	

// 	ticker := time.NewTicker(delay)
// 	defer ticker.Stop()
// 	for {
// 		out_ready := countFileInDir("./ภายนอก/เตรียมแนบ",1)
// 		out_done := countFileInDir("./ภายนอก/เตรียมแนบ/แนบแล้ว",0)
// 		out_sum := out_ready + out_done
// 		in_ready := countFileInDir("./ภายใน/เตรียมแนบ",1)
// 		in_done :=  countFileInDir("./ภายใน/เตรียมแนบ/แนบแล้ว",0)
// 		in_sum := in_ready + in_done
// 		// in_out := out_sum + in_sum
// 		if in_sum > out_sum {
// 			max = in_sum
// 		}else{
// 			max = out_sum
// 		}
		
// 		select {
// 		case <-ticker.C:
// 			// var brBarColor []cell.Color
// 			var values []int
// 			for i := 0; i < 4; i++ {
// 				switch i {
// 				case 0:
// 					// if out_done == out_sum {
// 					// 	brBarColor = append(brBarColor,  cell.ColorGreen)
// 					// }else{
// 					// 	brBarColor = append(brBarColor,  cell.ColorRed)
// 					// }
				
// 					values = append(values,  out_done)
// 				case 1:
// 					values = append(values, out_sum )
// 				case 2:
// 					// if in_done == in_sum {
// 					// 	brBarColor = append(brBarColor,  cell.ColorGreen)
// 					// }else{
// 					// 	brBarColor = append(brBarColor,  cell.ColorRed)
// 					// }
// 					values = append(values, in_done )
// 				case 3:
// 					values = append(values, in_sum)
// 				}
// 				// values = append(values, int(rand.Int31n(max+1)))
// 			}

// 			if err := bc.Values(values, max); err != nil {
// 				panic(err)
// 			}

// 			// bc.BarColors(brBarColor)


// 		case <-ctx.Done():
// 			return
// 		}
// 	}
// }

func playGauge(ctx context.Context, g *gauge.Gauge, step int, delay time.Duration, pt playType) {
	progress := 0
	mult := 1

	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			switch pt {
			case playTypePercent:
				if err := g.Percent(progress); err != nil {
					panic(err)
				}
			case playTypeAbsolute:
				if err := g.Absolute(progress, 100); err != nil {
					panic(err)
				}
			}

			progress += step * mult
			if progress > 100 || 100-progress < step {
				progress = 100
			} else if progress < 0 || progress < step {
				progress = 0
			}

			if progress == 100 {
				mult = -1
			} else if progress == 0 {
				mult = 1
			}

		case <-ctx.Done():
			return
		}
	}
}


func playGauge_in(ctx context.Context, g *gauge.Gauge, step int, delay time.Duration, pt playType) {
	progress := 0
	mult := 1

	

	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			in_ready := countFileInDir("./ภายใน/เตรียมแนบ",1)
			in_done :=  countFileInDir("./ภายใน/เตรียมแนบ/แนบแล้ว",0)
			in_sum := in_ready + in_done

			if in_sum == 0 {
				in_sum = 101010101
			}
			switch pt {
			case playTypePercent:
				if err := g.Percent(progress); err != nil {
					panic(err)
				}
			case playTypeAbsolute:
				if err := g.Absolute(in_done, in_sum); err != nil {
					panic(err)
				}
			}


			progress += step * mult
			if progress > 100 || 100-progress < step {
				progress = 100
			} else if progress < 0 || progress < step {
				progress = 0
			}

			if progress == 100 {
				mult = -1
			} else if progress == 0 {
				mult = 1
			}

		case <-ctx.Done():
			return
		}
	}
}



func playGauge_out(ctx context.Context, g *gauge.Gauge, step int, delay time.Duration, pt playType) {
	progress := 0
	mult := 1

	

	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			
			out_ready := countFileInDir("./ภายนอก/เตรียมแนบ",1)
			out_done := countFileInDir("./ภายนอก/เตรียมแนบ/แนบแล้ว",0)
			out_sum := out_ready + out_done

			if out_sum == 0 {
				out_sum = 101010101
			}
			switch pt {
			case playTypePercent:
				if err := g.Percent(progress); err != nil {
					panic(err)
				}
			case playTypeAbsolute:
				if err := g.Absolute(out_done, out_sum); err != nil {
					panic(err)
				}
			}
			// out_pro_text := false
			// if out_sum == 0 {
			// 	out_pro_text = false
			// }else{
			// 	out_pro_text = true
			// }
		
			progress += step * mult
			if progress > 100 || 100-progress < step {
				progress = 100
			} else if progress < 0 || progress < step {
				progress = 0
			}

			if progress == 100 {
				mult = -1
			} else if progress == 0 {
				mult = 1
			}

		case <-ctx.Done():
			return
		}
	}
}



func playGauge_all(ctx context.Context, g *gauge.Gauge, step int, delay time.Duration, pt playType) {
	progress := 0
	mult := 1

	

	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			
			in_ready := countFileInDir("./ภายใน/เตรียมแนบ",1)
			in_done :=  countFileInDir("./ภายใน/เตรียมแนบ/แนบแล้ว",0)
			in_sum := in_ready + in_done

			out_ready := countFileInDir("./ภายนอก/เตรียมแนบ",1)
			out_done := countFileInDir("./ภายนอก/เตรียมแนบ/แนบแล้ว",0)
			out_sum := out_ready + out_done

			all_done := in_done + out_done
			all_sum := in_sum + out_sum
			
			if out_sum == 0 {
				out_sum = 101010101
			}
			switch pt {
			case playTypePercent:
				if err := g.Percent(progress); err != nil {
					panic(err)
				}
			case playTypeAbsolute:
				if err := g.Absolute(all_done, all_sum); err != nil {
					panic(err)
				}
			}
			// out_pro_text := false
			// if out_sum == 0 {
			// 	out_pro_text = false
			// }else{
			// 	out_pro_text = true
			// }
		
			progress += step * mult
			if progress > 100 || 100-progress < step {
				progress = 100
			} else if progress < 0 || progress < step {
				progress = 0
			}

			if progress == 100 {
				mult = -1
			} else if progress == 0 {
				mult = 1
			}

		case <-ctx.Done():
			return
		}
	}
}

// regisInNum

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
		YYMMDD := string(time.Now().AddDate(543, 0, 0).Format("060102"))
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
		YYMMDD := string(time.Now().AddDate(543, 0, 0).Format("06.01.02"))
		mkdirYYMMDD := os.Mkdir(YYMMDD,0666)
		if mkdirYYMMDD != nil {
			showErr("error :",mkdirYYMMDD)
		}else{
			showLogString("dirDate : ",YYMMDD)
		}
		os.Mkdir(YYMMDD+"/ออกเลข",0666)
		os.Mkdir(YYMMDD+"/ภายนอก",0666)
		os.Mkdir(YYMMDD+"/ภายใน",0666)
		os.Mkdir(YYMMDD+"/ไม่เสนอ",0666)
		os.Mkdir(YYMMDD+"/แจ้งเวียน",0666)
		
		os.Mkdir(YYMMDD+"/ภายนอก/เตรียมแนบ",0666)
		os.Mkdir(YYMMDD+"/ภายใน/เตรียมแนบ",0666)
		os.Mkdir(YYMMDD+"/ไม่เสนอ/เตรียมแนบ",0666)
		os.Mkdir(YYMMDD+"/ภายนอก/เตรียมแนบ/แนบแล้ว",0666)
		os.Mkdir(YYMMDD+"/ภายใน/เตรียมแนบ/แนบแล้ว",0666)
		os.Mkdir(YYMMDD+"/ไม่เสนอ/เตรียมแนบ/แนบแล้ว",0666)
		os.Mkdir(YYMMDD+"/แจ้งเวียน/แปลงแล้ว",0666)
		os.Mkdir(YYMMDD+"/แจ้งเวียน/แปลงแล้ว/เวียนแล้ว",0666)
		
		open.Start(YYMMDD)

		return 0
	})

	genCSV := cli.NewCommand("genCSV", "...").
	WithShortcut("csv").
	WithAction(func(args []string, options map[string]string) int {

		targetDirs := []string{"./ภายใน/เตรียมแนบ", "./ภายนอก/เตรียมแนบ"}
		

		var fileNames []string
		var docUIDs [][]string
	

		for _, targetDir := range targetDirs {
			fileInfos, err := ioutil.ReadDir(targetDir)
			if err != nil {
				fmt.Printf("Error reading directory %s: %v\n", targetDir, err)
				continue
			}

			for _, fileInfo := range fileInfos {
				if !fileInfo.IsDir() {
					fileName := fileInfo.Name()
					filePath := filepath.Join(targetDir, fileName)
					fileNames = append(fileNames, filePath)
				}
			}

			
			for _, fileInfo := range fileInfos {
				if !fileInfo.IsDir() {
					fileName := fileInfo.Name()
					parts := strings.Split(fileName, "_")
					rawDocUID := parts[0]
					docUIDParts := strings.Split(rawDocUID, " ")
					docUIDs = append(docUIDs, docUIDParts)
				}
			}
		}

	
		file, err := os.Create("output.csv")
		if err != nil {
			fmt.Println("Error:", err)
			return 0
		}
		defer file.Close()
	
		// Write BOM
		file.Write([]byte{0xEF, 0xBB, 0xBF})
	
		writer := csv.NewWriter(file)
		err = writer.WriteAll(docUIDs)
		if err != nil {
			fmt.Println("Error:", err)
			return 0
		}
	
		// Check file encoding
		content, err := ioutil.ReadFile("output.csv")
		if err != nil {
			fmt.Println("Error:", err)
			return 0
		}
	
		encoding := "unknown"
		if utf8.Valid(content) {
			encoding = "UTF-8"
			if bytes := content[:3]; bytes[0] == 0xEF && bytes[1] == 0xBB && bytes[2] == 0xBF {
				encoding = "UTF-8-BOM"
			}
		}
	
		fmt.Printf("File encoding: %s\n", encoding)		

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



			borderless, err := text.New()
			if err != nil {
				panic(err)
			}
			borderless.Write("Text without border.");


			in_pro, err := gauge.New(
				gauge.Height(1),
				gauge.Border(linestyle.Light),
				gauge.BorderTitle("in doc progress"),
			)

		    

			out_pro, err := gauge.New(
				gauge.Height(1),
				gauge.Border(linestyle.Light),
				gauge.BorderTitle("out doc progress"),
			)
			
			
			all_pro, err := gauge.New(
				gauge.Height(1),
				gauge.Border(linestyle.Light),
				gauge.BorderTitle("all doc progress"),
			)

			go playGauge_in(ctx, in_pro, 17, 500*time.Millisecond, playTypeAbsolute)
			go playGauge_out(ctx, out_pro, 17, 500*time.Millisecond, playTypeAbsolute)
			go playGauge_all(ctx, all_pro, 17, 500*time.Millisecond, playTypeAbsolute)

			// bc, err := barchart.New(
				
			// 	barchart.BarColors([]cell.Color{
			// 		cell.ColorYellow,
			// 		cell.ColorYellow,
			// 		cell.ColorYellow,
			// 		cell.ColorYellow,
			// 		// cell.ColorGreen,
			// 	}),
			// 	barchart.ValueColors([]cell.Color{
			// 		cell.ColorRed,
			// 		cell.ColorRed,
			// 		cell.ColorRed,
			// 		cell.ColorRed,
			// 	}),
			// 	barchart.ShowValues(),
			// 	barchart.BarWidth(20),
			// 	barchart.Labels([]string{
			// 		"ภายนอก - แนบแล้ว",
			// 		"ภายนอก - ทั้งหมด",
			// 		"ภายใน - แนบแล้ว",
			// 		"ภายใน - ทั้งหมด",
			// 	}),
			// )
			// if err != nil {
			// 	panic(err)
			// }

			
		

			// go playBarChart(ctx, bc, 1*time.Second)
		
			c, err := container.New(
				t,
				container.Border(linestyle.Light),
				container.BorderTitle("PRESS Q TO QUIT"),
				container.SplitVertical(
					
					container.Left(
						container.Border(linestyle.Light),
						container.BorderTitle("Document Count"),
						container.SplitHorizontal(
							container.Top(
								container.SplitHorizontal(
									container.Top(
										container.PlaceWidget(in_pro),
									),
									container.Bottom(
										container.PlaceWidget(out_pro),
									),
								),
							),
							container.Bottom(
								container.PlaceWidget(all_pro),
							),
						),
					
						
						
					),
					container.Right(
						container.Border(linestyle.Light),
						container.BorderTitle("something"),
						container.PlaceWidget(borderless),
					),
				),
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
	WithCommand(showPath).
	WithCommand(genCSV)
	

	

	os.Exit(app.Run(os.Args, os.Stdout))
}