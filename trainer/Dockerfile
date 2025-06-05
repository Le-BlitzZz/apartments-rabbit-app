FROM python:3.13-slim

# Add dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    bash \
    make \
    curl \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

RUN curl -sSL https://install.python-poetry.org | python3 - && \
    ln -s ~/.local/bin/poetry /usr/local/bin/poetry && \
    poetry config virtualenvs.create false

CMD ["tail", "-f", "/dev/null"]
