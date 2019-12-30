# GatherFile
gather files from a directory tree, copy and rename the files into new directory, useful for AI training photos

usage: gatherfile targetDir sourceDir +filenameContain -filenameWithout
for example: gatherfile d:\train\type1 d:\photos +.jpg +car -sport
this will gather all files which name cantain ".jpg" and "car", but not contain "sport", into d:\train\type1 from d:\photos

开发这个小工具，是因为训练TensorFlow的时候，大量图片位于不同目录，文件名又重名，收集太麻烦，所以做了这个工具。
