FROM ${RepositoryUri}rust:1.58.1-slim-bullseye as builder
WORKDIR /app
COPY . .
ENV RUST_BACKTRACE=full
RUN apt update && apt install -y libudev-dev pkg-config && \
  cargo build --release

FROM ${RepositoryUri}debian:bullseye-20220125-slim
ARG version
ENV version=$version
MAINTAINER justin
COPY --from=builder /app/build/entrypoint.sh /opt/entrypoint.sh
COPY --from=builder /app/target/release/rust-hello /opt/main
WORKDIR /opt
ENTRYPOINT ["/opt/entrypoint.sh"]
CMD ./main

