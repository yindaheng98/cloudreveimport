# cloudreveimport

Import your existing files into [Cloudreve](https://github.com/cloudreve/Cloudreve)!

## Install

### Build executable file

```sh
git clone https://github.com/yindaheng98/cloudreveimport
cd cloudreveimport
go build
```

Then you shall get an executable file `cloudreveimport`, you can try it:

```sh
./cloudreveimport -h
Usage of D:\Documents\MyPrograms\cloudreveimport\cloudreveimport.exe:
  -c string
        Path to the config file. (default "D:\\Documents\\MyPrograms\\cloudreveimport\\conf.ini")
  -m string
        A folder to be imported.
  -t string
        Import the folder to which folder in cloudreve.
  -u string
        Email of the target user. (default "admin@cloudreve.org")
```

### Install Python package

```sh
pip install cloudreveimport
```

## Usage

### Command line

Just run it in the device where the Cloudreve is installed:

Assume your Cloudreve config file is `/root/conf.ini`, and you want to import a folder `/mnt/gallery` into the user `me@cloudreve.org`'s `my/gallery` folder in Cloudreve.
Just run:

```sh
./cloudreveimport -c "/root/conf.ini" -u "me@cloudreve.org" -m "/mnt/gallery" -t "my/gallery"
```

Then open `me@cloudreve.org`'s `my/gallery` folder in Cloudreve, and you will see those files in `/mnt/gallery` here.

### Python

Or you can write your scripts in Python:

```python
import time
from cloudreveimport import Invoker
invoker = Invoker(
    execuable="path to executable file you just built",
    config="path to your conf.ini for your cloudreve",
    email="your email in cloudreve"
)
invoker.start()
invoker.import_file(
    source_name="path to the file on your disk",
    dst_path=["path", "you", "want", "to", "import", "to", "in", "cloudreve"],
    size=100, # filesize in Bytes
)
invoker.update_file_stat(
    dst_path=["path", "to", "file", "in", "cloudreve"],
    ctime=int(time.time()), # created time
    mtime=int(time.time()), # updated time
    size=100, # filesize in Bytes
)
invoker.delete_file(
    dst_path=["path", "to", "file", "in", "cloudreve"],
)
invoker.update_folder_time(
    dst_path=["path", "to", "folder", "in", "cloudreve"],
    ctime=int(time.time()), # created time
    mtime=int(time.time()), # updated time
)
```

Note: 上述文件操作是直接对数据库进行修改，会导致用户已用空间与实际不符，所以在完成文件的导入或文件大小修改后需要执行数据库脚本[校准用户容量](https://docs.cloudreve.org/v/en/manage/db-script#xiao-zhun-yong-hu-rong-liang)。