import random
import time
from cloudreveimport import Invoker

invoker = Invoker(
    executable="D:\\Documents\\MyPrograms\\cloudreveimport\\cloudreveimport.exe",
    config="D:\\Documents\\MyPrograms\\cloudreveimport\\test\\conf.ini",
    email="yindaheng98@gmail.com",
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
    ctime=0,
    mtime=int(time.time()-30*24*3600),
    size=100
)
invoker.update_folder_time(
    dst_path=["www", "ggg", "ooo", f"{k}.png"],
    ctime=0,
    mtime=int(time.time()-30*24*3600),
)
invoker.update_folder_time(
    dst_path=["www", "ggg", "ooo"],
    ctime=0,
    mtime=int(time.time()-30*24*3600),
)
invoker.update_folder_time(
    dst_path=["groups"],
    ctime=int(time.time()-30*24*3600), # it works
    mtime=int(time.time()-30*24*3600), # it works!
)
invoker.join()
