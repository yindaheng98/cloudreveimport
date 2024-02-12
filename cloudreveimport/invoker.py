import json
import subprocess
from threading import Thread


def stdout_reader(process: subprocess.Popen):
    while True:
        output = process.stdout.readline().decode('utf-8')
        if not output and process.poll() is not None:
            break
        if len(output.strip()) <= 0:
            continue
        print("stdout", output.strip())


def stderr_reader(process: subprocess.Popen):
    while True:
        output = process.stderr.readline().decode('utf-8')
        if not output and process.poll() is not None:
            break
        if len(output.strip()) <= 0:
            continue
        print("stderr", output.strip())


class Invoker:
    def __init__(self, executable, *args):
        self.executable = executable
        self.args = args
        self.process = None
        self.stdout_reader = None
        self.stderr_reader = None

    def start(self):
        self.process = subprocess.Popen(
            args=[self.executable, *self.args],
            stdin=subprocess.PIPE,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE
        )
        self.stdout_reader = Thread(target=stdout_reader, args=(self.process,))
        self.stderr_reader = Thread(target=stderr_reader, args=(self.process,))
        self.stdout_reader.start()
        self.stderr_reader.start()

    def invoke(self, command, **kwargs):
        if not self.process:
            print("not started")
            return
        data = json.dumps({
            **kwargs,
            "command": command
        })
        self.process.stdin.write((data + "\n").encode("utf8"))

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
