// gatherfile_main
// by ggenien@163.com
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	c_contain = "+"
	c_without = "-"
)

var (
	target_dir                         string   //target dir
	filename_contain, filename_without []string //filename contain or without
	count                              int      //files count, for rename

	c_sep = string(os.PathSeparator) //path separator
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("gatherfile targetDir sourceDir +filenameContain -filenameWithout")
		return
	}

	target_dir = os.Args[1]
	source_dir := os.Args[2]
	if strings.HasSuffix(target_dir, c_sep) {
		target_dir = strings.TrimSuffix(target_dir, c_sep)
	}
	if strings.HasSuffix(source_dir, c_sep) {
		source_dir = strings.TrimSuffix(source_dir, c_sep)
	}

	filename_contain = make([]string, 0, 10)
	filename_without = make([]string, 0, 10)
	count = 0

	//filename contain or without
	wrong := false
	for i := 3; i < len(os.Args); i++ {
		p := os.Args[i]
		if strings.HasPrefix(p, c_contain) && (len(p) > len(c_contain)) {
			filename_contain = append(filename_contain, p[1:])
		} else if strings.HasPrefix(p, "-") && (len(p) > len(c_without)) {
			filename_without = append(filename_without, p[1:])
		} else {
			fmt.Println("parameter should begin with + or -:", p)
			wrong = true
		}
	}
	if wrong {
		return
	}

	if f_create_multi_path(target_dir) {
		f_deal_path(source_dir)
	}

	fmt.Println("files gathered:", count)
}

// deal a path and files in it
func f_deal_path(the_path string) {
	files, err := ioutil.ReadDir(the_path)
	if err != nil {
		return
	}

	for i := range files {
		fn := files[i].Name()

		if files[i].IsDir() {
			f_deal_path(the_path + c_sep + fn)
		} else {
			for i := range filename_contain {
				if !strings.Contains(fn, filename_contain[i]) {
					goto l_next_file
				}
			}
			for i := range filename_without {
				if strings.Contains(fn, filename_without[i]) {
					goto l_next_file
				}
			}
			count++
			fmt.Printf("%08d %s\n", count, fn)
			tar_fn := target_dir + c_sep + fmt.Sprintf("%08d%s", count, path.Ext(fn))
			f_copy_file(the_path+c_sep+fn, tar_fn)
		}

	l_next_file:
	}
}

func f_create_multi_path(the_path string) bool {
	dirs := strings.Split(the_path, c_sep)
	if len(dirs) <= 0 {
		return false
	}
	d := ""
	for i := range dirs {
		d += dirs[i] + c_sep
		f_create_path(d)
	}
	return true
}

// create a path
func f_create_path(the_path string) {
	_, err := os.Stat(the_path)
	if err == nil { //dir exists
		return
	}
	if os.IsNotExist(err) { //dir not exists, create
		os.Mkdir(the_path, os.ModeDir)
	}
}

//copy and rename a file
func f_copy_file(src_sep, tar_dir string) bool {
	src_file, err := os.Open(src_sep)
	if err != nil {
		return false
	}
	defer src_file.Close()
	dst_file, err := os.Create(tar_dir)
	if err != nil {
		return false
	}
	defer dst_file.Close()

	_, err = io.Copy(dst_file, src_file)
	return err == nil
}
