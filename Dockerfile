FROM golang:1.18

# RUN apt-get install -qq -y llvm
# RUN apt-get install -qq -y clang
# RUN apt-get install -qq -y make

RUN echo 'deb http://apt.llvm.org/bullseye/ llvm-toolchain-bullseye-15 main' | tee /etc/apt/sources.list.d/llvm.list

RUN wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | apt-key add -
RUN apt-get -qq -y update
RUN apt-get -qq -y upgrade
RUN apt-get install -y clang-15 llvm-15-dev lld-15 libclang-15-dev
RUN apt-get install -y make

WORKDIR /work
RUN go install honnef.co/go/tools/cmd/staticcheck@2022.1
RUN go install -v golang.org/x/lint/golint@latest

COPY . .

# A little hacky.  A better way?:
RUN ln -s /usr/bin/clang-15 /usr/bin/clang

RUN make test
