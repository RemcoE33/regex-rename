# Regex Rename

Simple cli to rename files in current, relative or fill path directory. You can use regex groups to to use as new filename. Currently only the first capture group is supported. 

I needed this for some repeated tasks.

```bash
go install -o rrn github.com/RemcoE33/regex-rename

rrn ^_\d{4}_(.*)$ /img

```
