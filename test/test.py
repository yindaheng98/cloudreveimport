import random
import time
from cloudreveimport import Invoker

invoker = Invoker(
    "D:\\Documents\\MyPrograms\\cloudreveimport\\cloudreveimport.exe",
    "-c", "D:\\Documents\\MyPrograms\\cloudreveimport\\test\\conf.ini",
    "-u", "yindaheng98@gmail.com",
    "-m", "-"
)
invoker.start()
for i in range(3):
    invoker.invoke(
        command="AAA",
        source_name="/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg",
        dst_path=["www", "ggg", "ooo", f"{i}.png"]
    )
k = random.randint(1, 1000)
invoker.import_file(
    source_name="/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg",
    dst_path=["www", "ggg", "ooo", f"{k}.png"]
)
invoker.update_file_time(
    dst_path=["www", "ggg", "ooo", f"{k}.png"],
    ctime=int(time.time()),
    mtime=int(time.time()-100),
)
invoker.update_folder_time(
    dst_path=["www", "ggg", "ooo", f"{k}.png"],
    ctime=int(time.time()),
    mtime=int(time.time()-100),
)
invoker.update_folder_time(
    dst_path=["www", "ggg", "ooo"],
    ctime=int(time.time()),
    mtime=int(time.time()-100),
)
invoker.join()
