import urllib
from locust import HttpLocust, TaskSet
from hypothesis import strategy

source = stragety.integers(min_value=0, max_value=10000)
sliceLen = strategy.integers(min_value=1, max_value=15)

class RandomBehavior(TaskSet):
    def request(self):
        nums = [source.example() for _ in range(1, sliceLen.example())]
        params = urllib.parse.urlencode({"num": ",".join([str(x) for x in nums])})
        self.client.get("/?" + params)

class Client(HttpLocust):
    task_set = RandomBehavior
    min_wait = 1000
    max_wait = 5000