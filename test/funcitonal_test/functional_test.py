"""
This module contains functional tests of the cubeclient tool.
"""
import subprocess
import pytest
import requests
import shlex


test_body = "gogog go \ngo"
@pytest.mark.server(url='/test', response=[test_body])
def test_web_resource_normal_response():
    cmd = "echo -e 'http://localhost:5000/test' | ./countgo"
    process = subprocess.Popen(['bash', '-c', cmd], stdout=subprocess.PIPE)
    out, _ = process.communicate()
    assert(out == b"Count for http://localhost:5000/test: 4\n")


def test_web_resource_one_worker():
    cmd = "echo -e 'http://localhost:5000/test' | ./countgo -k=1"
    process = subprocess.Popen(['bash', '-c', cmd], stdout=subprocess.PIPE)
    out, _ = process.communicate()
    assert(out == b"Count for http://localhost:5000/test: 4\n")


def test_file_resource_emty_file():
    cmd = "echo -e './test/test_files/empty_test.txt' | ./countgo -k=1"
    process = subprocess.Popen(['bash', '-c', cmd], stdout=subprocess.PIPE)
    out, _ = process.communicate()
    assert(out == b"Count for ./test/test_files/empty_test.txt: 0\n")


def test_file_resource_normal_file():
    cmd = "echo -e './test/test_files/simple_test.txt' | ./countgo -k=1"
    process = subprocess.Popen(['bash', '-c', cmd], stdout=subprocess.PIPE)
    out, _ = process.communicate()
    assert(out == b"Count for ./test/test_files/simple_test.txt: 4\n")


def test_multiple_resources():
    multiple_resources_cmd = "echo -e \
'./test/test_files/simple_test.txt\nhttp://localhost:5000/test\n\
./test/test_files/empty_test.txt' | ./countgo -k=2"
    process = subprocess.Popen(
        ['bash', '-c', multiple_resources_cmd], stdout=subprocess.PIPE)
    out, _ = process.communicate()
    assert(
        b'Count for ./test/test_files/simple_test.txt: 4' in out and
        b'Count for ./test/test_files/empty_test.txt: 0' in out and
        b'Count for http://localhost:5000/test: 4' in out
    )


def test_big_input_data():
    cmd = "echo -e '" + "./test/test_files/simple_test.txt\n"*100 + "'\
         | ./countgo"
    process = subprocess.Popen(['bash', '-c', cmd], stdout=subprocess.PIPE)
    out, _ = process.communicate()
    assert(out.count(b'Count for ./test/test_files/simple_test.txt: 4') == 100)


def test_file_resource_no_such_file_error():
    cmd = "echo -e './test/test_files/no_such_file.txt' | ./countgo -k=1"
    process = subprocess.Popen(['bash', '-c', cmd], stdout=subprocess.PIPE)
    out, _ = process.communicate()
    assert(
        out == b"Failed to count 'go' in ./test/test_files/no_such_file.txt\n"
        )


def test__web_resource_unavailable_error():
    cmd = "echo -e 'http://localhost:0/test' | ./countgo -k=1"
    process = subprocess.Popen(['bash', '-c', cmd], stdout=subprocess.PIPE)
    out, _ = process.communicate()
    assert(
        out == b"Failed to count 'go' in http://localhost:0/test\n"
        )
