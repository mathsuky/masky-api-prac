FROM python:3.10-slim

WORKDIR /app

COPY requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 5100

ENV FLASK_APP=app.py

CMD ["flask", "run", "--host=0.0.0.0", "--port=5100"]
