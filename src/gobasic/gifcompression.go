package gobasic

import (
	"image"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

const vieux string = "/home/lzd/study/Go_study/go_web_programming/"

func GifCompression() {
	log.Println("处理当前文件夹的gif,并把旧文件放入ori_image")
	_, err := os.Stat(vieux)
	if err != nil {
		log.Println(err)
		os.MkdirAll(vieux, 0666) //8进制 要有0
	}
	log.Println(regexp.QuoteMeta(`^(.\\)?go_web_programming\\?`), regexp.MustCompile(regexp.QuoteMeta(`^(.\\)?go_web_programming\\?`)))
	filepath.Walk("/home/lzd/study/Go_study/go_web_programming/test.gif", func(path string, info os.FileInfo, err error) error {
		//log.Println(path,filepath.Ext(path))//,filepath.Dir(path))
		//matched,err:=regexp.MatchString(
		//log.Println( regexp.MustCompile( (`^(.\\)?ori_image\\?`)).MatchString( filepath.Dir(path)) )
		// if info.IsDir() || regexp.MustCompile((`^(.\\)?go_web_programming\\?`)).MatchString(filepath.Dir(path)) || strings.ToLower(filepath.Ext(path)) != ".gif" {
		// 	return nil
		// }
		// if filepath.Dir(path) != "." {
		// 	return nil
		// }
		log.Println(path, filepath.Ext(path)) //,filepath.Dir(path))
		log.Println("处理", path)
		defer func() {
			//注意 如果加上这上会导致无法打印出堆栈错误的详细信息 开发阶段不要加
			/*if r := recover(); r != nil {
			    log.Printf("Runtime error caught: %v", r)
			} */
		}()
		skip_frame(path)
		return nil
	})

}
func skip_frame(s string) {
	f, err := os.Open(s)
	echo_err(err)
	pic, err := gif.DecodeAll(f)
	echo_err(err)
	count := len(pic.Image)
	log.Println("帧数量 ", count)
	// if count < 50 {
	// 	return
	// }

	// count / 2
	// k := count / 35 //保留35帧
	mid := "_抽帧"
	//newpath
	neo := make([]*image.Paletted, count/20)
	neoDisposal := make([]byte, count/20)
	for i := 0; i < count/20; i++ {
		neo[i] = pic.Image[int(i*20)]
		neoDisposal[i] = pic.Disposal[int(i*20)]
		pic.Delay[i] = int(pic.Delay[i] * 20 * 4)
	}
	pic.Image = neo
	pic.Delay = pic.Delay[:count/20]
	pic.Disposal = neoDisposal
	newpath := filepath.Base(s) + mid + ".gif"
	output, err := os.Create(newpath)
	echo_err(err)
	err = gif.EncodeAll(output, pic)
	echo_err(err)

	//处理后移动走
	f.Close()
	output.Close()
	// err = os.Rename(s, vieux+s)
	// echo_err(err)
}
func echo_err(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
