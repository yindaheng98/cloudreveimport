import json
import subprocess
from multiprocessing import Process


def stdout_reader(process: subprocess.Popen):
    while True:
        output = process.stdout.readline().decode('utf-8')
        if not output and process.poll() is not None:
            break
        print("stdout", output.strip())


def stderr_reader(process: subprocess.Popen):
    while True:
        output = process.stderr.readline().decode('utf-8')
        if not output and process.poll() is not None:
            break
        print("stderr", output.strip())


class Invoker:
    def __init__(self, executable, *args):
        self.executable = executable
        self.args = args
        self.process = None
        self.stdout_reader = None
        self.stderr_reader = None

    def start(self):
        self.process = subprocess.Popen(args=self.args, executable=self.executable, stdout=subprocess.PIPE)
        self.stdout_reader = Process(target=stdout_reader, args=(self.process,))
        self.stderr_reader = Process(target=stderr_reader, args=(self.process,))
        self.stdout_reader.start()
        self.stderr_reader.start()

    def invoke(self, command, **kwargs):
        if not self.process:
            print("not started")
            return
        self.process.stdin.write(json.dumps({
            **kwargs,
            "command": command
        }))

    def import_file(self, dst_path, src_name):
        self.invoke(command="ImportFile", dst_path=dst_path, src_name=src_name)

    def update_file_time(self, dst_path, ctime, mtime):
        self.invoke(command="UpdateFileTime", dst_path=dst_path, created_at=ctime, updated_at=mtime)

    def update_folder_time(self, dst_path, ctime, mtime):
        self.invoke(command="UpdateFolderTime", dst_path=dst_path, created_at=ctime, updated_at=mtime)

    def join(self):
        self.process.stdin.close()
        self.stdout_reader.join()
        self.stderr_reader.join()
        self.process.kill()
        self.process = None
        self.stdout_reader = None
        self.stderr_reader = None
