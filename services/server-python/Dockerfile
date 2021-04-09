FROM python:3.9-slim

WORKDIR /work

COPY requirements.txt .

RUN pip install -U pip \
    && pip install -r requirements.txt

COPY . .

CMD ["python3", "/work/main.py"]
