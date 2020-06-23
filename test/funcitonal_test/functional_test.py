"""
This module contains functional tests of the cubeclient tool.
"""
import subprocess
import pytest
import requests
import shlex


test_body = "gogog go \ngo"
@pytest.mark.server(url='/test', response=[test_body])
def test_web_resourse_deafult_goroutines():
    cmd = "echo -e 'http://localhost:5000/test' | ./countgo"
    process = subprocess.Popen(['bash', '-c', cmd], stdout=subprocess.PIPE)         
    out, _ = process.communicate()
    assert(out == b"Count for http://localhost:5000/test: 4\n")

def test_web_resourse_one_goroutine():
    cmd = "echo -e 'http://localhost:5000/test' | ./countgo -k=1"
    process = subprocess.Popen(['bash', '-c', cmd], stdout=subprocess.PIPE)         
    out, _ = process.communicate()
    assert(out == b"Count for http://localhost:5000/test: 4\n")