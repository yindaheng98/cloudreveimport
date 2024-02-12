from cloudreveimport import Invoker

invoker = Invoker(
    "D:\\Documents\\MyPrograms\\cloudreveimport\\cloudreveimport.exe",
    "-c", "D:\\Documents\\MyPrograms\\cloudreveimport\\test\\conf.ini",
    "-u", "yindaheng98@gmail.com",
    "-m", "-"
)
invoker.start()
for i in range(10):
    invoker.invoke(
        command="AAA",
        source_name="/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg",
        dst_path=["www", "ggg", "ooo", f"{i}.png"]
    )
invoker.join()
