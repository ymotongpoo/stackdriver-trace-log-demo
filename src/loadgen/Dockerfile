FROM python:3.8-slim
RUN apt-get -qq update \
    && apt-get install -y --no-install-recommends \
        g++ build-essential libffi-dev
COPY requirements.txt .
COPY constraints.txt .
RUN pip install -r requirements.txt -c constraints.txt
COPY . .
ENTRYPOINT ./loadgen.sh