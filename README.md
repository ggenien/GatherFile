# GatherFile
**Search for files from the source folder and its subfolders, copy them to the target folder and rename them. You can set search conditions so that the filename contains or does not contain anything.**

When training TensorFlow, a large number of pictures are located in different directories, and the file names are often duplicate. It's too troublesome to collect, so I wrote this tool.

Only one source file gatherfile.go, compiled with GO 1.12.7.
I have tested it in windows 10 x64, and CentOS 7.4 x64.

**usage:**

    gatherfile  targetDir   sourceDir  [/c] [+filenameContain] [-filenameWithout]
       /c: use if you need case sensitive, default is not sensitive
       +filenameContain: gather filename which contain a string, can be used multiple times
       +filenameWithout: gather filename which not contain a string, can be used multiple times

**example:**

    gatherfile d:\cars\train\type1 d:\original +car +.jpg -sport
	gatherfile . d:\original
	./gatherfile xyz data +.log

**功能：从源文件夹及其子文件夹中搜索符合条件的文件，复制到目标文件夹并改名。可设置搜索条件，使文件名包含或不包含什么。**

在训练TensorFlow的时候，大量图片位于不同目录，文件名又经常重名，收集太麻烦，所以写了这个工具。
