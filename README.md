# cloudreveimport

Import your existing files into [Cloudreve](https://github.com/cloudreve/Cloudreve)!

## Install

```sh
go install github.com/yindaheng98/cloudreveimport
pip install cloudreveimport
```

## Usage

```python
import time
from cloudreveimport import Invoker
invoker = Invoker(
    "path to execuatable file you just install by 'go install ...'",
    "-c", "path to your conf.ini for your cloudreve",
    "-u", "your email in cloudreve",
    "-m", "-" # This means read commands from stdin
)
invoker.start()
invoker.import_file(
    source_name="path to the file on your disk",
    dst_path=["path", "you", "want", "to", "import", "to", "in", "cloudreve"]
)
invoker.update_file_time(
    dst_path=["path", "to", "file", "in", "cloudreve"],
    ctime=int(time.time()), # created time
    mtime=int(time.time()), # updated time
)
invoker.update_folder_time(
    dst_path=["path", "to", "folder", "in", "cloudreve"],
    ctime=int(time.time()), # created time
    mtime=int(time.time()), # updated time
)
```