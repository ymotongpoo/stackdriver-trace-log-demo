# Copyright 2019 Yoshi Yamaguchi
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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