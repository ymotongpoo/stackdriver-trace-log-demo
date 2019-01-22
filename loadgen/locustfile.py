import urllib
from locust import HttpLocust, TaskSet, task
from hypothesis import strategies

source = strategies.integers(min_value=0, max_value=10000)
sliceLen = strategies.integers(min_value=1, max_value=15)

class RandomBehavior(TaskSet):
    @task(1)
    def request(self):
        nums = [source.example() for _ in range(1, sliceLen.example())]
        params = urllib.parse.urlencode({"num": ",".join([str(x) for x in nums])})
        self.client.get("/?" + params)

class Client(HttpLocust):
    task_set = RandomBehavior
    min_wait = 3000
    max_wait = 10000